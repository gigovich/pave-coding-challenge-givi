package bill

import (
	"context"

	"encore.app/bill/repository"
	"encore.app/bill/workflow"
	"github.com/shopspring/decimal"
)

type ChargeRequest struct {
	Amount string `json:"amount"`
}

type ChargeResponse struct {
	Charges []repository.BillCharge `json:"charges"`
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

	var bill repository.Bill
	query, err := s.client.QueryWorkflow(ctx, wID, "", workflow.QueryBill)
	if err != nil {
		return nil, err
	}
	if err := query.Get(&bill); err != nil {
		return nil, err
	}

	return &ChargeResponse{Charges: bill.Charges}, nil
}
