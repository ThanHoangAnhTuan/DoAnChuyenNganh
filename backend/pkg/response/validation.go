package response

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	Field   string      `json:"field"`
	Message string      `json:"message"`
	Value   interface{} `json:"value"`
}

func FormatValidationErrorsToStruct(err error) []ValidationError {
	var errors []ValidationError

	for _, err := range err.(validator.ValidationErrors) {
		var message string
		switch err.Tag() {
		case "required":
			message = "Trường này là bắt buộc"
		case "email":
			message = "Email không hợp lệ"
		case "min":
			message = fmt.Sprintf("Tối thiểu %s ký tự", err.Param())
		case "max":
			message = fmt.Sprintf("Tối đa %s ký tự", err.Param())
		default:
			message = "Giá trị không hợp lệ"
		}

		errors = append(errors, ValidationError{
			Field:   err.Field(),
			Message: message,
			Value:   err.Value(),
		})
	}

	return errors
}
