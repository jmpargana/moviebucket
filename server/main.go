package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

func run() error {
	connStr := "user=pggotest dbname=pggotest verify-ssl=never"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	if _, err := db.Exec(schema); err != nil {
		return err
	}
	srv := newServer(db)
	return http.ListenAndServe(":8000", srv.router)
}

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
}
