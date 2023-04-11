package validator

import (
	"reflect"
	"strings"
)

const separatorPos = 2

// nolint:gocritic
// this replaces struct field name for json tag when reporting errors
func (c *Validator) registerJSONTag(field reflect.StructField) string {
	name := strings.SplitN(field.Tag.Get("json"), ",", separatorPos)[0]
	if name == "-" {
		return ""
	}

	return name
}
