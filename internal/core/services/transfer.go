package services

import (
	"github.com/sMARCHz/go-secretaria-finance/internal/core/domain"
	"github.com/sMARCHz/go-secretaria-finance/internal/core/dto"
	"github.com/sMARCHz/go-secretaria-finance/internal/core/errors"
)

func (f *financeService) Transfer(req *dto.TransferRequest) (*dto.TransferResponse, *errors.AppError) {
	fromAccount, err := f.repository.GetAccountByName(req.FromAccountName)
	if err != nil {
		return nil, err
	}

	toAccount, err := f.repository.GetAccountByName(req.ToAccountName)
	if err != nil {
		return nil, err
	}

	if fromAccount.Balance < req.Amount {
		return nil, errors.UnprocessableEntityServerError("transferer's balance can't be less than the transfer amount")
	}

	transfer := &domain.TransferInput{
		FromAccountID: fromAccount.AccountID,
		ToAccountID:   toAccount.AccountID,
		Amount:        req.Amount,
		Description:   req.Description,
	}
	fromAccount, err = f.repository.Transfer(transfer)
	if err != nil {
		return nil, err
	}

	return fromAccount.ToTransferResponseDto(), nil
}
