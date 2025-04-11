package repository

import (
	"time"

	"github.com/sMARCHz/go-secretaria-finance/internal/core/domain"
	"github.com/sMARCHz/go-secretaria-finance/internal/core/errors"
)

type TransactionType string

const (
	TransactionTypeWithdraw TransactionType = "WITHDRAW"
	TransactionTypeDeposit  TransactionType = "DEPOSIT"
)

type FinanceRepository interface {
	Withdraw(domain.TransactionInput) (*domain.Account, *errors.AppError)
	Deposit(domain.TransactionInput) (*domain.Account, *errors.AppError)
	Transfer(domain.TransferInput) (*domain.Account, *errors.AppError)
	GetAllAccountBalance() ([]domain.Account, *errors.AppError)
	GetEntryByDaterange(From time.Time, To time.Time) ([]domain.Entry, *errors.AppError)

	GetAccountByName(name string) (*domain.Account, *errors.AppError)
	GetCategoryIDByAbbrNameAndTransactionType(abbrName string, txnType TransactionType) (int, *errors.AppError)
}
