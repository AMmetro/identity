package utils

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

func FormatValidationErrors(err error) []string {
	var vErrs validator.ValidationErrors
	if errors.As(err, &vErrs) {
		out := make([]string, 0, len(vErrs))
		for _, ve := range vErrs {
			field := ve.Field() // short name of the struct field
			switch ve.Tag() {
			case "required":
				out = append(out, fmt.Sprintf("%s is required", field))
			default:
				out = append(out, fmt.Sprintf("%s failed on the '%s' tag", field, ve.Tag()))
			}
		}
		return out
	}
	return []string{err.Error()}
}
