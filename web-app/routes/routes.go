package routes

import (
	"net/http"

	"github.com/wes-santos/alura-golang-study/web-app/controllers"
)

func LoadRoutes() {
	http.HandleFunc("/", controllers.GetAllProducts)
	http.HandleFunc("/new", controllers.NewProduct)
	http.HandleFunc("/edit", controllers.EditProduct)
	http.HandleFunc("/insert", controllers.AddProduct)
	http.HandleFunc("/delete", controllers.DeleteProduct)
	http.HandleFunc("/update", controllers.UpdateProduct)
}
