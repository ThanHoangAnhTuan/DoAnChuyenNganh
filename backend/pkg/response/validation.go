package response

import (
	"fmt"
	"reflect"

	"github.com/go-playground/validator/v10"
)

type FieldError struct {
	Field   string      `json:"field"`
	Message string      `json:"message"`
	Value   interface{} `json:"value"`
}

func translateMessage(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "Trường này là bắt buộc"
	case "email":
		return "Email không hợp lệ"
	case "min":
		return fmt.Sprintf("Giá trị tối thiểu là %s ký tự", fe.Param())
	case "max":
		return fmt.Sprintf("Giá trị tối đa là %s ký tự", fe.Param())
	case "alphanum":
		return "Chỉ được chứa chữ và số"
	case "len":
		return fmt.Sprintf("Độ dài phải đúng %s ký tự", fe.Param())
	case "numeric":
		return "Chỉ được chứa số"
	case "eq":
		return fmt.Sprintf("Phải có giá trị đúng bằng '%s'", fe.Param())
	case "ne":
		return fmt.Sprintf("Không được có giá trị là '%s'", fe.Param())
	case "url":
		return "URL không hợp lệ"
	case "uuid":
		return "UUID không hợp lệ"
	case "strongpassword":
		return "Mật khẩu phải có ít nhất 8 ký tự, chứa ít nhất 1 chữ hoa, 1 chữ thường, 1 chữ số và 1 ký tự đặc biệt"
	default:
		return "Giá trị không hợp lệ"
	}
}

func FormatValidationErrorsToStruct(err error, s interface{}) []FieldError {
	var errs []FieldError
	t := reflect.TypeOf(s)

	// Nếu là con trỏ thì lấy phần tử bên dưới
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	for _, err := range err.(validator.ValidationErrors) {
		fieldName := err.Field()

		// Lấy tên trường theo tag json (nếu có)
		if f, ok := t.FieldByName(err.StructField()); ok {
			jsonTag := f.Tag.Get("json")
			if jsonTag != "" {
				fieldName = jsonTag
			}
		}

		errs = append(errs, FieldError{
			Field:   fieldName,
			Message: translateMessage(err),
			Value:   err.Value(),
		})
	}
	return errs
}
