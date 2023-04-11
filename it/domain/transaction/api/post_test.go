package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"transaction-service/internal/domain/transaction"
	"transaction-service/internal/pkg/file"
	transactionfixture "transaction-service/internal/testresources/fixtures/transaction"
	"transaction-service/internal/testresources/samples"
	"transaction-service/it"
	"transaction-service/it/config/db"
	httphelper "transaction-service/it/config/http"
	"transaction-service/it/config/seeds"

	"github.com/go-resty/resty/v2"
	"github.com/labstack/echo/v4"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("an transaction submission", Ordered, Label("post"), func() {
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
		baseURL = fmt.Sprintf("%s/%s", it.LocalServer().URL(), "transactions")

		if err := seeds.Load(ctx, "account/accounts.sql"); err != nil {
			Fail(err.Error())
			defer GinkgoRecover()
		}
	})

	BeforeEach(func() {
		cli = resty.New()
		cli.Header = map[string][]string{
			echo.HeaderContentType: {echo.MIMEApplicationJSON},
		}
		req = cli.R().SetContext(ctx)
	})

	AfterAll(func() { _ = db.Truncate(ctx, "accounts", "transactions") })

	When("request is invalid", func() {
		DescribeTable("and payload is parseable but",
			func(field string, request *transaction.PostTransactionRequest) {
				body, parseErr := json.Marshal(request)
				Expect(parseErr).ToNot(HaveOccurred())
				res, err = req.SetBody(body).Post(baseURL)

				Expect(err).ToNot(HaveOccurred())
				httpErr := httphelper.ParseError(res)

				fmt.Fprintf(GinkgoWriter, "Some log text: %v\n", httpErr)

				Expect(httpErr.Status).To(Equal(http.StatusBadRequest))
				errMessage := fmt.Sprintf("invalid field(s): %s.", field)
				Expect(httpErr.Message).To(ContainSubstring(errMessage))
			},

			Entry("account_id is invalid", "account_id", transactionfixture.RequestWith(
				func(r *transaction.PostTransactionRequest) {
					r.AccountId = -1
				}),
			),

			Entry("operation_type_id is invalid", "operation_type_id", transactionfixture.RequestWith(
				func(r *transaction.PostTransactionRequest) {
					r.OperationTypeId = -1
				}),
			),

			Entry("amount is zero", "amount", transactionfixture.RequestWith(
				func(r *transaction.PostTransactionRequest) {
					r.Amount = 0
				}),
			),
		)
	})

	When("request is valid", func() {
		var (
			body    []byte
			samples = path.Join(samples.Path, "transaction")
		)

		JustBeforeEach(func() { res, err = req.SetBody(body).Post(baseURL) })

		toTransaction := func(r *resty.Response) transaction.Transaction {
			var transaction transaction.Transaction
			Expect(json.Unmarshal(r.Body(), &transaction)).ToNot(HaveOccurred())
			return transaction
		}

		Context("and account was successfully persisted", func() {
			BeforeEach(func() {
				body, _ = file.LoadBytes(path.Join(samples, "new_transaction.json"))
			})

			It("should respond with new transaction", func() {
				Expect(err).ToNot(HaveOccurred())
				t := toTransaction(res)
				Expect(t.AccountId).To(Equal(1))
				Expect(t.OperationTypeId).To(Equal(1))
				Expect(t.Amount).To(Equal(1.59))
			})
		})
	})
})
