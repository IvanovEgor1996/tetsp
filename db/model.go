package db

import "github.com/google/uuid"

type Balances struct {
	tableName struct{} `pg:"balances"`

	User    uuid.UUID `pg:", pk"`
	Balance int
}
