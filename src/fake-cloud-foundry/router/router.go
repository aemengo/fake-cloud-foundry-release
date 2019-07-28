package router

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type Router struct {
	handlers map[string]http.Handler
	logger   *log.Logger
}

func New() *Router {
	return &Router{
		handlers: map[string]http.Handler{},
		logger:   log.New(os.Stdout, "[ROUTER] ", log.LstdFlags),
	}
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.logger.Printf("Receiving [%s] request for %s%s...\n", req.Method, req.Host, req.URL.Path)

	handler, ok := r.handlers[req.Host]
	if !ok {
		http.Error(w, fmt.Sprintf(`404 Not Found: Requested route ('%s') does not exist.`, req.Host), http.StatusNotFound)
		return
	}

	handler.ServeHTTP(w, req)
}

func (r *Router) Add(host string, handler http.Handler) {
	r.handlers[host] = handler
}