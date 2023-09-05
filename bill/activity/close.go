package activity

import (
	"context"

	"encore.app/bill/repository"
)

// CloseBill activity close a bill for the charges.
func CloseBill(ctx context.Context, bill repository.Bill) error {
	if err := bill.Close(ctx); err != nil {
		return err
	}
	return nil
}
