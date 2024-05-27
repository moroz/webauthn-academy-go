package handler

import (
	"log"
	"net/http"

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

func (h *userHandler) New(w http.ResponseWriter, r *http.Request) {
	err := users.New(types.NewUserParams{}, nil).Render(r.Context(), w)
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

	_, err, validationErrors := h.us.RegisterUser(params)

	if err != nil || validationErrors != nil {
		addFlash(r, w, types.FlashMessage{
			Severity: types.FlashMessageSeverity_Error,
			Content:  "One or more errors prevented the record from being saved. Please review the errors in the form below.",
		})
		w.Header().Add("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusUnprocessableEntity)
		err := users.New(params, validationErrors).Render(r.Context(), w)
		if err != nil {
			log.Print(err)
		}
		return
	}

	http.Redirect(w, r, "/sign-in", http.StatusSeeOther)
}
