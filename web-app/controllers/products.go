package controllers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/wes-santos/alura-golang-study/web-app/models"
)

var templates = template.Must(template.ParseGlob("/Users/weslley/Studies/alura/golang/web-app/templates/*.html"))

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := models.GetAllProducts()

	if err != nil {
		log.Println("Error scanning line: ", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	templates.ExecuteTemplate(w, "Index", products)
}

func NewProduct(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "New", nil)
}

func EditProduct(w http.ResponseWriter, r *http.Request) {
	productId, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(
			w,
			fmt.Sprintf("Error parsing id: %v", err),
			http.StatusBadRequest,
		)
		return
	}

	product, err := models.GetProductById(productId)
	if err != nil {
		http.Error(
			w,
			fmt.Sprintf("Error fetching product from database: %v", err),
			http.StatusInternalServerError,
		)
	}

	templates.ExecuteTemplate(w, "Update", *product)
}

func AddProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	if r.Method == "POST" {
		price, err := strconv.ParseFloat(r.FormValue("price"), 64)
		if err != nil {
			http.Error(w, "Error parsing price", http.StatusBadRequest)
			return
		}
		quantity, err := strconv.Atoi(r.FormValue("quantity"))
		if err != nil {
			http.Error(w, "Error parsing quantity", http.StatusBadRequest)
			return
		}
		product = models.Product{
			Name:        r.FormValue("name"),
			Description: r.FormValue("description"),
			Price:       price,
			Quantity:    quantity,
		}

		err = models.AddProduct(product)
		if err != nil {
			http.Error(
				w,
				fmt.Sprintf("Error inserting data to DB %v", err),
				http.StatusInternalServerError)
			return
		}
	}

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	productId := params.Get("id")

	err := models.DeleteProduct(productId)
	if err != nil {
		http.Error(
			w,
			fmt.Sprintf("Error deleting product with id %s. Error: %v", productId, err),
			http.StatusInternalServerError,
		)
		return
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	if r.Method == "POST" {
		productId, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			http.Error(w, "Error parsing product id", http.StatusBadRequest)
			return
		}
		price, err := strconv.ParseFloat(r.FormValue("price"), 64)
		if err != nil {
			http.Error(w, "Error parsing price", http.StatusBadRequest)
			return
		}
		quantity, err := strconv.Atoi(r.FormValue("quantity"))
		if err != nil {
			http.Error(w, "Error parsing quantity", http.StatusBadRequest)
			return
		}
		product = models.Product{
			Id:          productId,
			Name:        r.FormValue("name"),
			Description: r.FormValue("description"),
			Price:       price,
			Quantity:    quantity,
		}

		err = models.UpdateProduct(product)
		if err != nil {
			http.Error(
				w,
				fmt.Sprintf("Error updating product into DB %v", err),
				http.StatusInternalServerError)
			return
		}
	}

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
