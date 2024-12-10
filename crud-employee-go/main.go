package main

import (
	"net/http"

	"github.com/dwiirawan/golang/crud-employee-go/database"
	"github.com/dwiirawan/golang/crud-employee-go/routes"
)

func main() {
	database.InitDatabase()

	server := http.NewServeMux()
	routes.MapRoutes(server)

	http.ListenAndServe(":8080", server)
}
