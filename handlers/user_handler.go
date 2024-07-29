package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"main/db"
	"main/models"
)

// funções / controller do user

//ANOTAÇÃO:
/*
	PONTEIROS EM GO:
	& = LOCAL DA MEMORIA
	* VALOR DA VARIAVEL QUE IRA APONTAR PARA ALGUM ENDEREÇO
*/

// GetAllUsers retorna todos os usuários
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	// declara a variavel que vai receber os dados do tipo models usuario como array
	var users []models.Usuario

	//query personalizada
	query := "SELECT * FROM usuario"

	//chamando e resgatando dados do banco
	result := db.DB.Raw(query).Scan(&users)
	//handling de erros
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	//respondendo o pedido
	json.NewEncoder(w).Encode(users)
}

// GetUser retorna um usuário por ID
func GetUser(w http.ResponseWriter, r *http.Request) {
	// Obtém os parâmetros da URL
	vars := mux.Vars(r)

	// Converte o valor do parâmetro `id` de string para inteiro
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Declara uma variável para armazenar o usuário
	var user models.Usuario

	// faz a query e consulta o banco
	query := "SELECT * FROM usuario WHERE id = $1"
	result := db.DB.Raw(query, id).Scan(&user)

	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	if result.RowsAffected == 0 {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}

// CreateUser cria um novo usuário
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.Usuario
	// tira de json e coloca em struct
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := "INSERT INTO usuario (nome, senha, email, tipo, tel, endereco, cpf, cep) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id"
	result := db.DB.Raw(query, user.Nome, user.Senha, user.Email, user.Tipo, user.Tel, user.Endereco, user.CPF, user.CEP).Scan(&user.ID)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	//tira de struct e coloca em json
	json.NewEncoder(w).Encode(user)
}

// UpdateUser atualiza um usuário por ID
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user models.Usuario
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := "UPDATE usuario SET nome = $1, senha = $2, email = $3, tipo = $4, tel = $5, endereco = $6, cpf = $7, cep = $8 WHERE id = $9"
	result := db.DB.Exec(query, user.Nome, user.Senha, user.Email, user.Tipo, user.Tel, user.Endereco, user.CPF, user.CEP, user.ID)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	if result.RowsAffected == 0 {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}

// DeleteUser deleta um usuário por ID
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	// pega o id da requisição  
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }

    query := "UPDATE usuario SET deleted_at = NOW() WHERE id = $1"
    result := db.DB.Exec(query, id)
    if result.Error != nil {
        http.Error(w, result.Error.Error(), http.StatusInternalServerError)
        return
    }

    if result.RowsAffected == 0 {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }

    w.WriteHeader(http.StatusAccepted)
}

