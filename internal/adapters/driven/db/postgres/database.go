package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/sMARCHz/go-secretaria-finance/internal/core/repository"
)

type financeRepository struct {
	db *sqlx.DB
}

func NewFinanceRepository(db *sqlx.DB) repository.FinanceRepository {
	return &financeRepository{
		db: db,
	}
}
