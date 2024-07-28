package router

import (
    "github.com/gorilla/mux"
    "main/handlers"
)

func NewRouter() *mux.Router {
    r := mux.NewRouter()
    r.HandleFunc("/users", handlers.GetAllUsers).Methods("GET")
    r.HandleFunc("/users/{id}", handlers.GetUser).Methods("GET")
    r.HandleFunc("/users", handlers.CreateUser).Methods("POST")
    r.HandleFunc("/users/{id}", handlers.UpdateUser).Methods("PUT")
    r.HandleFunc("/users/{id}", handlers.DeleteUser).Methods("DELETE")
    return r
}