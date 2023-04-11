package handler

import (
	transaction "transaction-service/internal/domain/transaction"
	"transaction-service/internal/domain/transaction/repository"
)

type Handler struct {
	repository transaction.Repository
}

func NewTransaction() transaction.Handler {
	return &Handler{
		repository: repository.New(),
	}
}
