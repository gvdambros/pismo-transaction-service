package error

import "net/http"

var (
	// BadRequest unmapped validation errors (i.e.: invalid uuid param)
	BadRequest = Error{
		Code:    "REQ-001",
		Status:  http.StatusBadRequest,
		Message: "invalid request: %+v.",
	}

	// NotFound unmapped 404's (i.e.: a wrong path)
	NotFound = Error{
		Code:    "REQ-002",
		Status:  http.StatusNotFound,
		Message: "resource not found: %+v.",
	}

	// ServerError unmapped unexpected error
	ServerError = Error{
		Code:    "ISE-000",
		Status:  http.StatusInternalServerError,
		Message: "unexpected server error: %+v.",
	}
)
