package activity

import (
	"context"

	"encore.app/bill/repository"
)

// CloseBill activity close a bill for the charges.
func CloseBill(ctx context.Context, bill repository.Bill) error {
	return bill.Close(ctx)
}
