package routes

import (
	"net/http"

	"github.com/wes-santos/alura-golang-study/web-app/controllers"
)

func LoadRoutes() {
	http.HandleFunc("/", controllers.GetAllProducts)
}
