package main

import (
	"log"
	"net/http"
	"time"

	"go-medium/models"
	"go-medium/repositories"
	"go-medium/routers"
)

func main() {
	log.Println("Populando banco de dados...")
	populateDatabase()

	log.Println("Iniciando servidor web...")
	log.Fatal(http.ListenAndServe(":8080", routers.GetRouter()))
}

func populateDatabase() {
	repositories.AddPost(models.Post{ID: 1, Text: "Hello, World!", Created: time.Now()})
}
