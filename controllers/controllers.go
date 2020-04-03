package controllers

import (
	"encoding/json"
	"go-medium/models"
	"log"
	"net/http"
	"time"

	"go-medium/repositories"

	"github.com/gorilla/mux"
)

func HomeController(res http.ResponseWriter, req *http.Request) {
	json.NewEncoder(res).Encode("Medium - Home")
}

func ListPostsController(res http.ResponseWriter, req *http.Request) {
	var posts = repositories.ListAllPosts()

	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(posts)
}

func AddPostController(res http.ResponseWriter, req *http.Request) {

	params := mux.Vars(req)

	var post models.Post
	_ = json.NewDecoder(req.Body).Decode(&post)

	var id = len(repositories.ListAllPosts()) + 1

	log.Println(params)

	post.ID = int64(id)
	post.Created = time.Now()

	repositories.AddPost(post)

	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(post)
}
