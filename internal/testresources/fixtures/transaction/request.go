package transaction

import (
	"transaction-service/internal/domain/transaction"
)

// Post transaction request fixture
func Request() *transaction.PostTransactionRequest {
	return &transaction.PostTransactionRequest{
		AccountId:       1,
		OperationTypeId: 1,
		Amount:          1,
	}
}

// RequestWith
func RequestWith(with func(r *transaction.PostTransactionRequest)) *transaction.PostTransactionRequest {
	r := Request()
	with(r)

	return r
}
