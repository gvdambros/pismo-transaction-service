package error

// Wrapper a custom error wrapper
type Wrapper interface {
	Error() string
	Custom() *Error
}

// Error a http error body
// @Description an API error
type Error struct {
	Code     string                 `json:"error_code"`
	Status   int                    `json:"status_code"`
	Message  string                 `json:"message"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// Error error message (for matching default golang's Error interface)
func (d *Error) Error() string {
	return d.Message
}

// Custom returns a self reference, used for wrapping up unknown errors
func (d *Error) Custom() *Error {
	return d
}
