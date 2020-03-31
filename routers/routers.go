package routers

import (
	"go-medium/controllers"

	"github.com/gorilla/mux"
)

// Registra as rotas da aplicacao
func GetRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", controllers.HomeController).Methods("GET")
	return router
}
