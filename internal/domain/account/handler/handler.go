package handler

import (
	account "transaction-service/internal/domain/account"
	"transaction-service/internal/domain/account/repository"
)

type Handler struct {
	repository account.Repository
}

func NewAccount() account.Handler {
	return &Handler{
		repository: repository.New(),
	}
}
