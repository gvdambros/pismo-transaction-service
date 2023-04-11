package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"transaction-service/internal/domain/account"
	"transaction-service/internal/pkg/file"
	accountfixture "transaction-service/internal/testresources/fixtures/account"
	"transaction-service/internal/testresources/samples"
	"transaction-service/it"
	"transaction-service/it/config/db"
	httphelper "transaction-service/it/config/http"

	"github.com/go-resty/resty/v2"
	"github.com/labstack/echo/v4"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("an account submission", Ordered, Label("post"), func() {
	var (
		ctx     = context.Background()
		cli     *resty.Client
		req     *resty.Request
		baseURL string

		res *resty.Response
		err error
	)

	BeforeAll(func() {
		it.Init()
		baseURL = fmt.Sprintf("%s/%s", it.LocalServer().URL(), "accounts")
	})

	BeforeEach(func() {
		cli = resty.New()
		cli.Header = map[string][]string{
			echo.HeaderContentType: {echo.MIMEApplicationJSON},
		}
		req = cli.R().SetContext(ctx)
	})

	When("request is invalid", func() {
		DescribeTable("and payload is parseable but",
			func(field string, request *account.PostAccountRequest) {
				body, parseErr := json.Marshal(request)
				Expect(parseErr).ToNot(HaveOccurred())
				res, err = req.SetBody(body).Post(baseURL)

				Expect(err).ToNot(HaveOccurred())
				httpErr := httphelper.ParseError(res)
				Expect(httpErr.Status).To(Equal(http.StatusBadRequest))
				errMessage := fmt.Sprintf("%s.", field)
				Expect(httpErr.Message).To(ContainSubstring(errMessage))
			},

			Entry("document_number is empty", "document_number", accountfixture.RequestWith(
				func(r *account.PostAccountRequest) {
					r.DocumentNumber = ""
				}),
			),
			Entry("document_number is a text", "document_number", accountfixture.RequestWith(
				func(r *account.PostAccountRequest) {
					r.DocumentNumber = "teste"
				}),
			),
			Entry("document_number has 10 numbers", "document_number", accountfixture.RequestWith(
				func(r *account.PostAccountRequest) {
					r.DocumentNumber = "0000000000"
				}),
			),
		)
	})

	When("request is valid", func() {
		var (
			body    []byte
			samples = path.Join(samples.Path, "account")
		)

		AfterAll(func() { _ = db.Truncate(ctx, "accounts") })

		JustBeforeEach(func() { res, err = req.SetBody(body).Post(baseURL) })

		toAccount := func(r *resty.Response) account.Account {
			var account account.Account
			Expect(json.Unmarshal(r.Body(), &account)).ToNot(HaveOccurred())
			return account
		}

		Context("and account was successfully persisted", func() {
			BeforeEach(func() {
				body, _ = file.LoadBytes(path.Join(samples, "new_account.json"))
			})

			It("should respond with new account", func() {
				Expect(err).ToNot(HaveOccurred())
				dp := toAccount(res)
				Expect(dp.DocumentNumber).To(Equal("00427255007"))
			})
		})
	})
})
