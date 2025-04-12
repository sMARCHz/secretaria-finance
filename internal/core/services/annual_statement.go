package services

import (
	"time"

	"github.com/sMARCHz/go-secretaria-finance/internal/core/dto"
	"github.com/sMARCHz/go-secretaria-finance/internal/core/errors"
)

func (f *financeService) GetOverviewAnnualStatement() (*dto.GetOverviewStatementResponse, *errors.AppError) {
	today := time.Now()
	from := time.Date(today.Year(), 1, 1, 0, 0, 0, 0, time.UTC)
	to := from.AddDate(1, 0, 0).Add(-time.Second)
	req := &dto.GetOverviewStatementRequest{
		From: from,
		To:   to,
	}
	return f.GetOverviewStatement(req)
}
