package test

import (
	"database/sql"
	"testing"

	"github.com/madjiebimaa/go-random-quotes/model/domain"
	"github.com/madjiebimaa/go-random-quotes/model/web"
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
