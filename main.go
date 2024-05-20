package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/moroz/webauthn-academy-go/handler"
)

func main() {
	db := sqlx.MustConnect("postgres", os.Getenv("DATABASE_URL"))

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	users := handler.UserHandler(db)
	r.Get("/", users.New)

	log.Println("Listening on port 3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
