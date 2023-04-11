package handler

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestAccountHandler(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Account handler suite", Label("account", "handler"))
}
