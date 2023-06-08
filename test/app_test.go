package test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"workshoptdd/config"
	"workshoptdd/routes"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	app *gin.Engine
)

func TestMain(m *testing.M) {
	db = config.InitDatabase("host=localhost user=postgres password=postgres port=5432 dbname=workshop_tdd_test")
	app = routes.InitRoutes(db)

	m.Run()
}

func TestHealthCheck(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/healthcheck", nil)

	app.ServeHTTP(w, req)

	response := w.Result()
	assert.Equal(t, http.StatusOK, response.StatusCode)
	body, _ := io.ReadAll(response.Body)

	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, "OK", responseBody["message"])
}
