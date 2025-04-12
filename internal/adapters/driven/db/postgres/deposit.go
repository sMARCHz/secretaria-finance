package postgres

import (
	"github.com/sMARCHz/go-secretaria-finance/internal/core/domain"
	"github.com/sMARCHz/go-secretaria-finance/internal/core/errors"
	"github.com/sMARCHz/go-secretaria-finance/pkg/logger"
)

func (f *financeRepository) Deposit(t *domain.TransactionInput) (*domain.Account, *errors.AppError) {
	tx, err := f.db.Begin()
	if err != nil {
		logger.Error("failed to begin transaction: ", err)
		return nil, errors.InternalServerError("failed to begin transaction")
	}

	// Insert entry
	query := "INSERT INTO entries(account_id, category_id, amount, description) VALUES($1, $2, $3, $4)"
	if _, err := tx.Exec(query, t.AccountID, t.CategoryID, t.Amount, t.Description); err != nil {
		tx.Rollback()
		logger.Error("failed to insert entries: ", err)
		return nil, errors.InternalServerError("failed to insert entries")
	}
	// Update the account's balance
	var account domain.Account
	err = tx.QueryRow("UPDATE accounts SET balance = balance + $1 WHERE account_id = $2 RETURNING name, balance, currency, created_at", t.Amount, t.AccountID).Scan(&account.Name, &account.Balance, &account.Currency, &account.CreatedAt)
	if err != nil {
		tx.Rollback()
		logger.Error("failed to update balance of account: ", err)
		return nil, errors.InternalServerError("failed to update balance of account")
	}

	// Commit
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		logger.Error("failed to commit transaction: ", err)
		return nil, errors.InternalServerError("failed to commit transaction")
	}

	logger.Infof("successfully deposit ฿%v to accountID=%v", t.Amount, t.AccountID)
	return &account, nil
}
