package repository

import (
	"context"
	"database/sql"
	"math/rand"

	"github.com/madjiebimaa/go-random-quotes/helper"
	"github.com/madjiebimaa/go-random-quotes/model/domain"
	"github.com/madjiebimaa/go-random-quotes/model/web"
)

type QuoteRepositoryImpl struct{}

func NewQuoteRepository() QuoteRepository {
	return &QuoteRepositoryImpl{}
}

func (repository *QuoteRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, quote domain.Quote) domain.Quote {
	SQL := "INSERT INTO quote (id, author_id, content) VALUES (?, ?, ?)"
	_, err := tx.ExecContext(ctx, SQL, quote.Id, quote.AuthorId, quote.Content)
	helper.PanicIfError(err)

	return quote
}

func (repository *QuoteRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, quoteId string) domain.Quote {
	SQL := "SELECT id, author_id, content FROM quote WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL, quoteId)
	helper.PanicIfError(err)
	defer rows.Close()

	var quote domain.Quote
	if rows.Next() {
		rows.Scan(&quote.Id, &quote.AuthorId, &quote.Content)
	}

	return quote
}

func (repository *QuoteRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Quote {
	SQL := "SELECT id, author_id, content, CHAR_LENGTH(content) AS content_length FROM quote"

	rows, err := repository.queryRowsFilter(ctx, tx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var quotes []domain.Quote
	for rows.Next() {
		var quote domain.Quote
		var temp interface{}
		rows.Scan(&quote.Id, &quote.AuthorId, &quote.Content, &temp)
		quotes = append(quotes, quote)
	}

	return quotes
}

func (repository *QuoteRepositoryImpl) FindRandom(ctx context.Context, tx *sql.Tx) domain.Quote {
	SQL := "SELECT COUNT(id) FROM quote"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)

	var countQuote int
	if rows.Next() {
		rows.Scan(&countQuote)
	}
	randNum := rand.Intn(countQuote)
	rows.Close()

	SQL = "WITH quote_numbered AS (SELECT id, author_id, content, row_number() over() AS rn FROM quote ) SELECT id, author_id, content FROM quote_numbered WHERE rn = ?"
	rows, err = tx.QueryContext(ctx, SQL, randNum)
	helper.PanicIfError(err)

	var quote domain.Quote
	if rows.Next() {
		rows.Scan(&quote.Id, &quote.AuthorId, &quote.Content)
	}
	rows.Close()

	return quote
}

func (repository *QuoteRepositoryImpl) FindByAuthorId(ctx context.Context, tx *sql.Tx, authorId string) []domain.Quote {
	SQL := "SELECT id, author_id, content FROM quote WHERE author_id = ?"
	rows, err := tx.QueryContext(ctx, SQL, authorId)
	helper.PanicIfError(err)
	defer rows.Close()

	var quotes []domain.Quote
	for rows.Next() {
		var quote domain.Quote
		rows.Scan(&quote.Id, &quote.AuthorId, &quote.Content)
		quotes = append(quotes, quote)
	}

	return quotes
}

func (repository *QuoteRepositoryImpl) queryRowsFilter(ctx context.Context, tx *sql.Tx, SQL string) (*sql.Rows, error) {
	var filterValues []interface{}
	var filterRequest web.FilterRequest
	helper.QueryToStruct(ctx, helper.FILTER_REQUEST, &filterRequest)

	if filterRequest.MinLength != 0 {
		filterValues = append(filterValues, filterRequest.MinLength)
		SQL += ` HAVING content_length > ?`
	}

	if filterRequest.MaxLength != 0 {
		filterValues = append(filterValues, filterRequest.MaxLength)
		SQL += ` AND content_length < ?`
	}

	if filterRequest.Limit != 0 {
		filterValues = append(filterValues, filterRequest.Limit)
		SQL += ` LIMIT ?`
	}

	if filterRequest.Offset != 0 {
		filterValues = append(filterValues, filterRequest.Offset)
		SQL += ` OFFSET ?`
	}

	return tx.QueryContext(ctx, SQL, filterValues...)
}
