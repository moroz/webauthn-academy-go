package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/moroz/webauthn-academy-go/config"
	"github.com/moroz/webauthn-academy-go/types"
)

var store = sessions.NewCookieStore(config.SessionSigner)

func ParseForm(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			w.Header().Add("content-type", "text/plain; charset=utf-8")
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "Bad Request")
			return
		}

		next.ServeHTTP(w, r)
	})
}

func FetchSession(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, config.SessionKey)
		ctx := context.WithValue(r.Context(), config.SessionContextKey, session)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func FetchFlash(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		flashes := make([]types.FlashMessage, 0)
		if session, ok := r.Context().Value(config.SessionContextKey).(*sessions.Session); ok {
			messages := session.Flashes()
			for _, msg := range messages {
				if msg, ok := msg.(types.FlashMessage); ok {
					flashes = append(flashes, msg)
				}
			}
		}
		ctx := context.WithValue(r.Context(), config.FlashContextKey, flashes)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
