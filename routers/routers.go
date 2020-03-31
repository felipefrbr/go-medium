package routers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// Registra as rotas da aplicacao
func GetRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", MainHandler).Methods("GET")
	return router
}

func MainHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Medium")
}
