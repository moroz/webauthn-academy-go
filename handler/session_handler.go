package handler

import (
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/moroz/webauthn-academy-go/config"
	"github.com/moroz/webauthn-academy-go/service"
	"github.com/moroz/webauthn-academy-go/templates/sessions"
	"github.com/moroz/webauthn-academy-go/types"

	gorilla "github.com/gorilla/sessions"
)

type sessionHandler struct {
	us service.UserService
	ts service.UserTokenService
}

func SessionHandler(db *sqlx.DB) sessionHandler {
	return sessionHandler{service.NewUserService(db), service.NewUserTokenService(db)}
}

func (h *sessionHandler) New(w http.ResponseWriter, r *http.Request) {
	err := sessions.New().Render(r.Context(), w)
	if err != nil {
		log.Print(err)
	}
}

func (h *sessionHandler) Create(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		handleError(w, http.StatusBadRequest)
		return
	}

}

func (h *sessionHandler) signUserIn(w http.ResponseWriter, r *http.Request, user *types.User) {
	session, ok := r.Context().Value(config.SessionContextKey).(*gorilla.Session)
	if !ok {
		handleError(w, 500)
		return
	}

	token, err := h.ts.GenerateUserSessionToken(user)
	if err != nil {
		log.Print(err)
		handleError(w, 500)
		return
	}

	session.Values[config.SessionUserTokenKey] = token
	if err := session.Save(r, w); err != nil {
		log.Print(err)
		handleError(w, 500)
	}
}
