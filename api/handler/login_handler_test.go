package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dasagho/playground/api/routers"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	router := routers.SetupRouter()

	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	expectedResponse := `{"message":"Everything right!!"}`

	assert.Equal(t, expectedResponse, w.Body.String())

}
