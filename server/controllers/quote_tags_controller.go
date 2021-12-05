package controllers

import "github.com/gofiber/fiber/v2"

type QuoteTagsController interface {
	Create(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
	FindById(c *fiber.Ctx) error
	FindAll(c *fiber.Ctx) error
}
