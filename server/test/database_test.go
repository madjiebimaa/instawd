package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/madjiebimaa/go-random-quotes/database"
	"github.com/madjiebimaa/go-random-quotes/helpers"
	"github.com/madjiebimaa/go-random-quotes/models/domain"
)

func TestTransaction(t *testing.T) {

	ctx := context.Background()

	tx, err := database.NewMySQL().Begin()
	helpers.PanicIfError(err)

	// do transaction

	rows, err := tx.QueryContext(ctx, "SELECT id FROM author")
	helpers.PanicIfError(err)
	defer rows.Close()

	var authors []domain.Author
	for rows.Next() {
		var author domain.Author
		rows.Scan(&author.Id)
		authors = append(authors, author)
	}

	fmt.Println(authors)

	rows, err = tx.QueryContext(ctx, "SELECT name FROM author")
	helpers.PanicIfError(err)
	defer rows.Close()

	var authorsName []domain.Author
	for rows.Next() {
		var author domain.Author
		rows.Scan(&author.Name)
		authorsName = append(authorsName, author)
	}

	fmt.Println(authorsName)

	err = tx.Commit()
	helpers.PanicIfError(err)
}
