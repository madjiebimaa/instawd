package repositories

import (
	"context"
	"database/sql"
	"math/rand"

	"github.com/madjiebimaa/go-random-quotes/helpers"
	"github.com/madjiebimaa/go-random-quotes/models/domain"
	"github.com/madjiebimaa/go-random-quotes/models/web"
)

type QuotesRepositoryImpl struct{}

func NewQuotesRepository() QuotesRepository {
	return &QuotesRepositoryImpl{}
}

func (repository *QuotesRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, quote domain.Quote) domain.Quote {
	SQL := "INSERT INTO quotes (id, author_id, content) VALUES (?, ?, ?)"
	_, err := tx.ExecContext(ctx, SQL, quote.Id, quote.AuthorId, quote.Content)
	helpers.PanicIfError(err)

	return quote
}

func (repository *QuotesRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, quoteId string) domain.Quote {
	SQL := "SELECT id, author_id, content FROM quotes WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL, quoteId)
	helpers.PanicIfError(err)
	defer rows.Close()

	var quote domain.Quote
	if rows.Next() {
		rows.Scan(&quote.Id, &quote.AuthorId, &quote.Content)
	}

	return quote
}

func (repository *QuotesRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Quote {
	SQL := "SELECT id, author_id, content, CHAR_LENGTH(content) AS content_length FROM quotes"

	rows, err := repository.queryRowsFilter(ctx, tx, SQL)
	helpers.PanicIfError(err)
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

func (repository *QuotesRepositoryImpl) FindRandom(ctx context.Context, tx *sql.Tx) domain.Quote {
	SQL := "SELECT COUNT(id) FROM quotes"
	rows, err := tx.QueryContext(ctx, SQL)
	helpers.PanicIfError(err)

	var countQuote int
	if rows.Next() {
		rows.Scan(&countQuote)
	}
	randNum := rand.Intn(countQuote)
	rows.Close()

	SQL = "WITH quotes_numbered AS (SELECT id, author_id, content, row_number() over() AS rn FROM quotes ) SELECT id, author_id, content FROM quotes_numbered WHERE rn = ?"
	rows, err = tx.QueryContext(ctx, SQL, randNum)
	helpers.PanicIfError(err)

	var quote domain.Quote
	if rows.Next() {
		rows.Scan(&quote.Id, &quote.AuthorId, &quote.Content)
	}
	rows.Close()

	return quote
}

func (repository *QuotesRepositoryImpl) FindByAuthorId(ctx context.Context, tx *sql.Tx, authorId string) []domain.Quote {
	SQL := "SELECT id, author_id, content FROM quotes WHERE author_id = ?"
	rows, err := tx.QueryContext(ctx, SQL, authorId)
	helpers.PanicIfError(err)
	defer rows.Close()

	var quotes []domain.Quote
	for rows.Next() {
		var quote domain.Quote
		rows.Scan(&quote.Id, &quote.AuthorId, &quote.Content)
		quotes = append(quotes, quote)
	}

	return quotes
}

func (repository *QuotesRepositoryImpl) queryRowsFilter(ctx context.Context, tx *sql.Tx, SQL string) (*sql.Rows, error) {
	var filterValues []interface{}
	var filterRequest web.FilterRequest
	helpers.QueryToStruct(ctx, helpers.FILTER_REQUEST, &filterRequest)

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
