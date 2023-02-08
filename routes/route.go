package routes

import (
	"net/http"

	"github.com/cleisommais/oauth-service-v1/controllers"
	"github.com/cleisommais/oauth-service-v1/models"
)

// Routes are the main setup for our Router
type Routes []models.Route

var RoutesSetup = Routes{
	models.Route{"Login", "GET", "/login", controllers.LoginController},
}

// MakeHandler allows us to pass an environment struct to our handlers, without resorting to global
// variables. It accepts an environment (Env) struct and our own handler function. It returns
// a function of the type http.HandlerFunc so can be passed on to the HandlerFunc in main.go.
func MakeHandler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Terry Pratchett tribute
		w.Header().Set("X-Clacks-Overhead", "GNU Terry Pratchett")
		// return function with AppEnv
		fn(w, r)
	}
}