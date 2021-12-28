package main

import (
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"

	"github.com/madjiebimaa/instawd/database"
	"github.com/madjiebimaa/instawd/routes"
)

func main() {
	app := fiber.New()
	db := database.NewMySQL()
	validate := validator.New()

	routes.NewAuthorsRoute(app, db, validate)
	routes.NewQuotesRoute(app, db, validate)
	routes.NewQuoteTagsRoute(app, db, validate)

	app.Listen(":3000")
}
