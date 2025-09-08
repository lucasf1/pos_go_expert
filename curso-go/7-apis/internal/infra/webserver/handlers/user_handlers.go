package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/lucasf1/goexpert/7-apis/internal/dto"
	"github.com/lucasf1/goexpert/7-apis/internal/entity"
	"github.com/lucasf1/goexpert/7-apis/internal/infra/database"
)

type UserHandler struct {
	UserDB database.UserInterface
}

func NewUserHandler(userDB database.UserInterface) *UserHandler {
	return &UserHandler{UserDB: userDB}
}

func (uh *UserHandler) Create(w http.ResponseWriter, r *http.Request) {

	var user dto.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u, err := entity.NewUser(user.Name, user.Email, user.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = uh.UserDB.Create(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
