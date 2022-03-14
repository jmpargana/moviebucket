package main

func (s *Server) routes() {
	s.router.HandleFunc("POST", "/", s.handleGreet())
	s.router.HandleFunc("POST", "/movies", s.handleMoviePost())
}
