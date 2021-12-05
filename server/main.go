package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

	_ "github.com/go-sql-driver/mysql"

	"github.com/madjiebimaa/go-random-quotes/database"
	"github.com/madjiebimaa/go-random-quotes/routes"
)

func main() {
	app := fiber.New()
	db := database.NewMySQL()
	validate := validator.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	routes.NewAuthorsRoute(app, db, validate)
	routes.NewQuotesRoute(app, db, validate)
	routes.NewQuoteTagsRoute(app, db, validate)

	app.Listen(":3000")
}
