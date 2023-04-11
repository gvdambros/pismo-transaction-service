package error

import "fmt"

// Build builds a custom error
func Build(err Error, args ...interface{}) *Error {
	err.Message = buildMessage(err.Message, args...)
	return &err
}

// BuildWithMeta builds a custom error with a metadata map
func BuildWithMeta(err Error, meta map[string]interface{}, args ...interface{}) *Error {
	err.Message = buildMessage(err.Message, args...)
	err.Metadata = meta

	return &err
}

// conditionally uses arguments for building an error message
func buildMessage(msg string, args ...interface{}) string {
	if len(args) < 1 {
		return msg
	}

	return fmt.Sprintf(msg, args...)
}
