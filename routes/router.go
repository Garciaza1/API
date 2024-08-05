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


// PRODUTOS:
    // posts
	router.HandleFunc("/Products/Post/Create", handlers.CreateProduto).Methods("POST") // *
    // gets
	router.HandleFunc("/Products/GetAll", handlers.GetAllProducts).Methods("GET") // *
	router.HandleFunc("/Product/Get/{id}", handlers.GetProduto).Methods("GET") // *
	router.HandleFunc("/Product/Get/Deleted/{id}", handlers.GetDeletedProduto).Methods("GET") // *
	router.HandleFunc("/User/Products/Get/Deleted/{user_id}", handlers.GetDeletedByUser).Methods("GET") // *
    
    // puts
	router.HandleFunc("/Products/Put/Edit", handlers.UpdateProduto).Methods("PUT") // *
	router.HandleFunc("/Products/Restaurar/${id}", handlers.RestaurarProduto).Methods("PUT") // * 
    
    // delets
	router.HandleFunc("/Products/Delete/{id}", handlers.SoftDeleteProduct).Methods("DELETE") // *


//VENDAS
    //posts
    router.HandleFunc("/Vendas/Post/Compra", handlers.CreateVenda).Methods("POST")
	    // router.post('/Vendas/Post/Compra', venda.createVenda);
    
    //puts
    router.HandleFunc("/Vendas/Put/Edit", handlers.EditVenda).Methods("PUT")
	    // router.put('/Vendas/Put/Edit', venda.EditVenda);
    router.HandleFunc("/Vendas/Put/Restaurar", handlers.RestauraVenda).Methods("PUT")
	    // router.put('/Vendas/Put/Restaurar', venda.restaurarVenda);
    router.HandleFunc("/Vendas/Put/Cancelar", handlers.CancelarVenda).Methods("PUT")
	    // router.put('/Vendas/Put/Cancelar', venda.CancelarVenda);
    
    //gets
    router.HandleFunc("/Vendas/Get/Compras/:user_id", handlers.GetComprasClient).Methods("GET")
	    // router.get('/Vendas/Get/Compras/:user_id', venda.getComprasClient);
    router.HandleFunc("/Vendas/GetAll", handlers.GetAll).Methods("GET")
	    // router.get('/Vendas/GetAll', venda.getAll);
    router.HandleFunc("/Vendas/Get/:id", handlers.GetVenda).Methods("GET")
	    // router.get('/Vendas/Get/:id', venda.getVenda);
    router.HandleFunc("/Vendas/Client/Get/:id", handlers.GetClientVendas).Methods("GET")
	    // router.get('/Vendas/Client/Get/:id', venda.getClientVenda);
    router.HandleFunc("/Vendas/Post/Compra", handlers.GetUserVendas).Methods("GET")
	    // router.get('/Vendas/User/Get/:user_id', venda.getUserVendas);
	
    //end of document
    return router
}
