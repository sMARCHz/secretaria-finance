package services

import (
	"time"

	"github.com/sMARCHz/go-secretaria-finance/internal/core/domain"
	"github.com/sMARCHz/go-secretaria-finance/internal/core/dto"
	"github.com/sMARCHz/go-secretaria-finance/internal/core/errors"
	"github.com/sMARCHz/go-secretaria-finance/pkg/logger"
)

func (f *financeService) GetOverviewStatement(req *dto.GetOverviewStatementRequest) (*dto.GetOverviewStatementResponse, *errors.AppError) {
	loc, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		logger.Error("failed to load time location") // TODO: add error in all log
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
	statements := map[string][]*domain.Entry{
		"revenue": {},
		"expense": {},
	}
	for _, entry := range entries {
		profit += entry.Amount
		if entry.Amount > 0 {
			totalRevenue += entry.Amount
			statements["revenue"] = append(statements["revenue"], entry)
		} else {
			entry.Amount = -entry.Amount // Make it positive for displaying amount
			totalExpense += entry.Amount
			statements["expense"] = append(statements["expense"], entry)
		}
	}

	// group entries by category
	revenue := &dto.OverviewStatementSection{
		Total: totalRevenue,
	}
	expense := &dto.OverviewStatementSection{
		Total: totalExpense,
	}
	for entryType, entries := range statements {
		categorizedEntry := sumByCategory(entries)
		switch entryType {
		case "revenue":
			revenue.Entries = categorizedEntry
		case "expense":
			expense.Entries = categorizedEntry
		}
	}

	return &dto.GetOverviewStatementResponse{
		Profit:  profit,
		Revenue: revenue,
		Expense: expense,
	}, nil
}

func sumByCategory(entries []*domain.Entry) []*dto.CategorizedEntry {
	summary := make(map[string]*dto.CategorizedEntry)
	for _, entry := range entries {
		categoryName := entry.Category.Name
		if categorizedEntry, exist := summary[categoryName]; exist {
			categorizedEntry.Amount += entry.Amount // TODO: Fix this naming
		} else {
			summary[categoryName] = &dto.CategorizedEntry{
				Category: categoryName,
				Amount:   entry.Amount,
			}
		}
	}

	categorizedEntries := make([]*dto.CategorizedEntry, 0)
	for _, v := range summary {
		categorizedEntries = append(categorizedEntries, v)
	}

	return categorizedEntries
}
