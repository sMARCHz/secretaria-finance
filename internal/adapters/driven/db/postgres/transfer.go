package postgres

import (
	"github.com/sMARCHz/go-secretaria-finance/internal/core/domain"
	"github.com/sMARCHz/go-secretaria-finance/internal/core/errors"
	"github.com/sMARCHz/go-secretaria-finance/pkg/logger"
)

func (f *financeRepository) Transfer(t *domain.TransferInput) (*domain.Account, *errors.AppError) {
	tx, err := f.db.Begin()
	if err != nil {
		logger.Error("failed to begin transaction: ", err)
		return nil, errors.InternalServerError("failed to begin transaction")
	}

	// Get categoryID of TRANSFER
	var categoryID int
	if err := tx.QueryRow("SELECT category_id FROM categories WHERE name = 'transfer' AND transaction_type = 'TRANSFER' LIMIT 1").Scan(&categoryID); err != nil {
		tx.Rollback()
		logger.Error("failed to get categoryID: ", err)
		return nil, errors.InternalServerError("failed to get categoryID")
	}

	// Insert transfer
	if _, err := tx.Exec("INSERT INTO transfers(from_account_id, to_account_id, amount) VALUES($1, $2, $3)", t.FromAccountID, t.ToAccountID, t.Amount); err != nil {
		tx.Rollback()
		logger.Error("failed to insert transfers: ", err)
		return nil, errors.InternalServerError("failed to insert transfers")
	}

	// Insert entries of fromAccount and toAccount
	if _, err := tx.Exec("INSERT INTO entries(account_id, category_id, amount, description) VALUES($1, $2, $3, $4)", t.FromAccountID, categoryID, -t.Amount, t.Description); err != nil {
		tx.Rollback()
		logger.Error("failed to insert entries: ", err)
		return nil, errors.InternalServerError("failed to insert entries")
	}
	if _, err := tx.Exec("INSERT INTO entries(account_id, category_id, amount, description) VALUES($1, $2, $3, $4)", t.ToAccountID, categoryID, t.Amount, t.Description); err != nil {
		tx.Rollback()
		logger.Error("failed to insert entries: ", err)
		return nil, errors.InternalServerError("failed to insert entries")
	}

	// Update balance of fromAccount and to Account
	var account domain.Account
	if err := tx.QueryRow("UPDATE accounts SET balance = balance + $1 WHERE account_id = $2 RETURNING name, balance, currency, created_at", -t.Amount, t.FromAccountID).Scan(&account.Name, &account.Balance, &account.Currency, &account.CreatedAt); err != nil {
		tx.Rollback()
		logger.Error("failed to update from_account balance: ", err)
		return nil, errors.InternalServerError("failed to update from_account balance")
	}
	if _, err := tx.Exec("UPDATE accounts SET balance = balance + $1 WHERE account_id = $2", t.Amount, t.ToAccountID); err != nil {
		tx.Rollback()
		logger.Error("failed to update to_account balance: ", err)
		return nil, errors.InternalServerError("failed to update to_account balance")
	}

	// Commit
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		logger.Error("failed to commit transaction: ", err)
		return nil, errors.InternalServerError("failed to commit transaction")
	}

	logger.Infof("successfully transfer ฿%v from accountID=%v to accountID=%v", t.Amount, t.FromAccountID, t.ToAccountID)
	return &account, nil
}
