package main

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	db := sqlx.MustConnect("postgres", "postgres://postgres:postgres@localhost/academy_dev?sslmode=disable")
	row := db.QueryRowx("select 2 + 2")
	var value int
	err := row.Scan(&value)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(value)
}
