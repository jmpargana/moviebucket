package main

func (s *Server) routes() {
	s.router.HandleFunc("POST", "/", s.handleGreet())
	s.router.HandleFunc("POST", "/movies", s.handleMoviePost())
	s.router.HandleFunc("POST", "/login", s.handleLogin())
	s.router.HandleFunc("POST", "/callback-gl", s.handleGoogleCallback())
}
