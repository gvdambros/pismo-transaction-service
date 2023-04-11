package handler

import (
	"transaction-service/internal/domain/account"

	httperror "transaction-service/internal/pkg/http/error"

	"github.com/klassmann/cpfcnpj"
)

func (h *Handler) validate(request *account.PostAccountRequest) error {
	cpf := cpfcnpj.NewCPF(request.DocumentNumber)

	if cpf.IsValid() {
		return nil
	}

	cnpj := cpfcnpj.NewCNPJ(request.DocumentNumber)

	if cnpj.IsValid() {
		return nil
	}

	return httperror.Build(httperror.BadRequest, "document_number")
}
