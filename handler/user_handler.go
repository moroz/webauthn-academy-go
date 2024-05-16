package handler

import (
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/moroz/webauthn-academy-go/handler/templates"
	"github.com/moroz/webauthn-academy-go/service"
	"github.com/moroz/webauthn-academy-go/types"
)

type userHandler struct {
	us service.UserService
}

func UserHandler(db *sqlx.DB) userHandler {
	return userHandler{service.NewUserService(db)}
}

type usersNewAssigns struct {
	RequestContext
	Params types.NewUserParams
}

func (h *userHandler) New(w http.ResponseWriter, r *http.Request) {
	err := templates.Users.New.Execute(w, usersNewAssigns{
		RequestContext: RequestContext{
			Title: "Register",
		},
		Params: types.NewUserParams{},
	})
	if err != nil {
		log.Print(err)
	}
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
