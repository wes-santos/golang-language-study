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
)

func SetupTestRoutes() *gin.Engine {
	routes := gin.Default()
	return routes
}

func TestCheckGreetingEndpointStatusCodeWithParam(t *testing.T) {
	r := SetupTestRoutes()

	r.GET("/:name", controllers.Greetings)
	req, _ := http.NewRequest("GET", "/wes", nil)
	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code, "Retrivied status code is different from 200")
}

func TestCheckGreetingEndpointBodyWithParam(t *testing.T) {
	r := SetupTestRoutes()

	r.GET("/:name", controllers.Greetings)

	reqName := "wes"

	req, _ := http.NewRequest("GET", "/"+reqName, nil)
	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)

	mockResp := fmt.Sprintf(`{"API says":"Hey, %s, how you doing?"}`, reqName)

	respBody, err := io.ReadAll(resp.Body)
	assert.Nil(t, err)

	assert.Equal(t, mockResp, string(respBody))
}
