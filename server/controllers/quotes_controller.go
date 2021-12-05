package controllers

import "github.com/gofiber/fiber/v2"

type QuotesController interface {
	Create(c *fiber.Ctx) error
	FindById(c *fiber.Ctx) error
	FindQuoteAndAuthor(c *fiber.Ctx) error
	FindAll(c *fiber.Ctx) error
	FindRandom(c *fiber.Ctx) error
	FindRandomAndAuthor(c *fiber.Ctx) error
}
