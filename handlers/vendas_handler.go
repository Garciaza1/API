package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"main/db"
	"main/models"

	"github.com/gorilla/mux"
)

func CreateVenda(w http.ResponseWriter, r *http.Request) {

	var venda models.Venda

	err := json.NewDecoder(r.Body).Decode(&venda)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := "INSERT INTO vendas (endereco, num_residencia, cep, total, quantidade, mtd_pay, cpf, user_id, product_id, vendedor_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)"
	result := db.DB.Raw(query, venda.Endereco, venda.NumResidencia, venda.CEP, venda.Total, venda.Quantidade, venda.MtdPay, venda.CPF, venda.UserID, venda.ProdutoID, venda.VendedorID)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(venda)
}

func GetAll(w http.ResponseWriter, r *http.Request) {

	var vendas models.Venda

	query := " SELECT * FROM vendas"

	result := db.DB.Raw(query).Scan(&vendas)

	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(vendas)
}

func GetVenda(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(w, "id invalido!", http.StatusBadRequest)
		return
	}

	var venda models.Venda

	query := "SELECT * FROM vendas WHERE id = $1"
	result := db.DB.Raw(query, id).Scan(&venda)

	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	if result.RowsAffected == 0 {
		http.Error(w, "Venda not Found!", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(venda)
}

func GetClientVendas(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["user_id"])

	if err != nil {
		http.Error(w, "id invalido!", http.StatusBadRequest)
		return
	}

	var vendas models.Venda
	query := "SELECT * FROM vendas WHERE user_id = $1"

	result := db.DB.Raw(query, id).Scan(&vendas)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	if result.RowsAffected == 0 {
		http.Error(w, "venda not found!", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(vendas)
}

func getComprasClient(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["user_id"])

	if err != nil {
		http.Error(w, "id invalido!", http.StatusBadRequest)
		return
	}

	var vendas models.Venda
	query := "SELECT  vendas.*, produto.nome, produto.garantia, produto.imagem FROM vendas JOIN produto ON vendas.product_id = produto.id WHERE vendas.user_id = $1"

	result := db.DB.Raw(query, id).Scan(&vendas)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	if result.RowsAffected == 0 {
		http.Error(w, "venda not found!", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(vendas)
}

func getUserVendas(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["user_id"])

	if err != nil {
		http.Error(w, "id invalido!", http.StatusBadRequest)
		return
	}

	var vendas models.Venda
	query := "SELECT * FROM vendas WHERE vendedor_id = $1"

	result := db.DB.Raw(query, id).Scan(&vendas)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	if result.RowsAffected == 0 {
		http.Error(w, "venda not found!", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(vendas)
}

func RestauraVenda(w http.ResponseWriter, r *http.Request) {

	var input struct {
		ID         int `json:"id"`
		ProductID  int `json:"product_id"`
		Quantidade int `json:"quantidade"`
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if input.ID == 0 || input.ProductID == 0 || input.Quantidade == 0 {
		http.Error(w, "Dados inválidos na requisição", http.StatusBadRequest)
		return
	}

	var produto models.Produto
	var venda models.Venda

	tx := db.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Capturando a quantidade atual do produto
	if err := tx.Where("id = ?", input.ProductID).First(&produto).Error; err != nil {
		tx.Rollback()
		http.Error(w, "Produto não encontrado", http.StatusNotFound)
		return
	}

	if produto.Quantidade < input.Quantidade {
		tx.Rollback()
		http.Error(w, "Quantidade insuficiente no estoque", http.StatusBadRequest)
		return
	}

	// Atualizando o status da venda
	if err := tx.Model(&venda).Where("id = ?", input.ID).Update("status", "Confirmada").Error; err != nil {
		tx.Rollback()
		http.Error(w, "Não foi possível atualizar a venda", http.StatusInternalServerError)
		return
	}

	// Atualizando a quantidade do produto
	produto.Quantidade -= input.Quantidade
	if err := tx.Save(&produto).Error; err != nil {
		tx.Rollback()
		http.Error(w, "Não foi possível atualizar o produto", http.StatusInternalServerError)
		return
	}

	// Commit da transação
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		http.Error(w, "Não foi possível confirmar a transação", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Venda restaurada com sucesso!"})
}

func CancelarVenda(w http.ResponseWriter, r *http.Request) {
	var input struct {
		ID         int `json:"id"`
		ProductID  int `json:"product_id"`
		Quantidade int `json:"quantidade"`
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if input.ID == 0 || input.ProductID == 0 || input.Quantidade == 0 {
		http.Error(w, "Dados inválidos na requisição", http.StatusBadRequest)
		return
	}

	tx := db.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Atualizar status da venda para 'Cancelada'
	queryCancelaVenda := "UPDATE vendas SET sts_venda = 'Cancelada' WHERE id = ?"
	if err := tx.Exec(queryCancelaVenda, input.ID).Error; err != nil {
		tx.Rollback()
		http.Error(w, "Não foi possível atualizar a venda", http.StatusInternalServerError)
		return
	}

	// Atualizar a quantidade do produto
	queryVoltaProduto := "UPDATE produto SET quantidade = quantidade + ? WHERE id = ?"
	if err := tx.Exec(queryVoltaProduto, input.Quantidade, input.ProductID).Error; err != nil {
		tx.Rollback()
		http.Error(w, "Não foi possível atualizar o produto", http.StatusInternalServerError)
		return
	}

	// Commit da transação
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		http.Error(w, "Não foi possível confirmar a transação", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Venda cancelada com sucesso!"})
}

func editVenda(w http.ResponseWriter, r *http.Request) {

	var venda models.Venda

	err := json.NewDecoder(r.Body).Decode(&venda)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tx := db.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// ver se a venda existe!
	var existingVenda models.Venda
	if err := tx.First(&existingVenda, venda.ID).Error; err != nil {
		tx.Rollback()
		http.Error(w, "Venda não encontrada", http.StatusNotFound)
		return
	}

	//update da venda antes
	updateVendaQuery := "UPDATE vendas SET endereco = ?, num_residencia = ?, cep = ?, total = ?, quantidade = ?, mtd_pay = ?, cpf = ?, user_id = ?, product_id = ?, vendedor_id = ? WHERE id = ?"
	if err := tx.Exec(updateVendaQuery, venda.Endereco, venda.NumResidencia, venda.CEP, venda.Total, venda.Quantidade, venda.MtdPay, venda.CPF, venda.UserID, venda.ProdutoID, venda.VendedorID, venda.ID).Error; err != nil {
		tx.Rollback()
		http.Error(w, "Não foi possível atualizar a venda", http.StatusInternalServerError)
		return
	}

	// update de produto em seguida	
	updateProdutoQuery := "UPDATE produto SET quantidade = quantidade - ? WHERE id = ?"
	if err := tx.Exec(updateProdutoQuery, venda.Quantidade, venda.ID).Error; err != nil {
		tx.Rollback()
		http.Error(w, "Não foi possível atualizar o produto", http.StatusInternalServerError)
		return
	}

	// Commit da transação
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		http.Error(w, "Não foi possível confirmar a transação", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Venda atualizada com sucesso!"})

}
