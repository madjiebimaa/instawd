package controller

import "github.com/gofiber/fiber/v2"

type AuthorController interface {
	Create(c *fiber.Ctx) error
	FindById(c *fiber.Ctx) error
	FindAll(c *fiber.Ctx) error
}
