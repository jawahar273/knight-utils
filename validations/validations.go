package validations

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/jawahar273/knight-utils/core"
)

// https://blog.depa.do/post/gin-validation-errors-handling

type JSONFormatter struct{}

// NewJSONFormatter will create a new JSON formatter and register a custom tag
// name func to gin's validator
func NewJSONFormatter() *JSONFormatter {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
	}

	return &JSONFormatter{}
}

func (JSONFormatter) generateValidatoinMessage(validatorErrors *validator.ValidationErrors) []core.ErrorFieldAndMessage {
	errs := []core.ErrorFieldAndMessage{}

	for _, f := range *validatorErrors {
		err := f.ActualTag()
		if f.Param() != " " {
			err = fmt.Sprintf("%s=%s", err, f.Param())
		}
		errs = append(errs, core.ErrorFieldAndMessage{Field: f.Field(), Message: err})
	}
	return errs
}

func BearerAuthHeader(authHeader string) string {
	if authHeader == "" {
		return ""
	}

	parts := strings.Split(authHeader, "Bearer")
	if len(parts) != 2 {
		return ""
	}

	token := strings.TrimSpace(parts[1])
	if len(token) < 1 {
		return ""
	}

	return token
}
