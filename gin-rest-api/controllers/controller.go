package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wes-santos/gin-rest-api/database"
	"github.com/wes-santos/gin-rest-api/models"
)

func GetAllStudents(c *gin.Context) {
	var students []models.Student

	database.DB.Find(&students)

	c.JSON(200, students)
}

func Greetings(c *gin.Context) {
	name, founded := c.Params.Get("name")
	if !founded {
		c.JSON(200, gin.H{
			"API says:": "Hey, what's up? Sorry, I can't get your name from URL.",
		})
		return
	}

	c.JSON(200, gin.H{
		"API says": "Hey, " + name + ", how you doing?",
	})
}

func AddNewStudent(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := models.ValidateStudentData(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	database.DB.Create(&student)
	c.JSON(http.StatusCreated, student)
}

func GetStudentById(c *gin.Context) {
	id := c.Params.ByName("id")

	var student models.Student
	database.DB.First(&student, id)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Student not found in database.",
		})
		return
	}

	c.JSON(http.StatusOK, student)
}

func DeleteStudent(c *gin.Context) {
	id := c.Params.ByName("id")

	database.DB.Delete(&models.Student{}, id)
	c.JSON(http.StatusNoContent, gin.H{})
}

func UpdateStudent(c *gin.Context) {
	id := c.Params.ByName("id")

	var student models.Student

	database.DB.First(&student, id)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Student not found in database. Create the student before update.",
		})
		return
	}

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := models.ValidateStudentData(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	database.DB.Model(&student).UpdateColumns(student)
	c.JSON(http.StatusAccepted, student)
}

func GetStudentByCPF(c *gin.Context) {
	var student models.Student
	cpf := c.Param("cpf")

	database.DB.Where(&models.Student{CPF: cpf}).First(&student)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Student not found in database.",
		})
		return
	}

	c.JSON(http.StatusOK, student)
}
