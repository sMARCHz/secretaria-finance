package domain

import (
	"testing"
	"time"

	"github.com/sMARCHz/go-secretaria-finance/internal/core/dto"
	"github.com/stretchr/testify/assert"
)

func TestToTransactionResponseDto(t *testing.T) {
	account := &Account{
		AccountID: 1,
		Name:      "john",
		Balance:   100,
		Currency:  "THB",
		CreatedAt: time.Now(),
	}

	res := account.ToTransactionResponseDto()

	expected := &dto.TransactionResponse{
		AccountName: account.Name,
		Balance:     account.Balance,
	}
	assert.Equal(t, expected, res)
}

func TestToTransferResponseDto(t *testing.T) {
	account := &Account{
		AccountID: 1,
		Name:      "john",
		Balance:   100,
		Currency:  "THB",
		CreatedAt: time.Now(),
	}

	res := account.ToTransferResponseDto()

	expected := &dto.TransferResponse{
		FromAccountName:    account.Name,
		FromAccountBalance: account.Balance,
	}
	assert.Equal(t, expected, res)
}

func TestToBalanceResponseDto(t *testing.T) {
	account := &Account{
		AccountID: 1,
		Name:      "john",
		Balance:   100,
		Currency:  "THB",
		CreatedAt: time.Now(),
	}

	res := account.ToBalanceResponseDto()

	expected := &dto.BalanceResponse{
		AccountName: account.Name,
		Balance:     account.Balance,
	}
	assert.Equal(t, expected, res)
}
