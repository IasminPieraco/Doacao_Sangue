package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var db *sql.DB // Variável global para o banco de dados

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

const (
	host     = "172.17.0.3"
	port     = 5432
	user     = "postgres"  // Usuário PostgreSQL
	password = "375217771" // Senha PostgreSQL
	dbname   = "web"       // Banco de dados PostgreSQL
)

func conectarBD() error {
	connStr := "user=postgres password=375217771 dbname=web sslmode=disable"
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("erro ao conectar ao banco de dados: %v", err)
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return fmt.Errorf("erro ao fazer ping no banco de dados: %v", err)
	}

	fmt.Println("Conexão com o banco de dados estabelecida com sucesso!")
	return nil
}

type Dado struct {
	ID   int    `json:"id"`
	Nome string `json:"nome"`
}

func inserirDado(c *gin.Context) {
	var novoDado Dado
	if err := c.BindJSON(&novoDado); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Inserir dados no banco de dados
	result, err := db.Exec("INSERT INTO tabela (nome) VALUES ($1)", novoDado.Nome)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Obter o ID do dado inserido
	id, _ := result.LastInsertId()

	// Retornar o ID do dado inserido
	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func main() {
	// Estabelecer a conexão com o banco de dados
	err := conectarBD()
	if err != nil {
		fmt.Printf("Erro ao conectar ao banco de dados: %v\n", err)
		return
	}
	defer db.Close()

	r := gin.Default()

	// Rota para inserir dados no banco de dados
	r.POST("/dados", inserirDado)

	// Rota de ping de exemplo
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	r.Run()
}
