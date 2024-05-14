package si_bank

// import (
// 	"context"
// 	"database/sql"
// 	"fmt"
// )

// type Store struct {
// 	*Queries
// 	db *sql.DB
// }

// func NewStore(db *sql.DB) *Store {
// 	return &Store{
// 		db:      db,
// 		Queries: New(db),
// 	}
// }

// func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
// 	tx, err := store.db.BeginTx(ctx, nil)

// 	if err != nil {
// 		return err
// 	}

// 	q := New(tx)
// 	err = fn(q)

// 	if err != nil {
// 		if rbErr := tx.Rollback(); rbErr != nil {
// 			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
// 		}

// 		return err
// 	}

// 	return tx.Commit()
// }

// type TransferTxParams struct {
// 	FromAcccountID int64 `json:"from_account_id"`
// 	ToAcccountID   int64 `json:"to_account_id"`
// 	Amount         int64 `json:"Amount"`
// }

// type TransferTxResult struct {
// 	Transfer    Transfer `json:"transfer"`
// 	FromAccount Account  `json:"from_account"`
// 	ToAccount   Account  `json:"to_account"`
// 	FromEntry   Entry    `json:"from_entry"`
// 	ToEntry     Entry    `json:"to_entry"`
// }

// func (store *Store) TransferTx(ctx context.Context, args TransferTxParams) (TransferTxResult, error) {
// 	var result TransferTxResult

// 	err := store.execTx(ctx, func(q *Queries) error {
// 		var err error

// 		result.Transfer, err = q.CreateTransfer
// 	})
// }
