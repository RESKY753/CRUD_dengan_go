package routes

import (
		"net/http"
	    "github.com/RESKY753/CRUD_GO/controller"
		"database/sql"
)

func MapRoutes(server *http.ServeMux, db *sql.DB) {
	// serve.HandleFunc(server.HandleFunc("/",controller.NewHelloWordController() ))
	server.HandleFunc("/", controller.NewHelloWordController())
	server.HandleFunc("/employee", controller.NewIndexEmployee(db))
	server.HandleFunc("/employee/create", controller.NewCreateEmployeeController(db))
	server.HandleFunc("/employee/update", controller.NewUpdateEmployeeController(db))
	server.HandleFunc("/employee/delete", controller.NewDeleteEmployeeController(db))
}