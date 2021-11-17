package repository

import (
	"context"
	"database/sql"

	"github.com/madjiebimaa/go-random-quotes/helper"
	"github.com/madjiebimaa/go-random-quotes/model/domain"
)

type AuthorRepositoryImpl struct{}

func NewAuthorRepository() AuthorRepository {
	return &AuthorRepositoryImpl{}
}

func (repository AuthorRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, author domain.Author) domain.Author {
	SQL := "INSERT INTO author (id, name, link, bio, description, quote_count, slug) VALUES (?, ?, ?, ?, ?, ?, ?)"
	_, err := tx.ExecContext(ctx, SQL, author.Id, author.Name, author.Link, author.Bio, author.Description, author.QuoteCount, author.Slug)
	helper.PanicIfError(err)

	return author
}

func (repository *AuthorRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, authorId string) domain.Author {
	SQL := "SELECT id, name, link, bio, description, quote_count, slug FROM author WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL, authorId)
	helper.PanicIfError(err)
	defer rows.Close()

	var author domain.Author
	if rows.Next() {
		rows.Scan(&author.Id, &author.Name, &author.Link, &author.Bio, &author.Description, &author.QuoteCount, &author.Slug)
	}

	return author
}

func (repository AuthorRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Author {
	SQL := "SELECT id, name, link, bio, description, quote_count, slug FROM author"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var authors []domain.Author
	for rows.Next() {
		var author domain.Author
		rows.Scan(&author.Id, &author.Name, &author.Link, &author.Bio, &author.Description, &author.QuoteCount, &author.Slug)
		authors = append(authors, author)
	}

	return authors
}
