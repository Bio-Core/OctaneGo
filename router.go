package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

//StaticDir is the path to the static files hosted
const StaticDir = "/public/"

//AngDir is the path to the static files hosted
const AngDir = "/node_modules/"

//NewRouter creates a new router and maps all the routes defined in routes.go
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.NotFoundHandler = http.HandlerFunc(errorHandler)

	router.
		PathPrefix("/octane" + StaticDir).
		Handler(http.StripPrefix("/octane"+StaticDir, http.FileServer(http.Dir("."+StaticDir))))
	router.
		PathPrefix("/octane" + AngDir).
		Handler(http.StripPrefix("/octane"+AngDir, http.FileServer(http.Dir("."+AngDir))))
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}
