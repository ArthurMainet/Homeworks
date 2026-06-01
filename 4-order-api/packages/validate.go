package packages

import "github.com/go-playground/validator/v10"

// Нужно дописать разбор конкретных ошибок на случай проблем с валидацией
func Validate[T any](isValid T) error {
	validate := validator.New()
	err := validate.Struct(isValid)
	return err
}
