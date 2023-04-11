package repository

import (
	"time"
	transaction "transaction-service/internal/domain/transaction"
	dbmodel "transaction-service/internal/domain/transaction/db/model"
)

func (r *Repository) convertDomainToDB(request *transaction.PostTransactionRequest) *dbmodel.Transaction {
	return &dbmodel.Transaction{
		AccountId:       request.AccountId,
		OperationTypeId: request.OperationTypeId,
		Amount:          request.Amount,
		EventDate:       time.Now(),
	}
}

func (r *Repository) convert(res *dbmodel.Transaction) *transaction.Transaction {
	return &transaction.Transaction{
		Id:              res.Id,
		AccountId:       res.AccountId,
		OperationTypeId: res.OperationTypeId,
		Amount:          res.Amount,
		EventDate:       res.EventDate,
	}
}
