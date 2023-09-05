package workflow

import (
	"time"

	"encore.app/bill/activity"
	"go.temporal.io/sdk/workflow"
)

type CreateBillParam struct {
	CustomerID          int
	TimePeriodInSeconds uint
}

func CreateBill(ctx workflow.Context, param CreateBillParam) (int, error) {
	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 1,
	}

	ctx = workflow.WithActivityOptions(ctx, options)

	var billID int
	err := workflow.
		ExecuteActivity(ctx, activity.CreateBill, param.CustomerID).
		Get(ctx, &billID)
	if err != nil {
		return err
	}

	err = workflow.SetQueryHandler(ctx, "query", func() (int, error) {
		return billID, nil
	})
	if err != nil {
		return nil
	}

	return billID, nil
}
