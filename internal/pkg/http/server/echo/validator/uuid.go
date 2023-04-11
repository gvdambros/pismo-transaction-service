package validator

import (
	"reflect"

	"github.com/google/uuid"
)

// ValidateUUId custom uuid validation
func (c *Validator) validateUUID(field reflect.Value) interface{} {
	if uuidVal, ok := field.Interface().(uuid.UUID); ok {
		return uuidVal.String()
	}

	return nil
}
