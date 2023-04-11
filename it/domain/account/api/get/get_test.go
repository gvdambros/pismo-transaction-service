package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"transaction-service/internal/domain/account"
	"transaction-service/it"
	"transaction-service/it/config/db"
	httphelper "transaction-service/it/config/http"
	"transaction-service/it/config/seeds"

	"github.com/go-resty/resty/v2"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("an account retrieval API", Ordered, Label("get"), func() {
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
		baseURL = fmt.Sprintf("%s/%s", it.LocalServer().URL(), "accounts/%s")

		if err := seeds.Load(ctx, "account/accounts.sql"); err != nil {
			Fail(err.Error())
			defer GinkgoRecover()
		}
	})

	AfterAll(func() { _ = db.Truncate(ctx, "accounts") })

	BeforeEach(func() {
		cli = resty.New()
		req = cli.R().SetContext(ctx)
	})

	When("couldn't find a matching account", func() {
		DescribeTable("while searching by",
			func(missingId string) {
				res, err = req.Get(fmt.Sprintf(baseURL, missingId))
				Expect(err).ToNot(HaveOccurred())
				httpErr := httphelper.ParseError(res)
				// should be NotFound
				Expect(httpErr.Status).To(Equal(http.StatusInternalServerError))
			},

			Entry("id", "2"),
		)
	})

	When("found an account", func() {
		toAccount := func(r *resty.Response) account.Account {
			var account account.Account
			Expect(json.Unmarshal(r.Body(), &account)).ToNot(HaveOccurred())
			return account
		}
		DescribeTable("while searching by",
			func(id string) {
				res, err = req.Get(fmt.Sprintf(baseURL, id))
				Expect(err).ToNot(HaveOccurred())
				account := toAccount(res)
				Expect(account.Id).To(Equal(1))
				Expect(account.DocumentNumber).To(Equal("00427255007"))
			},

			Entry("id", "1"),
		)
	})
})
