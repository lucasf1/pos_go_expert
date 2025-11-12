package webserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type WebServer struct {
	Router        chi.Router
	Routes        map[string]map[string]http.HandlerFunc
	WebServerPort string
}

func NewWebServer(serverPort string) *WebServer {
	return &WebServer{
		Router:        chi.NewRouter(),
		Routes:        make(map[string]map[string]http.HandlerFunc),
		WebServerPort: serverPort,
	}
}

func (s *WebServer) AddRoute(method string, path string, handler http.HandlerFunc) {
	if _, ok := s.Routes[path]; !ok {
		s.Routes[path] = make(map[string]http.HandlerFunc)
	}
	s.Routes[path][method] = handler
}

// loop through the handlers and add them to the router
// register middeleware logger
// start the server
func (s *WebServer) Start() {
	
	s.Router.Use(middleware.Logger)
	for path, methods := range s.Routes {
		for method, handler := range methods {
			s.Router.Method(method, path, handler)
		}
	}
	http.ListenAndServe(":"+s.WebServerPort, s.Router)
}