package handler_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/dasagho/playground/api/routers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var router *gin.Engine = routers.SetupRouter()

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	exitCode := m.Run()

	os.Exit(exitCode)
}

func TestLogin(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	t.Run("Raíz código HTTP 200", func(t *testing.T) {
		status := assert.Equal(t, http.StatusOK, w.Code)
		if !status {
			t.Errorf("Código de respuesta erroneo %d", w.Code)
		}
	})

	t.Run("Raíz respuesta todo ok", func(t *testing.T) {
		expectedResponse := `{"message":"Everything right!!"}`
		res := assert.Equal(t, expectedResponse, w.Body.String())
		if !res {
			t.Errorf("Respuesta erronea")
		}
	})
}
