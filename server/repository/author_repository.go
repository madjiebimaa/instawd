package repository

import (
	"context"
	"database/sql"

	"github.com/madjiebimaa/go-random-quotes/model/domain"
)

type AuthorRepository interface {
	FindById(ctx context.Context, tx *sql.Tx, authorId string) domain.Author
}
