package repositories

import (
	"context"
	"database/sql"

	"github.com/madjiebimaa/go-random-quotes/helpers"
	"github.com/madjiebimaa/go-random-quotes/models/domain"
)

type QuoteTagsRepositoryImpl struct{}

func NewQuoteTagsRepository() QuoteTagsRepository {
	return &QuoteTagsRepositoryImpl{}
}

func (repository *QuoteTagsRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, quoteTag domain.QuoteTag) domain.QuoteTag {
	SQL := "INSERT INTO quote_tags (id, name) VALUES (?, ?)"
	_, err := tx.ExecContext(ctx, SQL, quoteTag.Id, quoteTag.Name)
	helpers.PanicIfError(err)

	return quoteTag
}

func (repository *QuoteTagsRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, quoteTag domain.QuoteTag) domain.QuoteTag {
	SQL := "DELETE FROM quote_tags WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, quoteTag.Id)
	helpers.PanicIfError(err)

	return quoteTag
}

func (repository *QuoteTagsRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, quoteTagId string) domain.QuoteTag {
	SQL := "SELECT id, name FROM quote_tags WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL, quoteTagId)
	helpers.PanicIfError(err)
	defer rows.Close()

	var quoteTag domain.QuoteTag
	if rows.Next() {
		rows.Scan(&quoteTag.Id, &quoteTag.Name)
	}

	return quoteTag
}

func (repository *QuoteTagsRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.QuoteTag {
	SQL := "SELECT id, name FROM quote_tags"
	rows, err := tx.QueryContext(ctx, SQL)
	helpers.PanicIfError(err)
	defer rows.Close()

	var quoteTags []domain.QuoteTag
	for rows.Next() {
		var quoteTag domain.QuoteTag
		rows.Scan(&quoteTag.Id, &quoteTag.Name)
		quoteTags = append(quoteTags, quoteTag)
	}
	return quoteTags
}
