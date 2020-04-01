package repositories

import (
	"go-medium/models"
)

var posts []models.Post

func ListAllPosts() []models.Post {
	return posts
}

func AddPost(post models.Post) {
	posts = append(posts, post)
}
