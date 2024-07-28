package main

import (
	// "fmt"
	"log"
	"net/http"
	"main/db"
	"main/routes"
)

func main() {
	db.Init()

	r := router.NewRouter()
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
