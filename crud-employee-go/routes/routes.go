package routes

import (
	"net/http"

	"github.com/dwiirawan/golang/crud-employee-go/controller"
)

func MapRoutes(server *http.ServeMux) {
	server.HandleFunc("/", controller.NewHelloWorldController())
	server.HandleFunc("/employee", controller.NewIndexEmployee())
}
