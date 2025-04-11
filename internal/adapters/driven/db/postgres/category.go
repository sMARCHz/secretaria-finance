package postgres

import (
	"database/sql"

	"github.com/sMARCHz/go-secretaria-finance/internal/core/errors"
	"github.com/sMARCHz/go-secretaria-finance/internal/core/repository"
	"github.com/sMARCHz/go-secretaria-finance/pkg/logger"
)

func (f *financeRepository) GetCategoryIDByAbbrNameAndTransactionType(categoryAbbrName string, txnType repository.TransactionType) (int, *errors.AppError) {
	var categoryID int
	err := f.db.Get(&categoryID, "SELECT category_id FROM categories WHERE name_abbr = $1 AND transaction_type = $2 LIMIT 1", categoryAbbrName, txnType)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Errorf("category not found where abbreviation='%v', transactionType='%v'", categoryAbbrName, txnType)
			return -1, errors.NotFoundError("category not found")
		}
		logger.Error("failed to get categoryID: ", err)
		return -1, errors.InternalServerError("failed to get categoryID")
	}
	return categoryID, nil
}
