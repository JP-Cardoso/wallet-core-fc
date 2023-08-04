package web

import (
	"encoding/json"
	"net/http"

	createaccount "github.com.br/devfullcycle/fc-ms-wallet/internal/usecase/create_account"
)

type WebAccountHandler struct {
	createAccountUseCase createaccount.CreateAccountUseCase
}

func NewWebAccountHandler(createAccountUseCase createaccount.CreateAccountUseCase) *WebAccountHandler {
	return &WebAccountHandler{
		createAccountUseCase: createAccountUseCase,
	}
}

func (web *WebAccountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var dto createaccount.CreateAccountInputDto //setando a váriavel para esse tipo
	err := json.NewDecoder(r.Body).Decode(&dto) //pegando os dados que vieram do request, e parseando eles para o var dto
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	output, err := web.createAccountUseCase.Execute(dto)
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
