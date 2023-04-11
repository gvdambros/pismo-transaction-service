package validator

import (
	"errors"
	"strings"
	httperror "transaction-service/internal/pkg/http/error"

	"github.com/go-playground/validator/v10"
)

// formats validation errors for usage in custom http errors
func (c *Validator) parseValidationErrors(err error) error {
	var errs validator.ValidationErrors
	if !errors.As(err, &errs) {
		return err
	}

	fields := make([]string, 0, len(errs)+1)

	for _, ve := range errs {
		fields = append(fields, ve.Field())
	}

	message := "invalid field(s): " + strings.Join(fields, ", ")

	return httperror.Build(httperror.BadRequest, message)
}
