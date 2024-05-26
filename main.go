package main

import (
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/moroz/webauthn-academy-go/config"
	"github.com/moroz/webauthn-academy-go/handler"
)

func main() {
	db := sqlx.MustConnect("postgres", config.DatabaseURL)
	r := handler.Router(db)
	log.Printf("Listening on %s", config.ListenOn)
	log.Fatal(http.ListenAndServe(config.ListenOn, r))
}
