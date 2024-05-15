package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/moroz/webauthn-academy-go/handler"
)

func MustGetenv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		msg := fmt.Sprintf("FATAL: Environment variable %s is not set", key)
		log.Fatal(msg)
	}
	return value
}

func main() {
	db := sqlx.MustConnect("postgres", MustGetenv("DATABASE_URL"))

	r := chi.NewRouter()

	users := handler.UserHandler(db)
	r.Get("/", users.New)

	sessions := handler.SessionHandler(db)
	r.Get("/sign-in", sessions.New)

	log.Println("Listening on port 3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
