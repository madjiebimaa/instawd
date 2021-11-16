package main

import (
	"context"

	"github.com/gofiber/fiber/v2"

	appl "github.com/madjiebimaa/go-random-quotes/app"
	"github.com/madjiebimaa/go-random-quotes/helper"
	"github.com/madjiebimaa/go-random-quotes/model/domain"
	"github.com/madjiebimaa/go-random-quotes/model/web"
)

type Response struct {
	Quote  domain.Quote
	Author domain.Author
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
	app.Get("/api/quotes/:quoteId", func(c *fiber.Ctx) error {
		quoteId := c.Params("quoteId")

		tx, err := appl.NewDB().Begin()
		helper.PanicIfError(err)
		defer helper.CommitOrRollBack(tx)

		ctx := context.Background()

		SQL := "SELECT id, content, author_id FROM quote WHERE id = ?"
		rows, err := tx.QueryContext(ctx, SQL, quoteId)
		helper.PanicIfError(err)
		defer rows.Close()

		var quote domain.Quote
		if rows.Next() {
			rows.Scan(&quote.Id, &quote.Content, &quote.AuthorId)
		}

		c.Status(fiber.StatusOK)
		c.Type(fiber.MIMEApplicationJSON)
		return c.JSON(quote)
	})
	app.Post("/api/quotes", func(c *fiber.Ctx) error {
		c.Accepts(fiber.MIMEApplicationJSON)

		var quote web.QuoteRequest
		err := c.BodyParser(&quote)
		helper.PanicIfError(err)

		c.Status(fiber.StatusOK)
		c.Type(fiber.MIMEApplicationJSON)
		return c.JSON(quote)
	})
	app.Get("/api/authors/:authorId", func(c *fiber.Ctx) error {
		authorId := c.Params("authorId")

		tx, err := appl.NewDB().Begin()
		helper.PanicIfError(err)
		defer helper.CommitOrRollBack(tx)

		ctx := context.Background()

		SQL := "SELECT id, name, link, bio, description, quote_count FROM author WHERE id = ?"
		rows, err := tx.QueryContext(ctx, SQL, authorId)
		helper.PanicIfError(err)
		defer rows.Close()

		var author domain.Author
		if rows.Next() {
			rows.Scan(&author.Id, &author.Name, &author.Link, &author.Bio, &author.Description, &author.QuoteCount)
		}

		c.Status(fiber.StatusOK)
		c.Type(fiber.MIMEApplicationJSON)
		return c.JSON(helper.ToAuthorResponse(author))
	})
	app.Post("/api/authors", func(c *fiber.Ctx) error {
		c.Accepts(fiber.MIMEApplicationJSON)

		var author web.AuthorRequest
		err := c.BodyParser(&author)
		helper.PanicIfError(err)

		tx, err := appl.NewDB().Begin()
		helper.PanicIfError(err)
		defer helper.CommitOrRollBack(tx)

		ctx := context.Background()

		author.Id = helper.RandomString(12)

		SQL := "INSERT INTO author (id, name, link, bio, description, quote_count) VALUES (?, ?, ?, ?, ?, ?)"
		_, err = tx.ExecContext(ctx, SQL, author.Id, author.Name, author.Link, author.Bio, author.Description, author.QuoteCount)
		helper.PanicIfError(err)

		c.Status(fiber.StatusOK)
		c.Type(fiber.MIMEApplicationJSON)
		return c.JSON(author)
	})

	app.Listen(":3000")
}
