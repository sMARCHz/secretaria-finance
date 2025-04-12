package dto

import (
	"net/http"
	"testing"

	"github.com/sMARCHz/go-secretaria-finance/internal/adapters/driving/grpc/pb"
	"github.com/stretchr/testify/assert"
)

func TestTransactionResponseToProto(t *testing.T) {
	tRes := &TransactionResponse{
		AccountName: "john",
		Balance:     100,
	}

	res := tRes.ToProto()

	assert.Equal(t, int32(http.StatusOK), res.Status)
	assert.Empty(t, res.Error)
	assert.Equal(t, tRes.AccountName, res.AccountName)
	assert.Equal(t, tRes.Balance, res.Balance)
}

func TestTransferResponseToProto(t *testing.T) {
	tRes := &TransferResponse{
		FromAccountName:    "john",
		FromAccountBalance: 100,
	}

	res := tRes.ToProto()

	assert.Equal(t, int32(http.StatusOK), res.Status)
	assert.Empty(t, res.Error)
	assert.Equal(t, tRes.FromAccountName, res.FromAccountName)
	assert.Equal(t, tRes.FromAccountBalance, res.Balance)
}

func TestGetOverviewStatementResponseToProto(t *testing.T) {
	oRes := &GetOverviewStatementResponse{
		Revenue: &OverviewStatementSection{
			Total: 1000,
			Entries: []*CategorizedEntry{
				{
					CategoryName: "salary",
					Amount:       1000,
				},
			},
		},
		Expense: &OverviewStatementSection{
			Total: 900,
			Entries: []*CategorizedEntry{
				{
					CategoryName: "shopping",
					Amount:       900,
				},
			},
		},
		Profit: 100,
	}

	res := oRes.ToProto()

	expectedRevenue := &pb.OverviewStatementSection{
		Total: 1000,
		Entries: []*pb.CategorizedEntry{
			{
				Category: "salary",
				Amount:   1000,
			},
		},
	}
	expectedExpense := &pb.OverviewStatementSection{
		Total: 900,
		Entries: []*pb.CategorizedEntry{
			{
				Category: "shopping",
				Amount:   900,
			},
		},
	}
	assert.Equal(t, int32(http.StatusOK), res.Status)
	assert.Empty(t, res.Error)
	assert.Equal(t, oRes.Profit, res.Profit)
	assert.Equal(t, expectedRevenue, res.Revenue)
	assert.Equal(t, expectedExpense, res.Expense)
}

func TestOverviewStatementSectionToProto(t *testing.T) {
	section := &OverviewStatementSection{
		Total: 100,
		Entries: []*CategorizedEntry{
			{
				CategoryName: "shopping",
				Amount:       100,
			},
		},
	}

	res := section.ToProto()

	expected := []*pb.CategorizedEntry{
		{
			Category: "shopping",
			Amount:   100,
		},
	}
	assert.Equal(t, 100.0, res.Total)
	assert.Equal(t, expected, res.Entries)
}

func TestCategorizedEntryToProto(t *testing.T) {
	entry := &CategorizedEntry{
		CategoryName: "shopping",
		Amount:       100,
	}

	res := entry.ToProto()

	assert.Equal(t, entry.CategoryName, res.Category)
	assert.Equal(t, entry.Amount, res.Amount)
}
