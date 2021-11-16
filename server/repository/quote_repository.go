package repository

import (
	"context"
	"database/sql"

	"github.com/madjiebimaa/go-random-quotes/model/domain"
)

type QuoteRepository interface {
	FindById(ctx context.Context, tx *sql.Tx, quoteId string) domain.Quote
}
