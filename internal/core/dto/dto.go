package dto

import (
	"net/http"
	"time"

	"github.com/sMARCHz/go-secretaria-finance/internal/adapters/driving/grpc/pb"
)

type TransactionRequest struct {
	AccountName string  `json:"account_name"`
	Category    string  `json:"category"`
	Description string  `json:"description"`
	Amount      float64 `json:"amount"`
}

type TransactionResponse struct {
	AccountName string  `json:"account_name"`
	Balance     float64 `json:"balance"`
}

func (t *TransactionResponse) ToProto() *pb.TransactionResponse {
	return &pb.TransactionResponse{
		Status:      http.StatusOK,
		AccountName: t.AccountName,
		Balance:     t.Balance,
	}
}

type TransferRequest struct {
	FromAccountName string  `json:"from_account"`
	ToAccountName   string  `json:"to_account"`
	Description     string  `json:"description"`
	Amount          float64 `json:"amount"`
}

type TransferResponse struct {
	FromAccountName    string  `json:"from_account"`
	FromAccountBalance float64 `json:"balance"`
}

func (t *TransferResponse) ToProto() *pb.TransferResponse {
	return &pb.TransferResponse{
		Status:          http.StatusOK,
		FromAccountName: t.FromAccountName,
		Balance:         t.FromAccountBalance,
	}
}

type BalanceResponse struct {
	AccountName string  `json:"account_name"`
	Balance     float64 `json:"balance"`
}

func (b *BalanceResponse) ToProto() *pb.AccountBalance {
	return &pb.AccountBalance{
		AccountName: b.AccountName,
		Balance:     b.Balance,
	}
}

type GetOverviewStatementRequest struct {
	From time.Time
	To   time.Time
}

type GetOverviewStatementResponse struct {
	Revenue *OverviewStatementSection
	Expense *OverviewStatementSection
	Profit  float64
}

func (g *GetOverviewStatementResponse) ToProto() *pb.OverviewStatementResponse {
	return &pb.OverviewStatementResponse{
		Status:  http.StatusOK,
		Revenue: g.Revenue.ToProto(),
		Expense: g.Expense.ToProto(),
		Profit:  g.Profit,
	}
}

type OverviewStatementSection struct {
	Total   float64
	Entries []*CategorizedEntry
}

func (o *OverviewStatementSection) ToProto() *pb.OverviewStatementSection {
	entries := make([]*pb.CategorizedEntry, len(o.Entries))
	for i, v := range o.Entries {
		entries[i] = v.ToProto()
	}
	return &pb.OverviewStatementSection{
		Total:   o.Total,
		Entries: entries,
	}
}

// TODO: Fix name
type CategorizedEntry struct {
	Category string
	Amount   float64
}

func (c *CategorizedEntry) ToProto() *pb.CategorizedEntry {
	return &pb.CategorizedEntry{
		Category: c.Category,
		Amount:   c.Amount,
	}
}
