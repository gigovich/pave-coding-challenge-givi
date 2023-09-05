package activity

import (
	"context"

	"encore.app/bill/repository"
	"github.com/shopspring/decimal"
)

// ChargeBill activity charges a new bill in the database and returns the Bill ID.
func ChargeBill(ctx context.Context, bill repository.Bill, amount decimal.Decimal) error {
	if err := bill.Charge(ctx, amount); err != nil {
		return err
	}
	return nil
}