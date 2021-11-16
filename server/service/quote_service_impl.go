package service

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/madjiebimaa/go-random-quotes/repository"
)

type QuoteServiceImpl struct {
	DB              *sql.Tx
	QuoteRepository *repository.QuoteRepository
}

func NewQuoteService(db *sql.Tx, quoteRepository repository.QuoteRepository) QuoteService {
	return &QuoteServiceImpl{
		DB:              db,
		QuoteRepository: &quoteRepository,
	}
}

func (service *QuoteServiceImpl) FindQuoteAndAuthor(ctx context.Context, request *http.Request) {

}
