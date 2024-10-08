package handler

import (
	"errors"
	"log"
	"net/http"

	"github.com/moroz/webauthn-academy-go/config"
	"github.com/moroz/webauthn-academy-go/db/queries"
	"github.com/moroz/webauthn-academy-go/service"
	"github.com/moroz/webauthn-academy-go/templates/sessions"

	gorilla "github.com/gorilla/sessions"
)

type sessionHandler struct {
	us service.UserService
	ts service.UserTokenService
}

func SessionHandler(db queries.DBTX) sessionHandler {
	return sessionHandler{service.NewUserService(db), service.NewUserTokenService(db)}
}

func (h *sessionHandler) New(w http.ResponseWriter, r *http.Request) {
	err := sessions.New("", nil).Render(r.Context(), w)
	if err != nil {
		log.Print(err)
	}
}

func (h *sessionHandler) Create(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		handleError(w, http.StatusBadRequest)
		return
	}

	email := r.PostForm.Get("email")

	user, err := h.us.AuthenticateUserByEmailPassword(r.Context(), email, r.PostForm.Get("password"))

	if err == nil {
		h.signUserIn(w, r, user)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	err = sessions.New(email, errors.New("Invalid email/password combination.")).Render(r.Context(), w)
	if err != nil {
		log.Print(err)
	}
}

func (h *sessionHandler) Delete(w http.ResponseWriter, r *http.Request) {
	h.signUserOut(w, r)

	http.Redirect(w, r, "/sign-in", http.StatusSeeOther)
}

func (h *sessionHandler) signUserIn(w http.ResponseWriter, r *http.Request, user *queries.User) {
	session, ok := r.Context().Value(config.SessionContextKey).(*gorilla.Session)
	if !ok {
		handleError(w, 500)
		return
	}

	token, err := h.ts.GenerateUserSessionToken(r.Context(), user)
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

func (h *sessionHandler) signUserOut(w http.ResponseWriter, r *http.Request) error {
	session, ok := r.Context().Value(config.SessionContextKey).(*gorilla.Session)
	if !ok {
		handleError(w, 500)
		return errors.New("signUserOut: session not found in request context")
	}

	delete(session.Values, config.SessionUserTokenKey)
	if err := session.Save(r, w); err != nil {
		log.Print(err)
		handleError(w, 500)
	}
	return nil
}
