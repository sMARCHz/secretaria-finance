package services

import (
	"github.com/sMARCHz/go-secretaria-finance/internal/core/domain"
	"github.com/sMARCHz/go-secretaria-finance/internal/core/dto"
	"github.com/sMARCHz/go-secretaria-finance/internal/core/errors"
	"github.com/sMARCHz/go-secretaria-finance/internal/core/repository"
	"github.com/sMARCHz/go-secretaria-finance/pkg/logger"
)

func (f *financeService) Withdraw(req *dto.TransactionRequest) (*dto.TransactionResponse, *errors.AppError) {
	category, err := f.repository.GetCategoryByAbbrNameAndTransactionType(req.Category, repository.TransactionTypeWithdraw)
	if err != nil {
		return nil, err
	}

	account, err := f.repository.GetAccountByName(req.AccountName)
	if err != nil {
		return nil, err
	}

	if account.Balance < req.Amount {
		logger.Error("balance cannot be less than the withdrawal amount")
		return nil, errors.UnprocessableEntityServerError("balance cannot be less than the withdrawal amount")
	}

	transaction := &domain.TransactionInput{
		AccountID:   account.AccountID,
		CategoryID:  category.CategoryID,
		Description: req.Description,
		Amount:      -req.Amount,
	}
	account, err = f.repository.Withdraw(transaction)
	if err != nil {
		return nil, err
	}

	return account.ToTransactionResponseDto(), nil
}
