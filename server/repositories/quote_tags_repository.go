package repositories

import (
	"context"
	"database/sql"

	"github.com/madjiebimaa/go-random-quotes/models/domain"
)

type QuoteTagsRepository interface {
	Create(ctx context.Context, tx *sql.Tx, quoteTag domain.QuoteTag) domain.QuoteTag
	Delete(ctx context.Context, tx *sql.Tx, quoteTag domain.QuoteTag) domain.QuoteTag
	FindById(ctx context.Context, tx *sql.Tx, quoteTagId string) domain.QuoteTag
	FindAll(ctx context.Context, tx *sql.Tx) []domain.QuoteTag
}
