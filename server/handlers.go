package main

import (
	"fmt"
	"net/http"
)

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
		tx, _ := s.db.Begin()
		tx.Exec("UPDATE counts SET count += 1 WHERE id = 1")
		tx.Commit()
		s.respond(w, r, fmt.Sprintf("hello %s", req.Name), 200)
	}
}

func (s *Server) handleMoviePost() http.HandlerFunc {
	type request struct {
		Name string `json:"name"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := s.decode(w, r, req); err != nil || len(req.Name) < 1 {
			s.fail(w, r, errors["movie_decode"])
			return
		}

		out, err := s.db.Exec("INSERT INTO movies (name) VALUES ('Harry Potter')")
		fmt.Println(out, err)

		s.respond(w, r, nil, 200)
	}
}
