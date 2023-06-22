package test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"workshoptdd/config"
	"workshoptdd/entity"
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

	// bersihin data
	db.Exec("TRUNCATE tasks RESTART IDENTITY")
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

func TestCreateTask(t *testing.T) {
	reqBody := strings.NewReader(`{
		"title": "example",
		"description": "example description"
	}`)

	var beforeCount int64
	db.Find(&entity.Task{}).Count(&beforeCount)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/tasks", reqBody)

	app.ServeHTTP(w, req)
	response := w.Result()

	assert.Equal(t, http.StatusCreated, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]string
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, "success create task", responseBody["message"])

	var afterCount int64
	db.Find(&entity.Task{}).Count(&afterCount)
	assert.Equal(t, afterCount, beforeCount+1)
}

func TestCreateTask_Fail(t *testing.T) {
	reqBody := strings.NewReader(`{
		"description": "example description"
	}`)

	var beforeCount int64
	db.Find(&entity.Task{}).Count(&beforeCount)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/tasks", reqBody)

	app.ServeHTTP(w, req)
	response := w.Result()

	assert.Equal(t, http.StatusBadRequest, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]string
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, "failed create task", responseBody["message"])

	var afterCount int64
	db.Find(&entity.Task{}).Count(&afterCount)
	assert.Equal(t, afterCount, beforeCount)
}

func TestGetTasks(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/tasks", nil)

	app.ServeHTTP(w, req)
	response := w.Result()

	assert.Equal(t, http.StatusOK, response.StatusCode)
	body, _ := io.ReadAll(response.Body)

	var responseBody []entity.Task
	json.Unmarshal(body, &responseBody)

	sampleTask := responseBody[0]
	if assert.NotEmpty(t, sampleTask) {
		assert.NotEmpty(t, sampleTask.Description)
		assert.NotEmpty(t, sampleTask.Title)
		assert.NotEmpty(t, sampleTask.ID)
	}
}

func TestDeleteTasks(t *testing.T) {
	dataToDelete := entity.Task{
		Title:       "delete",
		Description: "delete description",
	}
	db.Create(&dataToDelete)
	var beforeCount int64
	db.Find(&entity.Task{}).Count(&beforeCount)

	url := fmt.Sprintf("/tasks/%d", dataToDelete.ID)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, url, nil)

	app.ServeHTTP(w, req)
	response := w.Result()

	assert.Equal(t, http.StatusOK, response.StatusCode)
	body, _ := io.ReadAll(response.Body)

	var responseBody map[string]string
	json.Unmarshal(body, &responseBody)
	assert.Equal(t, "success delete task", responseBody["message"])

	var afterCount int64
	db.Find(&entity.Task{}).Count(&afterCount)
	assert.Equal(t, beforeCount-1, afterCount)
}

func TestDeleteTask_Fail(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, "/tasks/asd", nil)

	app.ServeHTTP(w, req)
	response := w.Result()

	assert.Equal(t, http.StatusBadRequest, response.StatusCode)
	body, _ := io.ReadAll(response.Body)

	var responseBody map[string]string
	json.Unmarshal(body, &responseBody)
	assert.Equal(t, "id invalid", responseBody["message"])
}
