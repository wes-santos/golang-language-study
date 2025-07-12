package main

import (
	"github.com/wes-santos/gin-rest-api/database"
	"github.com/wes-santos/gin-rest-api/routes"
)

func main() {
	database.ConnectWithDatabase()
	routes.HandleRequests()
}
