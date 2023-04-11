package account

import (
	"context"
)

type (
	Account struct {
		Id             int    `json:"id"`
		DocumentNumber string `json:"document_number"`
	}

	PostAccountRequest struct {
		DocumentNumber string `json:"document_number" validate:"required"`
	}

	GetAccountRequest struct {
		Id int `param:"id" validate:"required"`
	}

	// Handler a service for handling accounts
	Handler interface {
		PostAccount(ctx context.Context, request *PostAccountRequest) (*Account, error)
		GetAccount(ctx context.Context, request *GetAccountRequest) (*Account, error)
	}

	// Repository a service for persisting a account locally
	Repository interface {
		SaveAccount(context.Context, *PostAccountRequest) (*Account, error)
		GetAccount(ctx context.Context, request *GetAccountRequest) (*Account, error)
	}
)
