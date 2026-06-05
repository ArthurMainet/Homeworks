package api

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	Field   string `json:"field"`
	Tag     string `json:"tag"`
	Message string `json:"message"`
}

type ValidationErrors struct {
	Errors []ValidationError
}

func (err ValidationErrors) Error() string {
	return "validation failed"
}

// Паттерн обработки ошибок валидации для удобства фронта.
func Validate[T any](isValid T) error {
	validate := validator.New()
	err := validate.Struct(isValid)
	if err != nil {
		var validErrs []ValidationError
		for _, err := range err.(validator.ValidationErrors) {
			validErrs = append(validErrs, ValidationError{
				Field:   err.Field(),
				Tag:     err.Tag(),
				Message: getErrMessage(err),
			})
		}
		return ValidationErrors{
			Errors: validErrs,
		}
	}
	return nil
}

func getErrMessage(err validator.FieldError) string {
	switch err.Tag() {
	case "requred":
		return "Это поле должно быть обязательным"
	case "email":
		return "Некорректный email"
	case "min":
		return fmt.Sprintf("Минимальная длина %s", err.Param())
	case "max":
		return fmt.Sprintf("Максимальная длина %s", err.Param())
	case "gt":
		return fmt.Sprintf("Значение должно быть больше %s", err.Param())
	default:
		return "Некорректное значение"
	}
}
