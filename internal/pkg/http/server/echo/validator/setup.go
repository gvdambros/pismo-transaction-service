package validator

import (
	"github.com/google/uuid"
)

func (c *Validator) setup() {
	c.validator.RegisterCustomTypeFunc(c.validateUUID, uuid.UUID{})
	c.validator.RegisterTagNameFunc(c.registerJSONTag)
}
