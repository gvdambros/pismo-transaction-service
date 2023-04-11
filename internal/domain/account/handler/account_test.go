package handler

import (
	account "transaction-service/internal/domain/account"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Verify account", func() {

	var (
		h = Handler{}
	)

	When("validating document_number", func() {
		DescribeTable("and documentNumber is invalid",
			func(request account.PostAccountRequest) {
				result := h.validate(&request)
				Expect(result).NotTo(BeNil())
			},

			Entry("empty document", account.PostAccountRequest{
				DocumentNumber: "",
			}),
			Entry("string", account.PostAccountRequest{
				DocumentNumber: "teste",
			}),
			Entry("8 numbers", account.PostAccountRequest{
				DocumentNumber: "12345678",
			}),
			Entry("10 numbers", account.PostAccountRequest{
				DocumentNumber: "1234567890",
			}),
			Entry("invalid 9 numbers cpf", account.PostAccountRequest{
				DocumentNumber: "51402689064",
			}),
			Entry("13 numbers", account.PostAccountRequest{
				DocumentNumber: "1234567890123",
			}),
			Entry("15 numbers", account.PostAccountRequest{
				DocumentNumber: "123456789012345",
			}),
			Entry("invalid 14 numbers cnpj", account.PostAccountRequest{
				DocumentNumber: "67210622000116",
			}),
		)

		DescribeTable("and documentNumber is valid",
			func(request account.PostAccountRequest) {
				result := h.validate(&request)
				Expect(result).To(BeNil())
			},

			Entry("9 numbers cpf", account.PostAccountRequest{
				DocumentNumber: "51402689063",
			}),
			Entry("formated 9 numbers cpf", account.PostAccountRequest{
				DocumentNumber: "514.026.890-63",
			}),
			Entry("14 numbers cnpj", account.PostAccountRequest{
				DocumentNumber: "67210622000115",
			}),
			Entry("formated 14 numbers cnpj", account.PostAccountRequest{
				DocumentNumber: "67.210.622/0001-15",
			}),
		)
	},
	)
})
