package domain

import (
	"database/sql"
	"time"

	"github.com/sMARCHz/go-secretaria-finance/internal/core/dto"
)

type Account struct {
	AccountID int       `db:"account_id"`
	Name      string    `db:"name"`
	Balance   float64   `db:"balance"`
	Currency  string    `db:"currency"`
	CreatedAt time.Time `db:"created_at"`
}

func (a *Account) ToTransactionResponseDto() *dto.TransactionResponse {
	return &dto.TransactionResponse{
		AccountName: a.Name,
		Balance:     a.Balance,
	}
}

func (a *Account) ToTransferResponseDto() *dto.TransferResponse {
	return &dto.TransferResponse{
		FromAccountName:    a.Name,
		FromAccountBalance: a.Balance,
	}
}

func (a *Account) ToBalanceResponseDto() *dto.BalanceResponse {
	return &dto.BalanceResponse{
		AccountName: a.Name,
		Balance:     a.Balance,
	}
}

// TODO: Fix this
type Entry struct {
	EntryID     int            `db:"entry_id"`
	AccountID   int            `db:"account_id"`
	Amount      float64        `db:"amount"`
	Description sql.NullString `db:"description"`
	CreatedAt   time.Time      `db:"created_at"`
}

type Category struct {
	CategoryID       int       `db:"category_id"`
	Name             string    `db:"name"`
	NameAbbriviation string    `db:"name_abbr"`
	TransactionType  string    `db:"transaction_type"`
	CreatedAt        time.Time `db:"created_at"`
}

type EntryWithCategory struct {
	Entry    `db:"entry"`
	Category `db:"category"`
}

type Transfer struct {
	TransferID    int       `db:"transfer_id"`
	FromAccountID int       `db:"from_account_id"`
	ToAccountID   int       `db:"to_account_id"`
	Amount        float64   `db:"amount"`
	CreatedAt     time.Time `db:"created_at"`
}

type TransactionInput struct {
	AccountID   int
	CategoryID  int
	Description string
	Amount      float64
	CreatedAt   time.Time
}

type TransferInput struct {
	FromAccountID int
	ToAccountID   int
	Amount        float64
	Description   string
	CreatedAt     time.Time
}
