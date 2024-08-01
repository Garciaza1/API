package router

import (
    "github.com/gorilla/mux"
    "main/handlers"
)

func NewRouter() *mux.Router {
    router := mux.NewRouter()

	// USERS
    router.HandleFunc("/Users/GetAll", handlers.GetAllUsers).Methods("GET")
    router.HandleFunc("/Users/FetchUser/{email}", handlers.GetUserEmail).Methods("GET")
    router.HandleFunc("/Users/GetUser/{id}", handlers.GetUser).Methods("GET")
    router.HandleFunc("/Users/Post/Cadastro", handlers.CreateUser).Methods("POST")
    router.HandleFunc("/Users/Post/Login", handlers.LoginUser).Methods("POST")
    router.HandleFunc("/Users/Put/EditUser", handlers.UpdateUser).Methods("PUT")
    router.HandleFunc("/users/Delete/{id}", handlers.DeleteUser).Methods("DELETE")
    
    // PRODUTOS 
    router.HandleFunc("/Products/Post/Create", handlers.CreateProduto).Methods("POST")
    router.HandleFunc("/Products/GetAll", handlers.GetAllProducts).Methods("GET")
    router.HandleFunc("/Product/Get/{id}", handlers.GetProduto).Methods("GET")
    router.HandleFunc("/Product/Get/Deleted/{id}", handlers.GetDeletedProduto).Methods("GET")
    router.HandleFunc("/User/Products/Get/Deleted/{user_id}", handlers.GetDeletedByUser).Methods("GET")
    // router.HandleFunc("/Users/GetUser/{id}", handlers.GetDeletedProduto).Methods("GET") pegar por vendedor deleted
    router.HandleFunc("/Products/Put/Edit", handlers.UpdateProduto).Methods("PUT")
    // router.HandleFunc("/Products/Restaurar/${id}", handlers.UpdateProduto).Methods("PUT") RESTAURAR 
    router.HandleFunc("/Products/Delete/{id}", handlers.SoftDeleteProduct).Methods("DELETE")
	
	//VENDAS
    
    return router
}