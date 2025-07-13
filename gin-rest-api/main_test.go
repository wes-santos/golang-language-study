package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/wes-santos/gin-rest-api/controllers"
	"github.com/wes-santos/gin-rest-api/database"
)

func SetupTestRoutes() *gin.Engine {
	routes := gin.Default()
	return routes
}

func TestCheckGreetingEndpointStatusCodeWithParam(t *testing.T) {
	r := SetupTestRoutes()

	r.GET("/:name", controllers.Greetings)
	req, _ := http.NewRequest("GET", "/wes", nil)
	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code, "Retrivied status code is different from 200")
}

func TestCheckGreetingEndpointBodyWithParam(t *testing.T) {
	r := SetupTestRoutes()

	r.GET("/:name", controllers.Greetings)

	reqName := "wes"

	req, _ := http.NewRequest("GET", "/"+reqName, nil)
	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)

	mockRes := fmt.Sprintf(`{"API says":"Hey, %s, how you doing?"}`, reqName)

	respBody, err := io.ReadAll(res.Body)
	assert.Nil(t, err)

	assert.Equal(t, mockRes, string(respBody))
}

func TestListAllStudentsHandler(t *testing.T) {
	database.ConnectWithDatabase()

	r := SetupTestRoutes()
	r.GET("/students", controllers.GetAllStudents)

	req, err := http.NewRequest("GET", "/students", nil)
	assert.Nil(t, err)

	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
}
