package routers

import (
	"github.com/felipefrbr/medium/controllers"
	"github.com/gorilla/mux"
)

// Registra as rotas da aplicacao
func GetRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", controllers.HomeController).Methods("GET")
	router.HandleFunc("/posts", controllers.ListPostsController).Methods("GET")
	router.HandleFunc("/post/{id}", controllers.GetPostController).Methods("GET")
	router.HandleFunc("/post", controllers.AddPostController).Methods("POST")
	return router
}
