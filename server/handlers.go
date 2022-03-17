package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
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

func (s *Server) handleLogin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		URL, err := url.Parse(oauthConfig.Endpoint.AuthURL)
		if err != nil {
			fmt.Println(err)
		}
		parameters := url.Values{}
		parameters.Add("client_id", oauthConfig.ClientID)
		parameters.Add("scope", strings.Join(oauthConfig.Scopes, " "))
		parameters.Add("redirect_uri", oauthConfig.RedirectURL)
		parameters.Add("response_type", "code")
		parameters.Add("state", "")
		URL.RawQuery = parameters.Encode()
		uri := URL.String()
		http.Redirect(w, r, uri, http.StatusTemporaryRedirect)
	}
}

func (s *Server) handleGoogleCallback() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		state := r.FormValue("state")
		code := r.FormValue("code")
		fmt.Println(state, code)
	}
}
