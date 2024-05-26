package handler

import (
	"context"
	"encoding/gob"
	"net/http"

	"github.com/gorilla/sessions"
	gorilla "github.com/gorilla/sessions"
	"github.com/jmoiron/sqlx"
	"github.com/moroz/webauthn-academy-go/config"
	"github.com/moroz/webauthn-academy-go/service"
	"github.com/moroz/webauthn-academy-go/types"
)

func init() {
	gob.Register(types.FlashMessage{})
}

var store = sessions.NewCookieStore(config.SessionSigner)

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
			if len(messages) > 0 {
				session.Save(r, w)
			}
		}
		ctx := context.WithValue(r.Context(), config.FlashContextKey, flashes)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func FetchUserFromSession(db *sqlx.DB) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			session, ok := r.Context().Value(config.SessionContextKey).(*gorilla.Session)
			if !ok {
				handleError(w, 500)
				return
			}

			var user *types.User

			if token, ok := session.Values[config.SessionUserTokenKey].([]byte); ok {
				srv := service.NewUserTokenService(db)
				userFromToken, err := srv.GetUserBySessionToken(token)
				if err != nil {
					user = userFromToken
				}
			}

			ctx := context.WithValue(r.Context(), config.UserContextKey, user)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func RequireAuthenticatedUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if user, ok := r.Context().Value(config.UserContextKey).(*types.User); ok && user != nil {
			next.ServeHTTP(w, r)
		}

		addFlash(r, w, types.FlashMessage{
			Severity: types.FlashMessageSeverity_Info,
			Content:  "You need to sign in to access this page.",
		})
		http.Redirect(w, r, "/sign-in", http.StatusMovedPermanently)
	})
}
