package service

import (
	"context"
	"net/http"
)

type QuoteService interface {
	FindQuoteAndAuthor(ctx context.Context, request *http.Request)
}
