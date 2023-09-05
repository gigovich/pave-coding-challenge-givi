package bill

import (
	"context"

	"encore.app/bill/workflow"
	"github.com/shopspring/decimal"
)

type ChargeRequest struct {
	Amount string `json:"amount"`
}

type ChargeResponse struct {
	Message string `json:"message"`
}

//encore:api public method=POST path=/bill/charge/:billID
func (s *Service) Charge(
	ctx context.Context, billID int, request *ChargeRequest,
) (*ChargeResponse, error) {
	wID := workflow.GetChargeWorkflowID(billID)

	amount, err := decimal.NewFromString(request.Amount)
	if err != nil {
		return nil, err
	}

	err = s.client.SignalWorkflow(ctx, wID, "", workflow.SignalChargeBill, amount)
	if err != nil {
		return nil, err
	}

	return &ChargeResponse{Message: "charged"}, nil
}
