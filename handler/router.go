package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/moroz/webauthn-academy-go/db/queries"
)

func Router(db queries.DBTX) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(FetchSession)
	r.Use(FetchFlash)
	r.Use(FetchUserFromSession(db))

	r.Group(func(r chi.Router) {
		r.Use(RequireAuthenticatedUser)

		dashboard := DashboardHandler()
		r.Get("/", dashboard.Index)

		sessions := SessionHandler(db)
		r.Get("/sign-out", sessions.Delete)
	})

	r.Group(func(r chi.Router) {
		r.Use(RedirectIfAuthenticated)

		users := UserHandler(db)
		r.Get("/sign-up", users.New)
		r.Post("/sign-up", users.Create)

		sessions := SessionHandler(db)
		r.Get("/sign-in", sessions.New)
		r.Post("/sign-in", sessions.Create)
	})

	return r
}
