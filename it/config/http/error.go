package http

import (
	"encoding/json"
	httperror "transaction-service/internal/pkg/http/error"

	"github.com/go-resty/resty/v2"
	"github.com/onsi/gomega"
)

// ParseError parses a response error
func ParseError(r *resty.Response) httperror.Error {
	var httpErr httperror.Error
	gomega.Expect(json.Unmarshal(r.Body(), &httpErr)).ToNot(gomega.HaveOccurred())

	return httpErr
}
