package routers

import (
	"go-medium/controllers"

	"github.com/gorilla/mux"
)

// Registra as rotas da aplicacao
func GetRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", controllers.HomeController).Methods("GET")
	router.HandleFunc("/posts", controllers.ListPostsController).Methods("GET")
	router.HandleFunc("/post", controllers.AddPostController).Methods("POST")
	return router
}
