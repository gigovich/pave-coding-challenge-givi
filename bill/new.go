package bill

import (
	"context"

	"encore.app/bill/workflow"
	"encore.dev/rlog"
	"go.temporal.io/sdk/client"
)

type NewRequest struct {
	Customer            int `json:"customer"`
	TimePeriodInSeconds int `json:"time_period_in_seconds"`
}

type NewResponse struct {
	BillID int `json:"bill_id"`
}

//encore:api public path=/new
func (s *Service) New(ctx context.Context, request *NewRequest) (*NewResponse, error) {
	options := client.StartWorkflowOptions{
		ID:        "bill-workflow",
		TaskQueue: feesTaskQueue,
	}

	we, err := s.client.ExecuteWorkflow(ctx, options, workflow.CreateBill, request)
	if err != nil {
		return nil, err
	}
	rlog.Error("started workflow", "id", we.GetID(), "run_id", we.GetRunID())

	var response NewResponse
	if err := we.Get(ctx, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
