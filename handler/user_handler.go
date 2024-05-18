package handler

import (
	"log"
	"net/http"

	"github.com/gookit/validate"
	"github.com/jmoiron/sqlx"
	"github.com/moroz/webauthn-academy-go/service"
	"github.com/moroz/webauthn-academy-go/templates/users"
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
	Errors validate.Errors
}

func (h *userHandler) New(w http.ResponseWriter, r *http.Request) {
	err := users.New(types.NewUserParams{}, nil).Render(r.Context(), w)
	if err != nil {
		log.Print(err)
	}
}

func (h *userHandler) Create(w http.ResponseWriter, r *http.Request) {
	var params types.NewUserParams
	err := decoder.Decode(&params, r.PostForm)
	if err != nil {
		handleError(w, http.StatusBadRequest)
		return
	}

	_, err, validationErrors := h.us.RegisterUser(params)

	if err != nil || validationErrors != nil {
		users.New(params, validationErrors).Render(r.Context(), w)
		err := users.New(types.NewUserParams{}, nil).Render(r.Context(), w)
		if err != nil {
			log.Print(err)
		}
		return
	}

	http.Redirect(w, r, "/sign-in", http.StatusMovedPermanently)
}
