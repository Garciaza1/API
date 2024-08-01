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
	log.Println("Starting server on :5000")
	log.Fatal(http.ListenAndServe(":5000", r))
}
