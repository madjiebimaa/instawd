package helpers

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/madjiebimaa/instawd/models"
)

func ToErrorFields(err error) []models.FieldErrors {
	var errors []models.FieldErrors
	for _, err := range err.(validator.ValidationErrors) {
		mess := fmt.Sprintf("invalid validation %s for %s tag", err.Field(), err.Tag())
		er := models.FieldErrors{
			Field:   err.Field(),
			Message: strings.ToLower(mess),
		}

		errors = append(errors, er)
	}

	return errors
}
