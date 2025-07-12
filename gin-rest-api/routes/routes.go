package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wes-santos/gin-rest-api/controllers"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/students", controllers.GetAllStudents)
	r.GET("/students/:id", controllers.GetStudentById)
	r.GET("students/cpf/:cpf", controllers.GetStudentByCPF)
	r.GET("/:name", controllers.Greetings)
	r.POST("/students", controllers.AddNewStudent)
	r.PUT("/students/:id", controllers.UpdateStudent)
	r.DELETE("/students/:id", controllers.DeleteStudent)
	r.Run()
}
