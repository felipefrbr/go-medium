package repositories

import (
	"github.com/felipefrbr/medium/models"
)

var posts []models.Post

func ListAllPosts() []models.Post {
	return posts
}

func GetPostById(id int64) *models.Post {
	for _, post := range posts {
		if post.ID == id {
			return &post
		}
	}
	return nil
}

func AddPost(post models.Post) {
	posts = append(posts, post)
}
