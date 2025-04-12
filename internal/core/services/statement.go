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
		logger.Error("failed to load time location: ", err)
		return nil, errors.InternalServerError("failed to load time location")
	}

	from := time.Date(req.From.Year(), req.From.Month(), req.From.Day(), 0, 0, 0, 0, loc)
	to := time.Date(req.To.Year(), req.To.Month(), req.To.Day(), 23, 59, 59, 999999999, loc)

	entries, appErr := f.repository.GetEntryByDaterange(from, to)
	if appErr != nil {
		return nil, appErr
	}

	// Calculate profit and split the entries 2 groups (revenue,expense)
	financialSummary := calculateFinancialSummary(entries)
	res := &dto.GetOverviewStatementResponse{
		Profit: financialSummary.profit,
		Revenue: &dto.OverviewStatementSection{
			Total: financialSummary.totalRevenue,
		},
		Expense: &dto.OverviewStatementSection{
			Total: financialSummary.totalExpense,
		},
	}

	for statementType, entries := range financialSummary.statements {
		categorizedEntry := sumByCategory(entries)
		switch statementType {
		case "revenue":
			res.Revenue.Entries = categorizedEntry
		case "expense":
			res.Expense.Entries = categorizedEntry
		}
	}

	return res, nil
}

type financialSummary struct {
	profit       float64
	totalRevenue float64
	totalExpense float64
	statements   map[string][]*domain.EntryWithCategory
}

func calculateFinancialSummary(entries []*domain.EntryWithCategory) *financialSummary {
	profit := 0.0
	totalRevenue := 0.0
	totalExpense := 0.0
	statements := map[string][]*domain.EntryWithCategory{
		"revenue": {},
		"expense": {},
	}

	for _, entry := range entries {
		profit += entry.Amount
		if entry.Amount > 0 {
			totalRevenue += entry.Amount
			statements["revenue"] = append(statements["revenue"], entry)
		} else {
			entry.Amount = -entry.Amount // Make it positive for displaying purpose
			totalExpense += entry.Amount
			statements["expense"] = append(statements["expense"], entry)
		}
	}

	return &financialSummary{
		profit:       profit,
		totalRevenue: totalRevenue,
		totalExpense: totalExpense,
		statements:   statements,
	}
}

func sumByCategory(entries []*domain.EntryWithCategory) []*dto.CategorizedEntry {
	summary := make(map[int]*dto.CategorizedEntry)
	for _, entry := range entries {
		if categorizedEntry, exist := summary[entry.CategoryID]; exist {
			categorizedEntry.Amount += entry.Amount
		} else {
			summary[entry.CategoryID] = &dto.CategorizedEntry{
				CategoryName: entry.Name,
				Amount:       entry.Amount,
			}
		}
	}

	res := make([]*dto.CategorizedEntry, 0)
	for _, v := range summary {
		res = append(res, v)
	}

	return res
}
