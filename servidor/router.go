package main

import (
	"net/http"
)

type Router struct {
	rules map[string]http.HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		rules: make(map[string]http.HandlerFunc),
	}
}

func (r *Router) FindHandler(path string) (http.HandlerFunc, bool) {
	handler, exist := r.rules[path]
	return handler, exist
}

func (r *Router) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	// fmt.Fprintf(w, "Hello word") // Estatico
	handler, exist := r.FindHandler(request.URL.Path) // Respuesta dinamica

	if !exist {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	handler(w, request)
}
