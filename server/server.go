package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/matryer/way"
)

type Server struct {
	db     *sql.DB
	router *way.Router
}

func newServer(db *sql.DB) *Server {
	router := way.NewRouter()
	srv := &Server{
		db:     db,
		router: router,
	}
	srv.routes()
	return srv
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *Server) respond(w http.ResponseWriter, r *http.Request, data interface{}, status int) {
	w.WriteHeader(status)
	if data != nil {
		err := json.NewEncoder(w).Encode(data)
		fmt.Println(err)
	}
}

func (s *Server) fail(w http.ResponseWriter, r *http.Request, err *Err) {
	w.WriteHeader(err.status)
	json.NewEncoder(w).Encode(err)
}

func (s *Server) decode(w http.ResponseWriter, r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}
