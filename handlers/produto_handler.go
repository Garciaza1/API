package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"main/db"
	"main/models"

	"github.com/gorilla/mux"
)

func GetAllProducts(w http.ResponseWriter, r *http.Request) {

	var produtos []models.Produto

	query := "SELECT * FROM produto"
	result := db.DB.Raw(query).Scan(&produtos)

	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(produtos)
}

func GetProduto(w http.ResponseWriter, r *http.Request) {
	// captura o id dos params
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid product id", http.StatusBadRequest)
		return
	}

	var produto models.Produto
	query := "SELECT * FROM produto WHERE id = $1"

	result := db.DB.Raw(query, id).Scan(&produto)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	if result.RowsAffected == 0 {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	log.Println("resultados encontrados")
	log.Print(result)

	json.NewEncoder(w).Encode(produto)
}

func GetDeletedProduto(w http.ResponseWriter, r *http.Request) {
	// captura o id dos params
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["product_id"])
	if err != nil {
		http.Error(w, "Invalid product id", http.StatusBadRequest)
		return
	}

	var produto models.Produto
	query := "SELECT * FROM produto WHERE id = $1 AND deleted_at IS NOT NULL"

	result := db.DB.Raw(query, id).Scan(&produto)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	if result.RowsAffected == 0 {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	log.Println("resultados encontrados")
	log.Print(result)

	json.NewEncoder(w).Encode(produto)
}

func getUserDeletedProduct(w http.ResponseWriter, r *http.Request) {
	
	// captura o id dos params
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["user_id"])
	if err != nil {
		http.Error(w, "Invalid product id", http.StatusBadRequest)
		return
	}

	var produto models.Produto
	query := "SELECT * FROM produto WHERE user_id = ? AND deleted_at IS NOT NULL"

	result := db.DB.Raw(query, id).Scan(&produto)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	if result.RowsAffected == 0 {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	log.Println("resultados encontrados")
	log.Print(result)

	json.NewEncoder(w).Encode(produto)
}

func CreateProduto(w http.ResponseWriter, r *http.Request) {

	var produto models.Produto

	// tira de json e coloca em struct
	err := json.NewDecoder(r.Body).Decode(&produto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := "INSERT INTO produto " +
		"(user_id, nome, descricao, imagem,  preco, quantidade, codigo, garantia, categoria, marca)" +
		" VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id"
	result := db.DB.Raw(query, produto.UserID, produto.Nome, produto.Descricao, produto.Imagem, produto.Preco, produto.Quantidade, produto.Codigo, produto.Garantia, produto.Categoria, produto.Marca)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Instância de produto criada com sucesso!")
	log.Print(result)

	json.NewEncoder(w).Encode(produto)
}

func UpdateProduto(w http.ResponseWriter, r *http.Request) {

	var produto models.Produto
	err := json.NewDecoder(r.Body).Decode(&produto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := "UPDATE produto SET nome = $1, descricao = $2, imagem = $3, preco = $4, quantidade = $5, codigo = $6, garantia = $7, categoria = $8, marca = $9 WHERE id = $10"
	result := db.DB.Exec(query, produto.Nome, produto.Descricao, produto.Imagem, produto.Preco, produto.Quantidade, produto.Codigo, produto.Garantia, produto.Categoria, produto.Marca, produto.ID)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	if result.RowsAffected == 0 {
		http.Error(w, "Produto not found!\n", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(produto)
}

func SoftDeleteProduct(w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	
	if err != nil{
		http.Error(w, "Product Id Invalid!", http.StatusBadRequest)
		return
	}

    query := "UPDATE produto SET deleted_at = NOW() WHERE id = $1"
	result := db.DB.Exec(query, id)

	if result != nil{
		http.Error(w,  result.Error.Error(), http.StatusInternalServerError)
		return
	}

	// if result.RowsAffected == 0{
	// 	http.Error(w, "Produto not found!", http.StatusNotFound)
	// 	return
	// }

    w.WriteHeader(http.StatusAccepted)
}
