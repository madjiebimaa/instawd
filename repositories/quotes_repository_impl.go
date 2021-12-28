package repositories

import (
	"context"
	"database/sql"
	"math/rand"

	"github.com/madjiebimaa/instawd/helpers"
	"github.com/madjiebimaa/instawd/models"
)

type QuotesRepository struct{}

func NewQuotesRepository() *QuotesRepository {
	return &QuotesRepository{}
}

func (r *QuotesRepository) Create(ctx context.Context, tx *sql.Tx, quote *models.Quote) error {
	sql := "INSERT INTO quotes (id, author_id, content) VALUES (?, ?, ?)"
	if _, err := tx.ExecContext(ctx, sql, quote.Id, quote.AuthorId, quote.Content); err != nil {
		return err
	}

	return nil
}

func (r *QuotesRepository) FindById(ctx context.Context, tx *sql.Tx, quoteID string) (*models.Quote, error) {
	sql := "SELECT id, author_id, content FROM quotes WHERE id = ?"
	row, err := tx.QueryContext(ctx, sql, quoteID)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	var quote models.Quote
	if row.Next() {
		row.Scan(&quote.Id, &quote.AuthorId, &quote.Content)
	}

	return &quote, err
}

func (r *QuotesRepository) queryRowsFilter(ctx context.Context, tx *sql.Tx, sql string) (*sql.Rows, error) {
	var filterValues []interface{}
	var filterRequest models.FilterRequest
	helpers.QueryToStruct(ctx, helpers.FILTER_REQUEST, &filterRequest)

	if filterRequest.MinLength != 0 {
		filterValues = append(filterValues, filterRequest.MinLength)
		sql += ` HAVING content_length > ?`
	}

	if filterRequest.MaxLength != 0 {
		filterValues = append(filterValues, filterRequest.MaxLength)
		sql += ` AND content_length < ?`
	}

	if filterRequest.Limit != 0 {
		filterValues = append(filterValues, filterRequest.Limit)
		sql += ` LIMIT ?`
	}

	if filterRequest.Offset != 0 {
		filterValues = append(filterValues, filterRequest.Offset)
		sql += ` OFFSET ?`
	}

	return tx.QueryContext(ctx, sql, filterValues...)
}

func (r *QuotesRepository) FindAll(ctx context.Context, tx *sql.Tx) ([]models.Quote, error) {
	sql := "SELECT id, author_id, content, CHAR_LENGTH(content) AS content_length FROM quotes"
	rows, err := r.queryRowsFilter(ctx, tx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var quotes []models.Quote
	for rows.Next() {
		var quote models.Quote
		var temp interface{}
		rows.Scan(&quote.Id, &quote.AuthorId, &quote.Content, &temp)
		quotes = append(quotes, quote)
	}

	return quotes, nil
}

func (r *QuotesRepository) FindRandom(ctx context.Context, tx *sql.Tx) (*models.Quote, error) {
	sql := "SELECT COUNT(id) FROM quotes"
	row, err := tx.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}

	var countQuote int
	if row.Next() {
		row.Scan(&countQuote)
	}
	randNum := rand.Intn(countQuote)
	row.Close()

	sql = "WITH quotes_numbered AS (SELECT id, author_id, content, row_number() over() AS rn FROM quotes ) SELECT id, author_id, content FROM quotes_numbered WHERE rn = ?"
	row, err = tx.QueryContext(ctx, sql, randNum)
	if err != nil {
		return nil, err
	}

	var quote models.Quote
	if row.Next() {
		row.Scan(&quote.Id, &quote.AuthorId, &quote.Content)
	}
	row.Close()

	return &quote, nil
}

func (r *QuotesRepository) FindByAuthorId(ctx context.Context, tx *sql.Tx, authorID string) ([]models.Quote, error) {
	sql := "SELECT id, author_id, content FROM quotes WHERE author_id = ?"
	rows, err := tx.QueryContext(ctx, sql, authorID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var quotes []models.Quote
	for rows.Next() {
		var quote models.Quote
		rows.Scan(&quote.Id, &quote.AuthorId, &quote.Content)
		quotes = append(quotes, quote)
	}

	return quotes, nil
}

func (r *AuthorsRepository) FindQuoteDetailByAuthor(ctx context.Context, tx *sql.Tx, quoteID string) (*models.AuthorAndQuote, error) {
	sql := `
		SELECT u.id AS user_id,
			u.name AS name,
			u.link AS link,
			q.id AS quote_id,
			q.content AS content
		FROM quotes AS q
			INNER JOIN users AS u
			ON (q.author_id = u.id)
		WHERE id = ? LIMIT 1
	`
	rows, err := tx.QueryContext(ctx, sql, quoteID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var authorQuote models.AuthorAndQuote
	if rows.Next() {
		var author models.Author
		var quote models.Quote

		rows.Scan(&author.Id, &author.Name, &author.Link, &quote.Id, &quote.Content)

		authorQuote.Author = author
		authorQuote.Quote = quote
	}

	return &authorQuote, nil
}

func (r *AuthorsRepository) FindQuotesByAuthor(ctx context.Context, tx *sql.Tx, authorID string) (*models.AuthorAndQuotes, error) {
	sql := `
		SELECT u.id AS user_id,
			u.name AS name,
			u.link AS link,
			u.bio AS bio,
			u.description AS description,
			u.quote_count AS quote_count,
			u.slug AS slug,
			q.id AS quote_id,
			q.content AS content
		FROM quotes AS q
			INNER JOIN users AS u
			ON (q.author_id = u.id)
		WHERE author_id = ?
	`
	rows, err := tx.QueryContext(ctx, sql, authorID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var authorQuotes models.AuthorAndQuotes
	for rows.Next() {
		var author models.Author
		var quote models.Quote

		rows.Scan(&author.Id, &author.Name, &author.Link, &author.Bio, &author.Description, &author.QuoteCount, &author.Slug, &quote.Id, &quote.Content)

		authorQuotes.Author = author
		authorQuotes.Quotes = append(authorQuotes.Quotes, quote)
	}

	return &authorQuotes, nil
}
