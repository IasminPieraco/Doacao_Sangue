package handlers

import (
	"database/sql"
	"net/http"
	"root/models"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	DB *sql.DB
}

func NewHandler(db *sql.DB) *Handler {
	return &Handler{DB: db}
}

func (h *Handler) InserirDado(c *gin.Context) {
	var novoDado models.Doador
	if err := c.BindJSON(&novoDado); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Inserir dados no banco de dados
	result, err := h.DB.Exec("INSERT INTO doador (nome) VALUES ($1)", novoDado.Nome)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Obter o ID do dado inserido
	id, _ := result.LastInsertId()

	// Retornar o ID do dado inserido
	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func (h *Handler) BuscarDoacoes(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"mensagem": "Lista de doações",
		//"doacoes": []string{"Doação 1", "Doação 2", "Doação 3"},
	})

}
