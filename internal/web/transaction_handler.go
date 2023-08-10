package web

import (
	"encoding/json"
	"fmt"
	"net/http"

	createtransaction "github.com.br/devfullcycle/fc-ms-wallet/internal/usecase/create_transaction"
)

type WebTransactionHandler struct {
	CreateTransactionUseCase createtransaction.CreateTransactionUseCase
}

func NewWebTransactionHandler(
	createTransactionUseCase createtransaction.CreateTransactionUseCase,
) *WebTransactionHandler {
	return &WebTransactionHandler{
		CreateTransactionUseCase: createTransactionUseCase,
	}
}

func (web *WebTransactionHandler) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var dto createtransaction.CreateTransactionInputDto //setando a váriavel para esse tipo
	err := json.NewDecoder(r.Body).Decode(&dto)         //pegando os dados que vieram do request, e parseando eles para o var dto
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	output, err := web.CreateTransactionUseCase.Execute(ctx, dto)
	fmt.Println("Deu erro", err)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json") //Setando o tipo de resposta no cabeçalho
	err = json.NewEncoder(w).Encode(output)            //Codifuca o output em JSON para enviar como retorno
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// w.Write([]byte())
	w.WriteHeader(http.StatusCreated)
}
