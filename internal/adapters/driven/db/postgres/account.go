package postgres

import (
	"database/sql"

	"github.com/sMARCHz/go-secretaria-finance/internal/core/domain"
	"github.com/sMARCHz/go-secretaria-finance/internal/core/errors"
	"github.com/sMARCHz/go-secretaria-finance/pkg/logger"
)

func (f *financeRepository) GetAccountByName(name string) (*domain.Account, *errors.AppError) {
	account := domain.Account{
		Name: name,
	}
	err := f.db.Get(&account, "SELECT account_id, balance, currency, created_at FROM accounts WHERE name = $1 LIMIT 1", name) // TODO: Add unique index
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.NotFoundError("account not found")
		}

		logger.Error("failed to get accountID: ", err)
		return nil, errors.InternalServerError("failed to get accountID")
	}
	return &account, nil
}
