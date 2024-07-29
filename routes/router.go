package router

import (
    "github.com/gorilla/mux"
    "main/handlers"
)

func NewRouter() *mux.Router {
    router := mux.NewRouter()
	// USERS
    router.HandleFunc("/users", handlers.GetAllUsers).Methods("GET")
    router.HandleFunc("/users/{id}", handlers.GetUser).Methods("GET")
    router.HandleFunc("/users", handlers.CreateUser).Methods("POST")
    router.HandleFunc("/users/{id}", handlers.UpdateUser).Methods("PUT")
    router.HandleFunc("/users/{id}", handlers.DeleteUser).Methods("DELETE")
    // PRODUTOS 
	
	return router
}