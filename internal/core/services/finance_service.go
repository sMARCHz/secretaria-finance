package services

import (
	"time"

	"github.com/sMARCHz/go-secretaria-finance/internal/core/domain"
	"github.com/sMARCHz/go-secretaria-finance/internal/core/dto"
	"github.com/sMARCHz/go-secretaria-finance/internal/core/errors"
	"github.com/sMARCHz/go-secretaria-finance/internal/core/repository"
	"github.com/sMARCHz/go-secretaria-finance/pkg/logger"
)

type FinanceService interface {
	Withdraw(dto.TransactionRequest) (*dto.TransactionResponse, *errors.AppError)
	Deposit(dto.TransactionRequest) (*dto.TransactionResponse, *errors.AppError)
	Transfer(dto.TransferRequest) (*dto.TransferResponse, *errors.AppError)
	GetBalance() ([]dto.BalanceResponse, *errors.AppError)
	GetOverviewStatement(dto.GetOverviewStatementRequest) (*dto.GetOverviewStatementResponse, *errors.AppError)
	GetOverviewMonthlyStatement() (*dto.GetOverviewStatementResponse, *errors.AppError)
	GetOverviewAnnualStatement() (*dto.GetOverviewStatementResponse, *errors.AppError)
}

type financeService struct {
	repository repository.FinanceRepository
}

func NewFinanceService(repo repository.FinanceRepository) FinanceService {
	return &financeService{repo}
}

func (f *financeService) Withdraw(req dto.TransactionRequest) (*dto.TransactionResponse, *errors.AppError) {
	categoryID, err := f.repository.GetCategoryIDByAbbrNameAndTransactionType(req.Category, repository.TransactionTypeWithdraw)
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

	transaction := domain.TransactionInput{
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

func (f *financeService) Deposit(req dto.TransactionRequest) (*dto.TransactionResponse, *errors.AppError) {
	categoryID, err := f.repository.GetCategoryIDByAbbrNameAndTransactionType(req.Category, repository.TransactionTypeDeposit)
	if err != nil {
		return nil, err
	}

	account, err := f.repository.GetAccountByName(req.AccountName)
	if err != nil {
		return nil, err
	}

	transaction := domain.TransactionInput{
		AccountID:   account.AccountID,
		CategoryID:  categoryID,
		Description: req.Description,
		Amount:      req.Amount,
	}
	account, err = f.repository.Deposit(transaction)
	if err != nil {
		return nil, err
	}

	return account.ToTransactionResponseDto(), nil
}

func (f *financeService) Transfer(req dto.TransferRequest) (*dto.TransferResponse, *errors.AppError) {
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

	transfer := domain.TransferInput{
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

func (f *financeService) GetBalance() ([]dto.BalanceResponse, *errors.AppError) {
	accounts, err := f.repository.GetAllAccountBalance()
	if err != nil {
		return nil, err
	}

	responses := make([]dto.BalanceResponse, len(accounts))
	for i, v := range accounts {
		responses[i] = *v.ToBalanceResponseDto()
	}

	return responses, nil
}

func (f *financeService) GetOverviewStatement(req dto.GetOverviewStatementRequest) (*dto.GetOverviewStatementResponse, *errors.AppError) {
	loc, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		logger.Error("failed to load time location")
		return nil, errors.InternalServerError("failed to load time location")
	}

	from := time.Date(req.From.Year(), req.From.Month(), req.From.Day(), 0, 0, 0, 0, loc)
	to := time.Date(req.To.Year(), req.To.Month(), req.To.Day(), 23, 59, 59, 0, loc)

	entries, appErr := f.repository.GetEntryByDaterange(from, to)
	if appErr != nil {
		return nil, appErr
	}

	// calculate profit and split the entries 2 groups(revenue,expense)
	profit := 0.0
	totalRevenue := 0.0
	totalExpense := 0.0
	statement := map[string][]domain.Entry{
		"revenue": {},
		"expense": {},
	}
	for _, v := range entries {
		profit += v.Amount
		if v.Amount > 0 {
			totalRevenue += v.Amount
			statement["revenue"] = append(statement["revenue"], v)
		} else {
			v.Amount = -v.Amount
			totalExpense += v.Amount
			statement["expense"] = append(statement["expense"], v)
		}
	}

	// group entries by category
	revenue := dto.OverviewStatementSection{
		Total: totalRevenue,
	}
	expense := dto.OverviewStatementSection{
		Total: totalExpense,
	}
	for k, entries := range statement {
		categorizedEntry := f.groupEntriesByCategory(entries)
		if k == "revenue" {
			revenue.Entries = categorizedEntry
		} else if k == "expense" {
			expense.Entries = categorizedEntry
		}
	}

	return &dto.GetOverviewStatementResponse{
		Profit:  profit,
		Revenue: revenue,
		Expense: expense,
	}, nil
}

func (*financeService) groupEntriesByCategory(entries []domain.Entry) []dto.CategorizedEntry {
	m := make(map[string]dto.CategorizedEntry)
	for _, e := range entries {
		categorizedEntry, present := m[e.Category.Name]
		if present {
			categorizedEntry.Amount += e.Amount
			m[e.Category.Name] = categorizedEntry
		} else {
			m[e.Category.Name] = dto.CategorizedEntry{Category: e.Category.Name, Amount: e.Amount}
		}
	}

	categorizedEntries := make([]dto.CategorizedEntry, 0)
	for _, v := range m {
		categorizedEntries = append(categorizedEntries, v)
	}

	return categorizedEntries
}

func (f *financeService) GetOverviewMonthlyStatement() (*dto.GetOverviewStatementResponse, *errors.AppError) {
	today := time.Now()
	from := time.Date(today.Year(), today.Month(), 1, 0, 0, 0, 0, time.UTC)
	to := from.AddDate(0, 1, -1)
	req := dto.GetOverviewStatementRequest{
		From: from,
		To:   to,
	}
	return f.GetOverviewStatement(req)
}

func (f *financeService) GetOverviewAnnualStatement() (*dto.GetOverviewStatementResponse, *errors.AppError) {
	today := time.Now()
	from := time.Date(today.Year(), 1, 1, 0, 0, 0, 0, time.UTC)
	to := time.Date(today.Year(), 12, 31, 0, 0, 0, 0, time.UTC)
	req := dto.GetOverviewStatementRequest{
		From: from,
		To:   to,
	}
	return f.GetOverviewStatement(req)
}
