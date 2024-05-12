package handler

import (
	"net/http"

	"github.com/gorilla/schema"
	"github.com/jmoiron/sqlx"
	"github.com/moroz/webauthn-academy-go/service"
	"github.com/moroz/webauthn-academy-go/types"
)

var decoder = schema.NewDecoder()

type userHandler struct {
	us service.UserService
}

func UserHandler(db *sqlx.DB) userHandler {
	return userHandler{service.NewUserService(db)}
}

func (h *userHandler) Create(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		handleError(w, http.StatusBadRequest)
		return
	}

	var params types.NewUserParams
	err = decoder.Decode(&params, r.PostForm)
	if err != nil {
		handleError(w, http.StatusBadRequest)
		return
	}
}
