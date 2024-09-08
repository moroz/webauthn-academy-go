package main

import (
	"context"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
	"github.com/moroz/webauthn-academy-go/config"
	"github.com/moroz/webauthn-academy-go/handler"
)

func main() {
	db, err := pgxpool.New(context.Background(), config.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}

	r := handler.Router(db)
	log.Printf("Listening on %s", config.ListenOn)
	log.Fatal(http.ListenAndServe(config.ListenOn, r))
}
