package controller

import "github.com/gofiber/fiber/v2"

type QuoteController interface {
	Create(c *fiber.Ctx) error
	FindById(c *fiber.Ctx) error
	FindAll(c *fiber.Ctx) error
	FindRandom(c *fiber.Ctx) error
}
