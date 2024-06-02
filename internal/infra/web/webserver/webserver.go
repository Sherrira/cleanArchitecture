package webserver

import (
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type WebServer struct {
	Router        chi.Router
	Handlers      map[string]http.HandlerFunc
	WebServerPort string
}

func NewWebServer(serverPort string) *WebServer {
	return &WebServer{
		Router:        chi.NewRouter(),
		Handlers:      make(map[string]http.HandlerFunc),
		WebServerPort: serverPort,
	}
}

// AddHandler adds a handler to the handlers map.
// httpMethod: GET, POST, PUT, DELETE.
// path: the path to the handler.
// handler: the handler function.
func (s *WebServer) AddHandler(httpMethod string, path string, handler http.HandlerFunc) {
	mapKey := httpMethod + ":" + path
	s.Handlers[mapKey] = handler
}

// loop through the handlers and add them to the router
// register middeleware logger
// start the server
func (s *WebServer) Start() {
	s.Router.Use(middleware.Logger)
	for mapKey, handler := range s.Handlers {
		parts := strings.Split(mapKey, ":")

		switch parts[0] {
		case "GET":
			s.Router.Get(parts[1], handler)
		case "POST":
			s.Router.Post(parts[1], handler)
		case "PUT":
			s.Router.Put(parts[1], handler)
		case "DELETE":
			s.Router.Delete(parts[1], handler)
		}
	}
	http.ListenAndServe(s.WebServerPort, s.Router)
}
