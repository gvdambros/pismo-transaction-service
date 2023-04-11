package handler

import (
	"context"
	transaction "transaction-service/internal/domain/transaction"

	"github.com/rs/zerolog/log"
)

func (h *Handler) Transaction(ctx context.Context, request *transaction.PostTransactionRequest) (*transaction.Transaction, error) {
	log.Info().Msgf("Creating transaction: %v", request)

	transaction, err := h.repository.SaveTransaction(ctx, request)

	if err != nil {
		return nil, err
	}

	log.Info().Msgf("Transaction created: %v", transaction)

	return transaction, nil
}
