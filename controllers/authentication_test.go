package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestLoginControllerWithValidIdAndLogin(t *testing.T) {
	req, err := http.NewRequest("GET", "/login/1/cleison", nil)
	if err != nil {
		t.Fatal(err)
	}
	vars := map[string]string{
		"id":    "1",
		"login": "cleison",
	}
	req = mux.SetURLVars(req, vars)
	res := httptest.NewRecorder()

	LoginController(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
	assert.JSONEq(t, `{"created_at":"0001-01-01T00:00:00Z", "first_name":"Cleison", "id":1, "last_name":"Melo", "login":"cleison", "updated_at":"0001-01-01T00:00:00Z"}`, res.Body.String())
}

func TestLoginControllerWithInvalidId(t *testing.T) {
	req, err := http.NewRequest("GET", "/login/invalid/cleison", nil)
	if err != nil {
		t.Fatal(err)
	}
	vars := map[string]string{
		"id":    "invalid",
		"login": "cleison",
	}
	req = mux.SetURLVars(req, vars)
	res := httptest.NewRecorder()

	LoginController(res, req)

	assert.Equal(t, http.StatusBadRequest, res.Code)
	assert.JSONEq(t, `{"status":400,"message":"Invalid id query parameter"}`, res.Body.String())
}

func TestLoginControllerWithoutId(t *testing.T) {
	req, err := http.NewRequest("GET", "/login/cleison", nil)
	if err != nil {
		t.Fatal(err)
	}
	vars := map[string]string{
		"login": "cleison",
	}
	req = mux.SetURLVars(req, vars)
	res := httptest.NewRecorder()

	LoginController(res, req)

	assert.Equal(t, http.StatusBadRequest, res.Code)
	assert.JSONEq(t, `{"status":400,"message":"Id query parameter is required"}`, res.Body.String())
}

func TestLoginControllerWithoutLogin(t *testing.T) {
	req, err := http.NewRequest("GET", "/login/cleison", nil)
	if err != nil {
		t.Fatal(err)
	}
	vars := map[string]string{
		"id": "1",
	}
	req = mux.SetURLVars(req, vars)
	res := httptest.NewRecorder()

	LoginController(res, req)

	assert.Equal(t, http.StatusBadRequest, res.Code)
	assert.JSONEq(t, `{"status":400,"message":"Login query parameter is required"}`, res.Body.String())
}
