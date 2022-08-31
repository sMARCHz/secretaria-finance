package repository

import (
	"github.com/sMARCHz/go-secretaria-finance/internal/core/domain"
	"github.com/sMARCHz/go-secretaria-finance/internal/core/errors"
)

type FinanceRepository interface {
	Withdraw(domain.Transaction) (domain.Account, *errors.AppError)
	Deposit(domain.Transaction) (domain.Account, *errors.AppError)
	Transfer(domain.Transfer) (domain.Account, *errors.AppError)
	GetAllAccountBalance() ([]domain.Account, *errors.AppError)

	GetAccountIDByName(string) (int, *errors.AppError)
	GetCategoryIDByAbbrNameAndTransactionType(string, string) (int, *errors.AppError)
}
