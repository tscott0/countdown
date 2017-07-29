package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type routes []route

func newRouter() *mux.Router {

	var allRoutes = routes{
		route{
			"LettersGame",
			"GET",
			"/letters/{letters}",
			lettersShow,
		},
	}

	router := mux.NewRouter().StrictSlash(true)
	for _, r := range allRoutes {
		router.
			Methods(r.Method).
			Path(r.Pattern).
			Name(r.Name).
			Handler(r.HandlerFunc)
	}

	return router
}
