package api

import (
	"net/http"
	account "transaction-service/internal/domain/account"
	_ "transaction-service/internal/pkg/http/error" // swagger only

	"github.com/labstack/echo/v4"
)

// Post godoc
// @Summary submits an account
// @Tags 	accounts
// @Accept	json
// @Produce	json
// @Success 200	{object}	account.Account
// @Failure 400	{object}	error.Error
// @Failure 404	{object}	error.Error
// @Failure 500	{object}	error.Error
// @Param	request body	account.PostAccountRequest true	"a account request"
// @Router	/accounts		[post]
func (a *API) Post(c echo.Context) error {
	ctx := c.Request().Context()

	var request account.PostAccountRequest
	if err := c.Bind(&request); err != nil {
		return err
	}

	if err := c.Validate(request); err != nil {
		return err
	}

	acc, err := a.handler.PostAccount(ctx, &request)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, acc)
}

// Find godoc
// @Summary submits an account
// @Tags 	accounts
// @Accept	json
// @Produce	json
// @Success 200	{object}	account.Account
// @Failure 400	{object}	error.Error
// @Failure 404	{object}	error.Error
// @Failure 500	{object}	error.Error
// @Param 	id	path	string	true	"an account id"
// @Router	/accounts/{id}		[get]
func (a *API) Get(c echo.Context) error {
	ctx := c.Request().Context()

	var request account.GetAccountRequest
	if err := c.Bind(&request); err != nil {
		return err
	}

	if err := c.Validate(request); err != nil {
		return err
	}

	acc, err := a.handler.GetAccount(ctx, &request)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, acc)
}
