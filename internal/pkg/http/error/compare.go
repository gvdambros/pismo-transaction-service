package error

import "errors"

// Is checks for error equality exactly like errors.Is
func Is(err error, expected Error) bool {
	var httpErr *Error
	if errors.As(err, &httpErr) {
		sameCode := httpErr.Code == expected.Code
		sameStatus := httpErr.Status == expected.Status

		return sameCode && sameStatus
	}

	return false
}
