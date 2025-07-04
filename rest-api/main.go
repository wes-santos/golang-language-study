package main

import (
	"fmt"

	"github.com/wes-santos/rest-api/models"
	"github.com/wes-santos/rest-api/routes"
)

func main() {
	models.Personalities = []models.Personality{
		{Name: "Name 1", History: "History 1"},
		{Name: "Name 2", History: "History 2"},
	}

	fmt.Println("Starting rest server with Go")
	routes.HandleRequest()
}
