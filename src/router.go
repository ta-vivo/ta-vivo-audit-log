package main

import (
	"encoding/json"
	"net/http"
)

// A Router is a map of strings to http.HandlerFuncs.
// @property rules - A map of string keys and http.HandlerFunc values.
type Router struct {
	rules map[string]map[string]http.HandlerFunc
}

// It creates a new Router object and returns a pointer to it
func NewRouter() *Router {
	return &Router{
		rules: make(map[string]map[string]http.HandlerFunc),
	}
}

// A method that takes a path and returns a handler and a boolean.
func (r *Router) FindHandler(path string, method string) (http.HandlerFunc, bool, bool) {
	_, exist := r.rules[path]
	handler, methodExist := r.rules[path][method]
	return handler, methodExist, exist
}

// A method that implements the ServeHTTP method of the http.Handler interface.
func (r *Router) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	handler, methodExist, exist := r.FindHandler(request.URL.Path, request.Method)
	w.Header().Set("Content-Type", "application/json")

	if !exist {
		var responseMap = make(map[string]string)
		responseMap["message"] = "Not found"
		response, _ := json.Marshal(responseMap)

		w.WriteHeader(http.StatusNotFound)
		w.Write(response)
		return
	}

	if !methodExist {
		var responseMap = make(map[string]string)
		responseMap["message"] = "Method not allowed"
		response, _ := json.Marshal(responseMap)

		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write(response)
		return
	}

	handler(w, request)
}
