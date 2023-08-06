package db

import (
	"context"
	// "database/sql"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Bank interface{
	Querier
	TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error)
	// CreateUserTx(ctx context.Context, arg CreateUserTxParams) (CreateUserTxResult, error)
	// VerifyEmailTx(ctx context.Context, arg VerifyEmailTxParams) (VerifyEmailTxResult, error)

}

// Bank provides all functions to execute SQL queries and transactions
type SQLBank struct {
	*Queries
	connPool *pgxpool.Pool
	
}

func NewBank(connPool *pgxpool.Pool) Bank {
	return &SQLBank{
		connPool:      connPool,
		Queries: New(connPool),
	}
}

// type SQLStore struct {
// 	connPool *pgxpool.Pool
// 	*Queries
// }

// // NewStore creates a new store
// func NewStore(connPool *pgxpool.Pool) Store {
// 	return &SQLStore{
// 		connPool: connPool,
// 		Queries:  New(connPool),
// 	}
// }



// func (bank *Bank) TransferTx(ctx context.Context, arg TransferTxParams)