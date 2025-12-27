package routes

import ("net/http"
	    "github.com/RESKY753/CRUD_GO/controller"
)

func MapRoutes(server *http.ServeMux) {
	// serve.HandleFunc(server.HandleFunc("/",controller.NewHelloWordController() ))
	server.HandleFunc("/", controller.NewHelloWordController())
	server.HandleFunc("/employee", controller.NewIndexEmployee())
}