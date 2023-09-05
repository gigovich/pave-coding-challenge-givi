package activity

import (
	"context"

	"encore.app/bill/repository"
)

// CreateBill activity creates a new bill in the database and returns the Bill ID.
func CreateBill(ctx context.Context, bill repository.Bill) (int, error) {
	created, err := bill.Create(ctx)
	if err != nil {
		return 0, err
	}
	return created.ID, nil
}
