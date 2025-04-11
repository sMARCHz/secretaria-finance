package services

import (
	"github.com/sMARCHz/go-secretaria-finance/internal/core/dto"
	"github.com/sMARCHz/go-secretaria-finance/internal/core/errors"
)

func (f *financeService) GetBalance() ([]*dto.BalanceResponse, *errors.AppError) {
	accounts, err := f.repository.GetAllAccountBalance()
	if err != nil {
		return nil, err
	}

	responses := make([]*dto.BalanceResponse, len(accounts))
	for i, v := range accounts {
		responses[i] = v.ToBalanceResponseDto()
	}

	return responses, nil
}
