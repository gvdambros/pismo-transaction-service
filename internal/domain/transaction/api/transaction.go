package api

import (
	"net/http"
	transaction "transaction-service/internal/domain/transaction"
	_ "transaction-service/internal/pkg/http/error" // swagger only

	"github.com/labstack/echo/v4"
)

// New godoc
// @Summary submits a transaction
// @Tags 	transactions
// @Accept	json
// @Produce	json
// @Success 200	{object}	transaction.Transaction
// @Failure 400	{object}	error.Error
// @Failure 404	{object}	error.Error
// @Failure 500	{object}	error.Error
// @Param	request body	transaction.PostTransactionRequest true	"a transaction request"
// @Router	/transactions		[post]
func (a *API) New(c echo.Context) error {
	ctx := c.Request().Context()

	var request transaction.PostTransactionRequest
	if err := c.Bind(&request); err != nil {
		return err
	}

	if err := c.Validate(request); err != nil {
		return err
	}

	transaction, err := a.handler.Transaction(ctx, &request)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, transaction)
}
