package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"root/handlers"

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

func main() {
	// Estabelecer a conexão com o banco de dados
	err := conectarBD()
	if err != nil {
		fmt.Printf("Erro ao conectar ao banco de dados: %v\n", err)
		return
	}

	defer db.Close()

	hander := handlers.NewHandler(db)

	r := gin.Default()

	// Rota para inserir dados no banco de dados
	r.POST("/dados", hander.InserirDado)

	// Rota de ping de exemplo
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	r.Run()
}
