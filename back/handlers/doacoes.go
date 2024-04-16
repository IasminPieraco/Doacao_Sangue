package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BuscarDoacoes(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"mensagem": "Lista de doações",
		//"doacoes": []string{"Doação 1", "Doação 2", "Doação 3"},
	})

}
