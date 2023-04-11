package account

import (
	"transaction-service/internal/domain/account"
)

// Post account request fixture
func Request() *account.PostAccountRequest {
	return &account.PostAccountRequest{
		DocumentNumber: "00427255007",
	}
}

// RequestWith
func RequestWith(with func(r *account.PostAccountRequest)) *account.PostAccountRequest {
	r := Request()
	with(r)

	return r
}
