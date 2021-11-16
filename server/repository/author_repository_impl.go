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

func (repository *AuthorRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, authorId string) domain.Author {
	SQL := "SELECT id, name, link, bio, description, quote_count FROM author WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL, authorId)
	helper.PanicIfError(err)
	defer rows.Close()

	var author domain.Author
	if rows.Next() {
		rows.Scan(&author.Id, &author.Name, &author.Link, &author.Bio, &author.Description, &author.QuoteCount)
	}

	return author
}
