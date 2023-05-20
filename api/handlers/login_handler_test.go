package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func init() {
	gin.SetMode(gin.TestMode)
	router = gin.Default()
}

func TestLoginHandler(t *testing.T) {
	router.POST("/login", PostLogin)

	requestBody := bytes.NewBuffer([]byte(`{"user": 55645093, "pass": "pepe05422"}`))
	badRequestBody := bytes.NewBuffer([]byte(`{"user": 5564509, "pass": "pepe0542"}`))

	t.Run("Sent invalid json", func(t *testing.T) {
		req, err := http.NewRequest("POST", "/login", nil)
		if err != nil {
			t.Fatalf("Couldn't create request: %s\n", err.Error())
		}

		req.Header.Set("Content-Type", "application/json")

		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		if rec.Code != http.StatusBadRequest {
			t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusBadRequest, rec.Code)
		}
	})

	t.Run("Sent bad credentials", func(t *testing.T) {
		req, err := http.NewRequest("POST", "/login", badRequestBody)
		if err != nil {
			t.Fatalf("Couldn't create request: %s\n", err.Error())
		}

		req.Header.Set("Content-Type", "application/json")

		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		if rec.Code != http.StatusNotFound {
			t.Fatalf("Expected to get status %d but instead got %d with error %s\n", http.StatusNotFound, rec.Code, rec.Body)
		}
	})

	t.Run("Correct login", func(t *testing.T) {
		req, err := http.NewRequest("POST", "/login", requestBody)
		if err != nil {
			t.Fatalf("Couldn't create request: %s\n", err.Error())
		}

		req.Header.Set("Content-Type", "application/json")

		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		if rec.Code != http.StatusOK {
			t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, rec.Code)
		}
	})
}
