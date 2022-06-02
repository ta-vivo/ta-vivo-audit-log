package main

import "net/http"

// A Server is a struct that contains a port and a router.
// @property {string} port - The port that the server will listen on.
// @property router - This is the router that will be used to handle the requests.
type Server struct {
	port   string
	router *Router
}

// NewServer returns a pointer to a new Server instance.
func NewServer(port string) *Server {
	return &Server{
		port:   port,
		router: NewRouter(),
	}
}

func (s *Server) Handle(method string, path string, handler http.HandlerFunc) {
	_, exist := s.router.rules[path]

	if !exist {
		s.router.rules[path] = make(map[string]http.HandlerFunc)
	}

	s.router.rules[path][method] = handler
}

// A function that takes a handler function and a list of middlewares. It then loops through the
// middlewares and calls each one with the handler function.
func (s *Server) AddMiddleware(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}

	return f
}

// Listening for requests on the port specified in the server struct.
func (s *Server) Listen() error {
	// Entry point
	http.Handle("/", s.router)
	err := http.ListenAndServe(s.port, nil)
	if err != nil {
		return err
	}

	return nil
}
