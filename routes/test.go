// routes/routes.go
package routes

import (
	"github.com/cleisommais/oauth-service-v1/controllers"
	"net/http"
)

// TestHandler is a handler function that serves the /test route
func TestHandler(w http.ResponseWriter, r *http.Request) {
	controllers.TestController(w, r)
}
