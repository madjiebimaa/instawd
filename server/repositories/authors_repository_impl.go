package repositories

import (
	"context"
	"database/sql"

	"github.com/madjiebimaa/go-random-quotes/helpers"
	"github.com/madjiebimaa/go-random-quotes/models/domain"
)

type AuthorsRepositoryImpl struct{}

func NewAuthorsRepository() AuthorsRepository {
	return &AuthorsRepositoryImpl{}
}

func (repository AuthorsRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, author domain.Author) domain.Author {
	SQL := "INSERT INTO authors (id, name, link, bio, description, quote_count, slug) VALUES (?, ?, ?, ?, ?, ?, ?)"
	_, err := tx.ExecContext(ctx, SQL, author.Id, author.Name, author.Link, author.Bio, author.Description, author.QuoteCount, author.Slug)
	helpers.PanicIfError(err)

	return author
}

func (repository *AuthorsRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, authorId string) domain.Author {
	SQL := "SELECT id, name, link, bio, description, quote_count, slug FROM authors WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL, authorId)
	helpers.PanicIfError(err)
	defer rows.Close()

	var author domain.Author
	if rows.Next() {
		rows.Scan(&author.Id, &author.Name, &author.Link, &author.Bio, &author.Description, &author.QuoteCount, &author.Slug)
	}

	return author
}

func (repository AuthorsRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Author {
	SQL := "SELECT id, name, link, bio, description, quote_count, slug FROM authors"
	rows, err := tx.QueryContext(ctx, SQL)
	helpers.PanicIfError(err)
	defer rows.Close()

	var authors []domain.Author
	for rows.Next() {
		var author domain.Author
		rows.Scan(&author.Id, &author.Name, &author.Link, &author.Bio, &author.Description, &author.QuoteCount, &author.Slug)
		authors = append(authors, author)
	}

	return authors
}

func (repository *AuthorsRepositoryImpl) FindBySlug(ctx context.Context, tx *sql.Tx, authorSlug string) domain.Author {
	SQL := "SELECT id, name, link, bio, description, quote_count, slug FROM authors WHERE slug = ?"
	rows, err := tx.QueryContext(ctx, SQL, authorSlug)
	helpers.PanicIfError(err)
	defer rows.Close()

	var author domain.Author
	if rows.Next() {
		rows.Scan(&author.Id, &author.Name, &author.Link, &author.Bio, &author.Description, &author.QuoteCount, &author.Slug)
	}

	return author
}
