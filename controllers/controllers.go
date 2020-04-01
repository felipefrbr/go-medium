package controllers

import (
	"encoding/json"
	"net/http"

	"go-medium/repositories"
)

func HomeController(res http.ResponseWriter, req *http.Request) {
	json.NewEncoder(res).Encode("Medium - Home")
}

func ListPostsController(res http.ResponseWriter, req *http.Request) {
	var posts = repositories.ListAllPosts()
	json.NewEncoder(res).Encode(posts)
}
