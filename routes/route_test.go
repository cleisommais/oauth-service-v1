package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cleisommais/oauth-service-v1/controllers"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestLoginRoute(t *testing.T) {
	req, err := http.NewRequest("GET", "/login", nil)
	if err != nil {
		t.Fatal(err)
	}
	vars := map[string]string{
		"id":    "1",
		"login": "cleison",
	}
	req = mux.SetURLVars(req, vars)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.LoginController)

	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.JSONEq(t, `{"created_at":"0001-01-01T00:00:00Z", "first_name":"Cleison", "id":1, "last_name":"Melo", "login":"cleison", "updated_at":"0001-01-01T00:00:00Z"}`, rr.Body.String())
}
