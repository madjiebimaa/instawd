package main

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

	appl "github.com/madjiebimaa/go-random-quotes/app"
	"github.com/madjiebimaa/go-random-quotes/controller"
	"github.com/madjiebimaa/go-random-quotes/helper"
	"github.com/madjiebimaa/go-random-quotes/model/domain"
	"github.com/madjiebimaa/go-random-quotes/repository"
	"github.com/madjiebimaa/go-random-quotes/service"
)

type Response struct {
	Quote  domain.Quote
	Author domain.Author
}

func main() {
	app := fiber.New()
	db := appl.NewDB()
	validate := validator.New()

	quoteRepository := repository.NewQuoteRepository()
	quoteService := service.NewQuoteService(quoteRepository, db, validate)
	quoteController := controller.NewQuoteController(quoteService)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Post("/api/quotes", quoteController.Create)
	app.Get("/api/quotes", quoteController.FindAll)
	app.Get("/api/quotes/:quoteId", quoteController.FindById)
	app.Get("/api/random", quoteController.FindRandom)

	app.Get("/api/authors/:authorId", func(c *fiber.Ctx) error {
		authorId := c.Params("authorId")

		ctx := context.Background()

		tx, err := appl.NewDB().Begin()
		helper.PanicIfError(err)
		defer helper.CommitOrRollBack(tx)

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

		// var author web.AuthorRequest
		// err := c.BodyParser(&author)
		// helper.PanicIfError(err)

		// ctx := context.Background()

		// tx, err := appl.NewDB().Begin()
		// helper.PanicIfError(err)
		// defer helper.CommitOrRollBack(tx)

		// author.Id = helper.RandomString(12)

		// SQL := "INSERT INTO author (id, name, link, bio, description, quote_count) VALUES (?, ?, ?, ?, ?, ?)"
		// _, err = tx.ExecContext(ctx, SQL, author.Id, author.Name, author.Link, author.Bio, author.Description, author.QuoteCount)
		// helper.PanicIfError(err)

		// c.Status(fiber.StatusOK)
		// c.Type(fiber.MIMEApplicationJSON)
		// return c.JSON(author)
		return c.SendString("Maintenance")
	})

	app.Listen(":3000")
}
