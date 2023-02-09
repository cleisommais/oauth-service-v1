package routes

import (
	"net/http"

	"github.com/cleisommais/oauth-service-v1/controllers"
	"github.com/cleisommais/oauth-service-v1/models"
)

// Routes are the main setup for our Router
type Routes []models.Route

var RoutesSetup = Routes{
	models.Route{Name: "Login", Method: "GET", Pattern: "/login", HandlerFunc: controllers.LoginController},
}

// function of the type http.HandlerFunc so can be passed on to the HandlerFunc in main.go.
func MakeHandler(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return handlerFunc
}
