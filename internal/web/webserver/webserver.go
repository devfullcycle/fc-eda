package webserver

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

type WebServer struct {
    Router chi.Router
    Handlers map[string]http.HandlerFunc
    WebServerPort string
}

func NewWebServer(webServerPort string) *WebServer {
    return &WebServer{
        Router: chi.NewRouter(),
        Handlers: make(map[string]http.HandlerFunc),
        WebServerPort: webServerPort,
    }
}

func (s *WebServer) AddHandler(path string, handler http.HandlerFunc) {
    s.Handlers[path] = handler
}

func (s *WebServer) Start() {
    s.Router.Use(middleware.Logger)
    for path, handler := range s.Handlers {
        s.Router.Post(path, handler)
    }
    http.ListenAndServe(s.WebServerPort, s.Router)
}