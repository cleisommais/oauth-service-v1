// routes/routes.go
package routes

import (
	"net/http"
	"github.com/cleisommais/oauth-service-v1/controllers"	
)

// TestHandler is a handler function that serves the /test route
func TestHandler(w http.ResponseWriter, r *http.Request) {
	controllers.TestController(w, r)
}