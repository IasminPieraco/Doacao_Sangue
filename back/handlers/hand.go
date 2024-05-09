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
	result, err := h.DB.Exec("INSERT INTO doador (nome,cpf,contato,tipo_sanguineo,rh,tipo_rh_corretos,situacao) VALUES ($1,$2,$3,$4,$5,$6,$7)", novoDado.Nome, novoDado.CPF, novoDado.Contato, novoDado.TipoDoSangue, novoDado.Rh, novoDado.TipoRHCorretos, novoDado.Situacao)
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
	// Realizar a consulta ao banco de dados para buscar as doações
	rows, err := h.DB.Query("SELECT codigo, nome, cpf, contato, tipo_sanguineo, rh, tipo_rh_corretos, situacao FROM doador")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	// Converter os resultados da consulta em uma lista de doadores
	var doadores []models.Doador
	for rows.Next() {
		var doador models.Doador
		err := rows.Scan(&doador.Codigo, &doador.Nome, &doador.CPF, &doador.Contato, &doador.TipoDoSangue, &doador.Rh, &doador.TipoRHCorretos, &doador.Situacao)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		doadores = append(doadores, doador)
	}

	// Fechar a consulta
	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Retornar os doadores como uma resposta JSON
	c.JSON(http.StatusOK, gin.H{
		"mensagem": "Lista de doadores",
		"doadores": doadores,
	})
}
