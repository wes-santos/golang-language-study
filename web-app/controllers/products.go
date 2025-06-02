package controllers

import (
	"html/template"
	"log"
	"net/http"

	"github.com/wes-santos/alura-golang-study/web-app/models"
)

var templates = template.Must(template.ParseGlob("web-app/templates/*.html"))

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := models.GetAllProducts()

	if err != nil {
		log.Println("Error scanning line: ", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	templates.ExecuteTemplate(w, "Index", products)
}
