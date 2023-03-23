package web

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com.br/devfullcycle/fc-ms-wallet/internal/usecase/create_account"
)

type WebAccountHandler struct {
    CreateAccountUseCase create_account.CreateAccountUseCase
}

func NewWebAccountHandler(createAccountUseCase create_account.CreateAccountUseCase) *WebAccountHandler {
    return &WebAccountHandler{
        CreateAccountUseCase: createAccountUseCase,
    }
}

func (h *WebAccountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
    var dto create_account.CreateAccountInputDTO
    err := json.NewDecoder(r.Body).Decode(&dto)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        fmt.Println(err)
        return
    }

    output, err := h.CreateAccountUseCase.Execute(dto)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Println(err)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    err = json.NewEncoder(w).Encode(output)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusCreated)
}