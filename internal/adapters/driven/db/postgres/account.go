package postgres

import (
	"database/sql"

	"github.com/sMARCHz/go-secretaria-finance/internal/core/domain"
	"github.com/sMARCHz/go-secretaria-finance/internal/core/errors"
	"github.com/sMARCHz/go-secretaria-finance/pkg/logger"
)

func (f *financeRepository) GetAccountByName(name string) (*domain.Account, *errors.AppError) {
	var account domain.Account
	err := f.db.Get(&account, "SELECT account_id, name, balance, currency, created_at FROM accounts WHERE name = $1 LIMIT 1", name)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Errorf("account not found where name='%v'", name)
			return nil, errors.NotFoundError("account not found")
		}
		logger.Error("failed to get accountID: ", err)
		return nil, errors.InternalServerError("failed to get accountID")
	}
	return &account, nil
}
