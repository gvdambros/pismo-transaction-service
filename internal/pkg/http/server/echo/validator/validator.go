package validator

import (
	"github.com/go-playground/validator/v10"
)

// Validator go-playground validator wrapper
type Validator struct{ validator *validator.Validate }

// New retrieves a custom validator reference
func New() *Validator {
	v := &Validator{validator: validator.New()}
	v.setup()

	return v
}

// Validate performs validation and emits a custom error on failure
func (c *Validator) Validate(i interface{}) error {
	if err := c.validator.Struct(i); err != nil {
		return c.parseValidationErrors(err)
	}

	return nil
}
