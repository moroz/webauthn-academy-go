package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jmoiron/sqlx"
)

func Router(db *sqlx.DB) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(FetchSession)
	r.Use(FetchFlash)

	users := UserHandler(db)
	r.Get("/", users.New)
	r.Post("/users/register", users.Create)

	sessions := SessionHandler(db)
	r.Get("/sign-in", sessions.New)
	r.Post("/sign-in", sessions.Create)

	return r
}
