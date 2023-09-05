package activity

import (
	"context"

	"encore.app/bill/repository"
)

func CreateBill(ctx context.Context, customerID int) (int, error) {
	stmt := repository.DB.QueryRow(
		ctx,
		"INSERT INTO bills (customer) VALUES ($1) RETURNING id",
		customerID,
	)
	if stmt.Err() != nil {
		return 0, stmt.Err()
	}

	var billID int
	return billID, stmt.Scan(&billID)
}
