package api

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestAccountAPI(t *testing.T) {
	if testing.Short() {
		t.Skip("skipped")
	}

	RegisterFailHandler(Fail)
	RunSpecs(t, "Transaction API suite", Label("transaction", "api", "integration"))
}
