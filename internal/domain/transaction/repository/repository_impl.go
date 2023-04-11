package repository

import (
	"context"
	transaction "transaction-service/internal/domain/transaction"
)

func (r *Repository) SaveTransaction(ctx context.Context, request *transaction.PostTransactionRequest) (*transaction.Transaction, error) {
	transaction := r.convertDomainToDB(request)

	result := r.database.Connection.Create(transaction)

	if result.Error != nil {
		return nil, result.Error
	}

	return r.convert(transaction), nil
}
