package notes

import (
	"github.com/gorilla/mux"
	"net/http"
)

var Repository = NewRepository("localhost", "pqgotest", "test", "password")
var Controller = NewNoteController(*Repository)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

var routes = []Route{
	{
		"GET /notes",
		"GET",
		"/notes",
		Controller.GetNotes,
	},
	{
		"GET /notes/{id}",
		"GET",
		"/notes/{id}",
		Controller.GetNoteById,
	},
}

func NewRouter() (router *mux.Router) {
	router = mux.NewRouter()

	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc

		router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(handler)
	}

	return
}
