package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

const (
	host     = "172.17.0.3"
	port     = 5432
	user     = "postgres"  // Usuário PostgreSQL
	password = "375217771" // Senha PostgreSQL
	dbname   = "web"       // Banco de dados PostgreSQL
)

type Doador struct {
	Codigo         int    `json:"codigo"`
	Nome           string `json:"nome"`
	CPF            string `json:"cpf"`
	Contato        string `json:"contato"`
	TipoDoSangue   string `json:"tipoSanguineo"`
	Rh             string `json:"fatorRh"`
	TipoRHCorretos *bool  `json:"tiporhcorretos"`
	Situacao       string `json:"situacao"`
}

var db *sql.DB

func main() {
	// Estabelece conexÃ£o com o banco de dados
	dbinfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	db, err = sql.Open("postgres", dbinfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Inicia o roteador
	router := mux.NewRouter()

	// Define as rotas CRUD
	router.HandleFunc("/informacoes", getTodos).Methods("POST")
	router.HandleFunc("/dados/{codigo}", getTodo).Methods("GET")
	router.HandleFunc("/dados", createTodo).Methods("POST")
	router.HandleFunc("/dados/{codigo}", updateTodo).Methods("PUT")
	router.HandleFunc("/dados/{codigo}", deleteTodo).Methods("DELETE")

	// Inicia o servidor na porta 8080
	//log.Fatal(http.ListenAndServe(":8081", router))
	// Adiciona o middleware CORS ao roteador
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:8080"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
	})

	// Adiciona o middleware CORS ao roteador
	handler := c.Handler(router)

	// Inicia o servidor na porta 8081
	log.Fatal(http.ListenAndServe(":8081", handler))
}

func getTodos(w http.ResponseWriter, r *http.Request) {
	var dadosNovos Doador
	_ = json.NewDecoder(r.Body).Decode(&dadosNovos)

	// Construção da consulta SQL
	query := "SELECT * FROM doador WHERE situacao =$1"
	parametros := []interface{}{"Ativo"} // Parâmetros para a consulta SQL

	// Verifica se as variáveis de pesquisa não estão vazias
	if dadosNovos.Nome != "" {
		query += " AND nome ILIKE $" + strconv.Itoa(len(parametros)+1)
		parametros = append(parametros, dadosNovos.Nome+"%")
	}
	if dadosNovos.CPF != "" {
		query += " AND cpf ILIKE $" + strconv.Itoa(len(parametros)+1)
		parametros = append(parametros, dadosNovos.CPF+"%")
	}
	if dadosNovos.Contato != "" {
		query += " AND contato ILIKE $" + strconv.Itoa(len(parametros)+1)
		parametros = append(parametros, dadosNovos.Contato+"%")
	}
	if dadosNovos.TipoDoSangue != "" {
		query += " AND tipo_sanguineo ILIKE $" + strconv.Itoa(len(parametros)+1)
		parametros = append(parametros, dadosNovos.TipoDoSangue+"%")
	}
	if dadosNovos.Rh != "" {
		query += " AND rh=$" + strconv.Itoa(len(parametros)+1)
		parametros = append(parametros, dadosNovos.Rh)
	}
	if dadosNovos.TipoRHCorretos != nil {
		query += " AND tipo_rh_corretos= $" + strconv.Itoa(len(parametros)+1)
		parametros = append(parametros, dadosNovos.TipoRHCorretos)
	}
	if dadosNovos.Codigo != -1 {
		query += " AND codigo=$" + strconv.Itoa(len(parametros)+1)
		parametros = append(parametros, dadosNovos.Codigo)
	}

	// Executa a consulta SQL
	rows, err := db.Query(query, parametros...)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Itera sobre os resultados
	var todos []Doador
	for rows.Next() {
		var novoDado Doador
		err := rows.Scan(&novoDado.Codigo, &novoDado.Nome, &novoDado.CPF, &novoDado.Contato, &novoDado.TipoDoSangue, &novoDado.Rh, &novoDado.TipoRHCorretos, &novoDado.Situacao)
		if err != nil {
			log.Fatal(err)
		}
		todos = append(todos, novoDado)
	}

	// Codifica os resultados como JSON e envia a resposta
	json.NewEncoder(w).Encode(todos)
}

func getTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	codigo := params["codigo"]
	var todo Doador
	err := db.QueryRow("SELECT * FROM doador WHERE codigo=$1%", codigo).Scan(&todo.Codigo, &todo.Nome, &todo.CPF, &todo.Contato, &todo.TipoDoSangue, &todo.Rh, &todo.TipoRHCorretos, &todo.Situacao)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(todo)
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	var novoDado Doador
	_ = json.NewDecoder(r.Body).Decode(&novoDado)
	_, err := db.Exec("INSERT INTO doador (nome,cpf,contato,tipo_sanguineo,rh,tipo_rh_corretos,situacao) VALUES ($1,$2,$3,$4,$5,$6,$7)", novoDado.Nome, novoDado.CPF, novoDado.Contato, novoDado.TipoDoSangue, novoDado.Rh, novoDado.TipoRHCorretos, novoDado.Situacao)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(novoDado)
}

func updateTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	codigo := params["codigo"]
	var novoDado Doador
	_ = json.NewDecoder(r.Body).Decode(&novoDado)
	_, err := db.Exec("UPDATE  doador SET nome=$1, cpf=$2, contato=$3, tipo_sanguineo=$4, rh=$5, tipo_rh_corretos=$6 WHERE codigo=$7", novoDado.Nome, novoDado.CPF, novoDado.Contato, novoDado.TipoDoSangue, novoDado.Rh, novoDado.TipoRHCorretos, codigo)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(novoDado)
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	codigo := params["codigo"]
	_, err := db.Exec("UPDATE doador SET situacao =$1 WHERE codigo=$2", "Inativo", codigo)
	if err != nil {
		log.Fatal(err)
	}
	w.WriteHeader(http.StatusNoContent)

}
