package handler

import (
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/schema"
	"github.com/gorilla/sessions"
	"github.com/moroz/webauthn-academy-go/config"
	"github.com/moroz/webauthn-academy-go/types"
)

var decoder = schema.NewDecoder()

func handleError(w http.ResponseWriter, status int) {
	w.Header().Add("content-type", "text/plain; charset=utf-8")
	w.WriteHeader(status)
	msg := http.StatusText(status)
	w.Write([]byte(msg))
}

func addFlash(r *http.Request, w http.ResponseWriter, msg types.FlashMessage) error {
	session, ok := r.Context().Value(config.SessionContextKey).(*sessions.Session)
	if !ok {
		return errors.New("Failed to fetch session")
	}

	session.AddFlash(msg)
	err := session.Save(r, w)
	if err != nil {
		log.Print(err)
		handleError(w, 500)
	}
	return err
}
