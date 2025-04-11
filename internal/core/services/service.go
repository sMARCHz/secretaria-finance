package services

import (
	"github.com/sMARCHz/go-secretaria-finance/internal/core/dto"
	"github.com/sMARCHz/go-secretaria-finance/internal/core/errors"
	"github.com/sMARCHz/go-secretaria-finance/internal/core/repository"
)

type FinanceService interface {
	Withdraw(*dto.TransactionRequest) (*dto.TransactionResponse, *errors.AppError)
	Deposit(*dto.TransactionRequest) (*dto.TransactionResponse, *errors.AppError)
	Transfer(*dto.TransferRequest) (*dto.TransferResponse, *errors.AppError)
	GetBalance() ([]*dto.BalanceResponse, *errors.AppError)
	GetOverviewStatement(*dto.GetOverviewStatementRequest) (*dto.GetOverviewStatementResponse, *errors.AppError)
	GetOverviewMonthlyStatement() (*dto.GetOverviewStatementResponse, *errors.AppError)
	GetOverviewAnnualStatement() (*dto.GetOverviewStatementResponse, *errors.AppError)
}

type financeService struct {
	repository repository.FinanceRepository
}

func NewFinanceService(repo repository.FinanceRepository) FinanceService {
	return &financeService{repo}
}
