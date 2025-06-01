package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func connectToDb() *sql.DB {
	connStr := "user=admin dbname=alura_store password=password host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err.Error())
	}
	return db
}

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

var templates = template.Must(template.ParseGlob("web-app/templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	db := connectToDb()
	defer db.Close()

	query := "SELECT id, name, description, price, quantity FROM products"
	rows, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var product Product
		err = rows.Scan(
			&product.Id,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.Quantity,
		)
		if err != nil {
			log.Println("Error scanning line: ", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		products = append(products, product)
	}
	if err = rows.Err(); err != nil {
		log.Println("Error after iterating lines: ", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	templates.ExecuteTemplate(w, "Index", products)
}
