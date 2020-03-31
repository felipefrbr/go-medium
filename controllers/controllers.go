package controllers

import (
	"encoding/json"
	"net/http"
)

func HomeController(res http.ResponseWriter, req *http.Request) {
	json.NewEncoder(res).Encode("Medium - Home")
}
