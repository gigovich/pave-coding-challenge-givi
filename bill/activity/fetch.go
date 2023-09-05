package activity

import (
	"context"

	"encore.app/bill/repository"
)

// FetchBill activity fetches a bill from the database and returns the Bill ID.
func FetchBill(ctx context.Context, bill repository.Bill) (repository.Bill, error) {
	if err := (&bill).Fetch(ctx); err != nil {
		return bill, err
	}
	return bill, nil
}
