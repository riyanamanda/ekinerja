package validation

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

func Validate[T any](data T) map[string]string {
	err := validator.New().Struct(data)
	res := map[string]string{}
	if err != nil {
		for _, v := range err.(validator.ValidationErrors) {
			field := strings.ToLower(v.StructField())
			switch v.Tag() {
			case "required":
				res[field] = fmt.Sprintf("The field %s is required", field)
			default:
				res[field] = v.Error()
			}
		}
	}
	return res
}
