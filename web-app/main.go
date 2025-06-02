package main

import (
	"net/http"

	_ "github.com/lib/pq"
	"github.com/wes-santos/alura-golang-study/web-app/routes"
)

func main() {
	routes.LoadRoutes()
	http.ListenAndServe(":8000", nil)
}
