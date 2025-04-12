package services

import (
	"github.com/sMARCHz/go-secretaria-finance/internal/core/domain"
	"github.com/sMARCHz/go-secretaria-finance/internal/core/dto"
	"github.com/sMARCHz/go-secretaria-finance/internal/core/errors"
	"github.com/sMARCHz/go-secretaria-finance/internal/core/repository"
)

func (f *financeService) Withdraw(req *dto.TransactionRequest) (*dto.TransactionResponse, *errors.AppError) {
	categoryID, err := f.repository.GetCategoryIDByAbbrNameAndTransactionType(req.Category, repository.TransactionTypeWithdraw) // TODO: Add unique index
	if err != nil {
		return nil, err
	}

	account, err := f.repository.GetAccountByName(req.AccountName)
	if err != nil {
		return nil, err
	}

	if account.Balance < req.Amount {
		return nil, errors.UnprocessableEntityServerError("balance can't be less than the withdrawal amount")
	}

	transaction := &domain.TransactionInput{
		AccountID:   account.AccountID,
		CategoryID:  categoryID,
		Description: req.Description,
		Amount:      -req.Amount,
	}
	account, err = f.repository.Withdraw(transaction)
	if err != nil {
		return nil, err
	}

	return account.ToTransactionResponseDto(), nil
}
