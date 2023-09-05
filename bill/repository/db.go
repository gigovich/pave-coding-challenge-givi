package repository

import (
	"encore.dev/storage/sqldb"
)

var DB = sqldb.NewDatabase("bill", sqldb.DatabaseConfig{
	Migrations: "./migrations",
})
