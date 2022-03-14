package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/matryer/way"
)

type Server struct {
	db     *sql.DB
	router *way.Router
}

func newServer() *Server {
	router := way.NewRouter()
	srv := &Server{
		db:     nil,
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

func (s *Server) decode(w http.ResponseWriter, r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}

func (s *Server) handleGreet() http.HandlerFunc {
	type request struct {
		Name string `json:"name"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := s.decode(w, r, req); err != nil {
			s.respond(w, r, "failed reading request body", 400)
			return
		}
		s.respond(w, r, fmt.Sprintf("hello %s", req.Name), 200)
	}
}

func (s *Server) routes() {
	s.router.HandleFunc("POST", "/", s.handleGreet())
}

func run() error {
	// dbsetup
	srv := newServer()
	return http.ListenAndServe(":8000", srv.router)
}

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
}
