package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"simple-go-project/db"
	"simple-go-project/models"
	"simple-go-project/utils"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/tasks", utils.AddTask)
	r.GET("/tasks", utils.GetTasks)
	r.GET("/tasks/:id", utils.GetTask)
	r.PUT("/tasks/:id", utils.UpdateTask)
	r.DELETE("/tasks/:id", utils.DeleteTask)
	return r
}

func TestAddTask(t *testing.T) {
	router := setupRouter()
	test_deadline := time.Now().Add(time.Hour)
	task := models.Task{HashKey: uuid.MustParse("00000000-0000-0000-0000-000000000000"), Name: "Test Task", Description: "Test Desc", Closed: true, Deadline: test_deadline}
	jsonStr, _ := json.Marshal(task)
	req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `"name":"Test Task"`)
	req, _ = http.NewRequest("DELETE", "/tasks/00000000-0000-0000-0000-000000000000", bytes.NewBuffer(jsonStr))
	router.ServeHTTP(w, req)
}

func TestAddTaskValidationName(t *testing.T) {
	router := setupRouter()
	test_deadline := time.Now().Add(time.Hour)
	task := models.Task{HashKey: uuid.MustParse("00000000-0000-0000-0000-000000000000"), Name: "T", Description: "Test Desc", Closed: true, Deadline: test_deadline}
	jsonStr, _ := json.Marshal(task)
	req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	// TODO: if validation text fixed assert must be changed as well
	assert.Contains(t, w.Body.String(), `Field: Name, Error: min`)
	req, _ = http.NewRequest("DELETE", "/tasks/00000000-0000-0000-0000-000000000000", bytes.NewBuffer(jsonStr))
	router.ServeHTTP(w, req)
}

func TestAddTaskValidationDeadline(t *testing.T) {
	router := setupRouter()
	test_deadline := time.Date(1000, time.April, 1, 0, 0, 0, 0, time.Local)
	task := models.Task{HashKey: uuid.MustParse("00000000-0000-0000-0000-000000000000"), Name: "Test Task", Description: "Test Desc", Closed: true, Deadline: test_deadline}
	jsonStr, _ := json.Marshal(task)
	req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	// TODO: if validation text fixed assert must be changed as well
	assert.Contains(t, w.Body.String(), `"Field: Deadline, Error: deadlineValidator`)
	req, _ = http.NewRequest("DELETE", "/tasks/00000000-0000-0000-0000-000000000000", bytes.NewBuffer(jsonStr))
	router.ServeHTTP(w, req)
}

func TestGetTasks(t *testing.T) {
	var dbase, _ = db.OpenDBWriteConnection()
	test_deadline := time.Now().Add(time.Hour)
	router := setupRouter()
	_, err := dbase.Exec("INSERT INTO todo_list.tasks (hash_key, name, description, closed, deadline) VALUES ('00000000-0000-0000-0000-000000000000', 'Test Task', 'Test Desc', true, $1)", test_deadline)
	if err != nil {
		panic(err)
	}

	req, _ := http.NewRequest("GET", "/tasks", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `"name":"Test Task"`)
	_, err = dbase.Exec("DELETE FROM todo_list.tasks WHERE hash_key = '00000000-0000-0000-0000-000000000000';")
	if err != nil {
		panic(err)
	}
	db.CloseDBConnection(dbase)
}
