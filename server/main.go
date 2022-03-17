package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	oauthConfig = &oauth2.Config{
		ClientID:     "410752640842-5nto8t0e3ov13acarb7itcqtn7gtvgbg.apps.googleusercontent.com",
		ClientSecret: "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",
		RedirectURL:  "http://localhost:8080/callback-gl",
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
)

func run() error {
	connStr := "user=user pass=pass dbname=pggotest sslmode=disable"
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
