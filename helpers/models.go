package helpers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/madjiebimaa/instawd/models"
)

func WebResponseSuccess(c *fiber.Ctx, now time.Time, status int, data interface{}) error {
	return c.Status(status).Type(fiber.MIMEApplicationJSON).JSON(models.WebResponse{
		Took:    uint(time.Until(now).Milliseconds()),
		Success: true,
		Status:  status,
		Data:    data,
		Errors:  nil,
	})
}

func WebResponseFail(c *fiber.Ctx, now time.Time, status int, errField string, errMessage error) error {
	return c.Status(status).Type(fiber.MIMEApplicationJSON).JSON(models.WebResponse{
		Took:    uint(time.Until(now).Milliseconds()),
		Success: false,
		Status:  status,
		Data:    nil,
		Errors: []models.FieldErrors{
			{Field: errField, Message: errMessage.Error()},
		},
	})
}

func WebResponseFails(c *fiber.Ctx, now time.Time, status int, errors []models.FieldErrors) error {
	return c.Status(status).Type(fiber.MIMEApplicationJSON).JSON(models.WebResponse{
		Took:    uint(time.Until(now).Milliseconds()),
		Success: false,
		Status:  status,
		Data:    nil,
		Errors:  errors,
	})
}
