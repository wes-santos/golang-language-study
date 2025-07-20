package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/wes-santos/gin-rest-api/controllers"
	"github.com/wes-santos/gin-rest-api/database"
	"github.com/wes-santos/gin-rest-api/models"
)

var ID int

func SetupTestRoutes() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	routes := gin.Default()
	return routes
}

func CreateStudentMock() {
	student := models.Student{Name: "Student Mock Name", CPF: "12345678901", RG: "123456789"}

	database.DB.Create(&student)
	ID = int(student.ID)
}

func DeleteStudentMock() {
	var student models.Student

	database.DB.Delete(&student, ID)

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
	CreateStudentMock()
	defer DeleteStudentMock()

	r := SetupTestRoutes()
	r.GET("/students", controllers.GetAllStudents)

	req, err := http.NewRequest("GET", "/students", nil)
	assert.Nil(t, err)

	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
}

func TestGetStudentByCPFHandler(t *testing.T) {
	database.ConnectWithDatabase()
	CreateStudentMock()
	defer DeleteStudentMock()

	r := SetupTestRoutes()
	r.GET("/students/cpf/:cpf", controllers.GetStudentByCPF)
	req, err := http.NewRequest("GET", "/students/cpf/12345678901", nil)
	assert.Nil(t, err)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
}

func TestGetStudentByIDHandler(t *testing.T) {
	database.ConnectWithDatabase()
	CreateStudentMock()
	defer DeleteStudentMock()

	r := SetupTestRoutes()
	r.GET("/students/:id", controllers.GetStudentById)
	endpointPath := "/students/" + strconv.Itoa(ID)

	req, err := http.NewRequest("GET", endpointPath, nil)
	assert.Nil(t, err)

	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	var mockStudent models.Student
	err = json.Unmarshal(res.Body.Bytes(), &mockStudent)
	assert.Nil(t, err)

	assert.Equal(t, "Student Mock Name", mockStudent.Name)
	assert.Equal(t, "12345678901", mockStudent.CPF)
	assert.Equal(t, "123456789", mockStudent.RG)
	assert.Equal(t, http.StatusOK, res.Code)
}
