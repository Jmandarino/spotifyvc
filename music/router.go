package music

import (
	"net/http"
	"github.com/gorilla/mux"
)

var controller = &Controller{DBconnection: DBconnection{}}

// Route defines a route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes defines the list of routes of our API
type Routes []Route
var routes = Routes {

	Route {
		"Index",
		"GET",
		"/",
		controller.Index,
	},
	Route {
		"TrackPlaylist",
		"POST",
		"/track",
		controller.TrackPlaylist,
	},
}

func NewRouter() *mux.Router{
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}
	return router
}