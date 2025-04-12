package postgres

import (
	"database/sql"

	"github.com/sMARCHz/go-secretaria-finance/internal/core/domain"
	"github.com/sMARCHz/go-secretaria-finance/internal/core/errors"
	"github.com/sMARCHz/go-secretaria-finance/internal/core/repository"
	"github.com/sMARCHz/go-secretaria-finance/pkg/logger"
)

func (f *financeRepository) GetCategoryByAbbrNameAndTransactionType(categoryAbbrName string, txnType repository.TransactionType) (*domain.Category, *errors.AppError) {
	category := domain.Category{
		NameAbbriviation: categoryAbbrName,
		TransactionType:  string(txnType),
	}
	err := f.db.Get(&category, "SELECT category_id, name, created_at FROM categories WHERE name_abbr = $1 AND transaction_type = $2 LIMIT 1", categoryAbbrName, txnType) // TODO: Add unique index
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.NotFoundError("category not found")
		}

		logger.Error("failed to get category: ", err)
		return nil, errors.InternalServerError("failed to get category")
	}
	return &category, nil
}
