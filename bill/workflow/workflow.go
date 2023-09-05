package workflow

import (
	"strconv"
	"time"

	"encore.app/bill/activity"
	"encore.app/bill/repository"

	"github.com/shopspring/decimal"
	// enumspb "go.temporal.io/api/enums/v1"
	"go.temporal.io/sdk/workflow"
)

const (
	// QueryBill is the name of the query that returns the bill ID
	QueryBill = "query"

	// ChargeBillSignal is the name of the signal that charges a bill
	SignalChargeBill = "charge-bill"
)

// GetChargeWorkflowID returns the workflow ID for the charge workflow
func GetChargeWorkflowID(billID int) string {
	return "bill-charges-workflow-" + strconv.Itoa(billID)
}

func CreateBill(ctx workflow.Context, bill repository.Bill) (int, error) {
	var billID = 0
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 5,
	})
	err := workflow.
		ExecuteActivity(ctx, activity.CreateBill, bill).
		Get(ctx, &billID)
	if err != nil {
		return 0, err
	}

	// we execute the charge workflow and abandon it until the bill is closed
	ctx = workflow.WithChildOptions(ctx, workflow.ChildWorkflowOptions{
		WorkflowID:         GetChargeWorkflowID(billID),
		TaskQueue:          workflow.GetInfo(ctx).TaskQueueName,
		WorkflowRunTimeout: time.Second*time.Duration(bill.TimePeriod) + time.Minute,
		// ParentClosePolicy:  enumspb.PARENT_CLOSE_POLICY_REQUEST_CANCEL,
	})

	ch := workflow.ExecuteChildWorkflow(ctx, ChargeBill, bill).GetChildWorkflowExecution()
	if err := ch.Get(ctx, nil); err != nil {
		return 0, err
	}

	return billID, nil
}

func ChargeBill(ctx workflow.Context, bill repository.Bill) error {
	selector := workflow.NewSelector(ctx)

	var closeBill bool
	println(1)
	closeBillTimer := workflow.NewTimer(ctx, time.Second*time.Duration(bill.TimePeriod))
	selector.AddFuture(closeBillTimer, func(future workflow.Future) {
		println(2)
		closeBill = true
	})

	err := workflow.SetQueryHandler(ctx, QueryBill, func() (repository.Bill, error) {
		return bill, nil
	})
	if err != nil {
		return err
	}

	var amount decimal.Decimal
	channel := workflow.GetSignalChannel(ctx, SignalChargeBill)
	selector.AddReceive(channel, func(c workflow.ReceiveChannel, more bool) {
		println(3)
		channel.ReceiveAsync(&amount)
	})

	selector.Select(ctx)
	println(4)
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute,
	})
	if closeBill {
		return workflow.
			ExecuteActivity(ctx, activity.CloseBill, bill).
			Get(ctx, nil)
	}

		println(5)
	err = workflow.
		ExecuteActivity(ctx, activity.ChargeBill, bill, amount).
		Get(ctx, nil)
	if err != nil {
		workflow.GetLogger(ctx).Error("failed to charge bill", "amount", amount, "err", err)
	}

		println(6)
	err = workflow.
		ExecuteActivity(ctx, activity.FetchBill, bill).
		Get(ctx, &bill)
	if err != nil {
		workflow.GetLogger(ctx).Error("failed to fetch bill", "err", err)
	}

	return &workflow.ContinueAsNewError{}
}
