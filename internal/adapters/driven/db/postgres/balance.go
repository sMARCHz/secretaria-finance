package postgres

import (
	"github.com/sMARCHz/go-secretaria-finance/internal/core/domain"
	"github.com/sMARCHz/go-secretaria-finance/internal/core/errors"
	"github.com/sMARCHz/go-secretaria-finance/pkg/logger"
)

func (f *financeRepository) GetAllAccountBalance() ([]*domain.Account, *errors.AppError) {
	var accounts []*domain.Account
	if err := f.db.Select(&accounts, "SELECT name, balance FROM accounts"); err != nil {
		logger.Error("failed to query accounts: ", err)
		return nil, errors.InternalServerError("failed to get all balance of accounts")
	}
	return accounts, nil
}
