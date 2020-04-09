package repositories

import (
	"github.com/felipefrbr/medium/models"
)

var posts []models.Post

func ListAllPosts() []models.Post {
	return posts
}

func GetPostById(id int64) models.Post {
	for _, post := range posts {
		if post.ID == id {
			return post
		}
	}
	return models.Post{}
}

func AddPost(post models.Post) {
	posts = append(posts, post)
}

func UpdatePost(post models.Post) {
	for i, item := range posts {
		if item.ID == post.ID {
			posts[i] = post
		}
	}

}
