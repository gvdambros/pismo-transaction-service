package handler

import (
	"context"
	account "transaction-service/internal/domain/account"

	"github.com/rs/zerolog/log"
)

func (h *Handler) PostAccount(ctx context.Context, request *account.PostAccountRequest) (*account.Account, error) {
	log.Info().Msgf("Creating account: %v", request)

	if err := h.validate(request); err != nil {
		return nil, err
	}

	a, err := h.repository.SaveAccount(ctx, request)

	if err != nil {
		return nil, err
	}

	log.Info().Msgf("Account created: %v", a)

	return a, nil
}

func (h *Handler) GetAccount(ctx context.Context, request *account.GetAccountRequest) (*account.Account, error) {
	a, err := h.repository.GetAccount(ctx, request)

	if err != nil {
		return nil, err
	}

	return a, nil
}
