package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Store interface {
	TransferTx(ctx context.Context, arg CreateTransferParams) (TransferTxResult, error)
	Querier
	CreateUserTx(ctx context.Context, arg CreateUserTxParams) (CreateUserTxResult, error)
}

// Store provides all functions to execute db queries and transactions
type SQLStore struct {
	*Queries
	db *pgxpool.Pool
}

func NewStore(db *pgxpool.Pool) Store {
	return &SQLStore{
		Queries: New(db),
		db:      db,
	}
}
