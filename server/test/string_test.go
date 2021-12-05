package test

import (
	"database/sql"
	"testing"

	"github.com/madjiebimaa/go-random-quotes/helpers"
	"github.com/madjiebimaa/go-random-quotes/models/domain"
	"github.com/madjiebimaa/go-random-quotes/models/web"
	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {

	t.Run("Exist value", func(t *testing.T) {
		request := web.AuthorCreateRequest{
			Bio: "Pemuda ganteng",
		}

		author := domain.Author{
			Bio: sql.NullString{String: request.Bio, Valid: request.Bio == ""},
		}

		assert.True(t, request.Bio == author.Bio.String)
		assert.False(t, author.Bio.Valid)
	})

	t.Run("Not Exist value", func(t *testing.T) {
		request := web.AuthorCreateRequest{
			Bio: "",
		}

		author := domain.Author{
			Bio: sql.NullString{String: request.Bio, Valid: request.Bio == ""},
		}

		assert.True(t, true, request.Bio == author.Bio.String)
		assert.True(t, author.Bio.Valid)
	})
}

func TestSlug(t *testing.T) {
	type Test struct {
		Input    string
		Expected string
	}

	testTable := []Test{
		{Input: "Muhammad Adjie Bimaditya", Expected: "muhammad-adjie-bimaditya"},
		{Input: "Relung", Expected: "relung"},
		{Input: "BUNGA DISANA INDAH SEKALI", Expected: "bunga-disana-indah-sekali"},
	}

	for _, test := range testTable {
		slug := helpers.ToSlugFromAuthorName(test.Input)
		assert.Equal(t, test.Expected, slug)
	}
}
