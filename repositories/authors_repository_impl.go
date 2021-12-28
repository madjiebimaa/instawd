package repositories

import (
	"context"
	"database/sql"

	"github.com/madjiebimaa/instawd/models"
)

type AuthorsRepository struct{}

func NewAuthorsRepository() *AuthorsRepository {
	return &AuthorsRepository{}
}

func (r *AuthorsRepository) Create(ctx context.Context, tx *sql.Tx, author *models.Author) error {
	sql := "INSERT INTO authors (id, name, link, bio, description, quote_count, slug) VALUES (?, ?, ?, ?, ?, ?, ?)"
	if _, err := tx.ExecContext(ctx, sql, author.Id, author.Name, author.Link, author.Bio, author.Description, author.QuoteCount, author.Slug); err != nil {
		return err
	}

	return nil
}

func (r *AuthorsRepository) FindById(ctx context.Context, tx *sql.Tx, authorID string) (*models.Author, error) {
	sql := "SELECT id, name, link, bio, description, quote_count, slug FROM authors WHERE id = ? LIMIT 1"
	row, err := tx.QueryContext(ctx, sql, authorID)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	var author models.Author
	if row.Next() {
		row.Scan(&author.Id, &author.Name, &author.Link, &author.Bio, &author.Description, &author.QuoteCount, &author.Slug)
	}

	return &author, nil
}

func (r *AuthorsRepository) FindAll(ctx context.Context, tx *sql.Tx) ([]models.Author, error) {
	sql := "SELECT id, name, link, bio, description, quote_count, slug FROM authors"
	rows, err := tx.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var authors []models.Author
	for rows.Next() {
		var author models.Author
		rows.Scan(&author.Id, &author.Name, &author.Link, &author.Bio, &author.Description, &author.QuoteCount, &author.Slug)
		authors = append(authors, author)
	}

	return authors, nil
}

func (r *AuthorsRepository) FindBySlug(ctx context.Context, tx *sql.Tx, authorSlug string) (*models.Author, error) {
	sql := "SELECT id, name, link, bio, description, quote_count, slug FROM authors WHERE slug = ? LIMIT 1"
	row, err := tx.QueryContext(ctx, sql, authorSlug)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	var author models.Author
	if row.Next() {
		row.Scan(&author.Id, &author.Name, &author.Link, &author.Bio, &author.Description, &author.QuoteCount, &author.Slug)
	}

	return &author, nil
}
