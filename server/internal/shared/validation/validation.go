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
			case "min":
				res[field] = fmt.Sprintf("The field %s must be at least %s characters", field, v.Param())
			case "max":
				res[field] = fmt.Sprintf("The field %s must be at most %s characters", field, v.Param())
			case "len":
				res[field] = fmt.Sprintf("The field %s must be exactly %s characters", field, v.Param())
			case "eq":
				res[field] = fmt.Sprintf("The field %s must be equal to %s", field, v.Param())
			case "ne":
				res[field] = fmt.Sprintf("The field %s must not be equal to %s", field, v.Param())
			case "gt":
				res[field] = fmt.Sprintf("The field %s must be greater than %s", field, v.Param())
			case "gte":
				res[field] = fmt.Sprintf("The field %s must be greater than or equal to %s", field, v.Param())
			case "lt":
				res[field] = fmt.Sprintf("The field %s must be less than %s", field, v.Param())
			case "lte":
				res[field] = fmt.Sprintf("The field %s must be less than or equal to %s", field, v.Param())
			case "email":
				res[field] = fmt.Sprintf("The field %s must be a valid email address", field)
			case "url":
				res[field] = fmt.Sprintf("The field %s must be a valid URL", field)
			case "uuid":
				res[field] = fmt.Sprintf("The field %s must be a valid UUID", field)
			case "oneof":
				res[field] = fmt.Sprintf("The field %s must be one of [%s]", field, v.Param())
			default:
				res[field] = v.Error()
			}
		}
	}
	return res
}
