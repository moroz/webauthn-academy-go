package handler

import (
	"errors"
	"log"
	"net/http"

	"github.com/moroz/webauthn-academy-go/db/queries"
	"github.com/moroz/webauthn-academy-go/service"
	"github.com/moroz/webauthn-academy-go/templates/users"
	"github.com/moroz/webauthn-academy-go/types"
)

type userHandler struct {
	us service.UserService
}

func UserHandler(db queries.DBTX) userHandler {
	return userHandler{service.NewUserService(db)}
}

var SignupError = errors.New("One or more errors prevented this user from being created. Please review the errors in the form below.")

func (h *userHandler) New(w http.ResponseWriter, r *http.Request) {
	err := users.New(types.NewUserParams{}, nil, nil).Render(r.Context(), w)
	if err != nil {
		log.Printf("Rendering error: %s", err)
	}
}

func (h *userHandler) Create(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // Ignore the error as it will be handled later

	var params types.NewUserParams
	if err := decoder.Decode(&params, r.PostForm); err != nil {
		handleError(w, http.StatusBadRequest)
		return
	}

	_, err, validationErrors := h.us.RegisterUser(r.Context(), params)

	if err != nil || validationErrors != nil {
		w.Header().Add("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusUnprocessableEntity)
		err := users.New(params, SignupError, validationErrors).Render(r.Context(), w)
		if err != nil {
			log.Print(err)
		}
		return
	}

	http.Redirect(w, r, "/sign-in", http.StatusSeeOther)
}
