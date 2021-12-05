package controllers

import "github.com/gofiber/fiber/v2"

type AuthorsController interface {
	Create(c *fiber.Ctx) error
	FindById(c *fiber.Ctx) error
	FindAll(c *fiber.Ctx) error
	FindBySlug(c *fiber.Ctx) error
	FindAuthorAndQuotes(c *fiber.Ctx) error
}