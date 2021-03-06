package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/felipefrbr/medium/models"
	"github.com/felipefrbr/medium/repositories"
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

func GetPostController(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	var postID, _ = strconv.ParseInt(params["id"], 10, 64)

	var post = repositories.GetPostById(postID)

	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(post)
}

func AddPostController(res http.ResponseWriter, req *http.Request) {

	var post models.Post
	_ = json.NewDecoder(req.Body).Decode(&post)

	var id = len(repositories.ListAllPosts()) + 1

	post.ID = int64(id)
	post.Created = time.Now()

	repositories.AddPost(post)

	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(post)
}

func UpdatePostController(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	var postID, _ = strconv.ParseInt(params["id"], 10, 64)

	post := repositories.GetPostById(postID)

	var postRequest models.Post
	_ = json.NewDecoder(req.Body).Decode(&postRequest)

	post.Text = postRequest.Text

	repositories.UpdatePost(post)

	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(post)
}
