package repository

import (
	"context"
	"database/sql"

	"github.com/madjiebimaa/go-random-quotes/helper"
	"github.com/madjiebimaa/go-random-quotes/model/domain"
)

type QuoteTagRepositoryImpl struct{}

func NewQuoteTagRepository() QuoteTagRepository {
	return &QuoteTagRepositoryImpl{}
}

func (repository *QuoteTagRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, quoteTag domain.QuoteTag) domain.QuoteTag {
	SQL := "INSERT INTO quote_tag (id, name) VALUES (?, ?)"
	_, err := tx.ExecContext(ctx, SQL, quoteTag.Id, quoteTag.Name)
	helper.PanicIfError(err)

	return quoteTag
}

func (repository *QuoteTagRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, quoteTag domain.QuoteTag) domain.QuoteTag {
	SQL := "DELETE FROM quote_tag WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, quoteTag.Id)
	helper.PanicIfError(err)

	return quoteTag
}

func (repository *QuoteTagRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, quoteTagId string) domain.QuoteTag {
	SQL := "SELECT id, name FROM quote_tag WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL, quoteTagId)
	helper.PanicIfError(err)
	defer rows.Close()

	var quoteTag domain.QuoteTag
	if rows.Next() {
		rows.Scan(&quoteTag.Id, &quoteTag.Name)
	}

	return quoteTag
}

func (repository *QuoteTagRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.QuoteTag {
	SQL := "SELECT id, name FROM quote_tag"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var quoteTags []domain.QuoteTag
	for rows.Next() {
		var quoteTag domain.QuoteTag
		rows.Scan(&quoteTag.Id, &quoteTag.Name)
		quoteTags = append(quoteTags, quoteTag)
	}
	return quoteTags
}
