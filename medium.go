package main

import (
	"log"
	"net/http"
	"time"

	"github.com/felipefrbr/medium/models"
	"github.com/felipefrbr/medium/repositories"
	"github.com/felipefrbr/medium/routers"
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
