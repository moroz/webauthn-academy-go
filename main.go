package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/moroz/webauthn-academy-go/config"
	"github.com/moroz/webauthn-academy-go/handler"
)

func main() {
	db := sqlx.MustConnect("postgres", config.DatabaseURL)

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(handler.FetchSession)
	r.Use(handler.FetchFlash)

	users := handler.UserHandler(db)
	r.Get("/", users.New)
	r.Post("/users/register", users.Create)

	sessions := handler.SessionHandler(db)
	r.Get("/sign-in", sessions.New)

	log.Println("Listening on port 3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
