package repository

import (
	"context"
	"database/sql"

	"github.com/madjiebimaa/go-random-quotes/helper"
	"github.com/madjiebimaa/go-random-quotes/model/domain"
)

type QuoteRepositoryImpl struct{}

func NewQuoteRepository() QuoteRepository {
	return &QuoteRepositoryImpl{}
}

func (repository *QuoteRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, quoteId string) domain.Quote {
	SQL := "SELECT id, content, author_id FROM quote WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL, quoteId)
	helper.PanicIfError(err)
	defer rows.Close()

	var quote domain.Quote
	if rows.Next() {
		rows.Scan(&quote.Id, &quote.Content, &quote.AuthorId)
	}

	return quote
}
