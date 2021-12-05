package repositories

import (
	"context"
	"database/sql"

	"github.com/madjiebimaa/go-random-quotes/models/domain"
)

type AuthorsRepository interface {
	Create(ctx context.Context, tx *sql.Tx, author domain.Author) domain.Author
	FindById(ctx context.Context, tx *sql.Tx, authorId string) domain.Author
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Author
	FindBySlug(ctx context.Context, tx *sql.Tx, authorSlug string) domain.Author
}
