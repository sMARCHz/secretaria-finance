package postgres

import (
	"time"

	"github.com/sMARCHz/go-secretaria-finance/internal/core/domain"
	"github.com/sMARCHz/go-secretaria-finance/internal/core/errors"
	"github.com/sMARCHz/go-secretaria-finance/pkg/logger"
)

func (f *financeRepository) GetEntryByDaterange(from time.Time, to time.Time) ([]*domain.EntryWithCategory, *errors.AppError) {
	var entries []*domain.EntryWithCategory
	query := `SELECT 
		e.entry_id "entry.entry_id",
		e.account_id "entry.account_id",
		e.amount "entry.amount",
		e.description "entry.description",
		e.created_at "entry.created_at",
		c.category_id "category.category_id",
		c.name "category.name"
	FROM entries e
	INNER JOIN categories c
	ON e.category_id = c.category_id 
	WHERE (e.created_at BETWEEN $1 AND $2)
	AND c."transaction_type" <> 'TRANSFER'`

	err := f.db.Select(&entries, query, from, to)
	if err != nil {
		logger.Error("failed to get entry by time range: ", err)
		return nil, errors.InternalServerError("failed to get entry by time range")
	}
	return entries, nil
}
