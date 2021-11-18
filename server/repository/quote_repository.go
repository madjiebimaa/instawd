package repository

import (
	"context"
	"database/sql"

	"github.com/madjiebimaa/go-random-quotes/model/domain"
)

type QuoteRepository interface {
	Create(ctx context.Context, tx *sql.Tx, quote domain.Quote) domain.Quote
	FindById(ctx context.Context, tx *sql.Tx, quoteId string) domain.Quote
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Quote
	FindRandom(ctx context.Context, tx *sql.Tx) domain.Quote
	FindByAuthorId(ctx context.Context, tx *sql.Tx, authorId string) []domain.Quote
}
