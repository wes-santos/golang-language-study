package main

import (
	"fmt"

	"github.com/wes-santos/rest-api/database"
	"github.com/wes-santos/rest-api/routes"
)

func main() {
	fmt.Println("Starting rest server with Go")
	database.ConnectWithDatabase()
	routes.HandleRequest()
}
