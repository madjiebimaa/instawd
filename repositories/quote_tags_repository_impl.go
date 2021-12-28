package repositories

import (
	"context"
	"database/sql"

	"github.com/madjiebimaa/instawd/models"
)

type QuoteTagsRepository struct{}

func NewQuoteTagsRepository() *QuoteTagsRepository {
	return &QuoteTagsRepository{}
}

func (r *QuoteTagsRepository) Create(ctx context.Context, tx *sql.Tx, quoteTag models.QuoteTag) error {
	sql := "INSERT INTO quote_tags (id, name) VALUES (?, ?)"
	if _, err := tx.ExecContext(ctx, sql, quoteTag.Id, quoteTag.Name); err != nil {
		return err
	}

	return nil
}

func (r *QuoteTagsRepository) Delete(ctx context.Context, tx *sql.Tx, quoteTagID string) error {
	sql := "DELETE FROM quote_tags WHERE id = ?"
	if _, err := tx.ExecContext(ctx, sql, quoteTagID); err != nil {
		return err
	}

	return nil
}

func (r *QuoteTagsRepository) FindById(ctx context.Context, tx *sql.Tx, quoteTagID string) (*models.QuoteTag, error) {
	sql := "SELECT id, name FROM quote_tags WHERE id = ? LIMIT 1"
	row, err := tx.QueryContext(ctx, sql, quoteTagID)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	var quoteTag models.QuoteTag
	if row.Next() {
		row.Scan(&quoteTag.Id, &quoteTag.Name)
	}

	return &quoteTag, nil
}

func (r *QuoteTagsRepository) FindByName(ctx context.Context, tx *sql.Tx, quoteTagName string) (*models.QuoteTag, error) {
	sql := "SELECT id, name FROM quote_tags WHERE name = ? LIMIT 1"
	row, err := tx.QueryContext(ctx, sql, quoteTagName)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	var quoteTag models.QuoteTag
	if row.Next() {
		row.Scan(&quoteTag.Id, &quoteTag.Name)
	}

	return &quoteTag, nil
}

func (r *QuoteTagsRepository) FindAll(ctx context.Context, tx *sql.Tx) ([]models.QuoteTag, error) {
	sql := "SELECT id, name FROM quote_tags"
	rows, err := tx.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var quoteTags []models.QuoteTag
	for rows.Next() {
		var quoteTag models.QuoteTag
		rows.Scan(&quoteTag.Id, &quoteTag.Name)
		quoteTags = append(quoteTags, quoteTag)
	}

	return quoteTags, nil
}
