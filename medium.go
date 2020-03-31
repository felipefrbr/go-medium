package main

import (
	"log"
	"net/http"

	"go-medium/routers"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080", routers.GetRouter()))
}
