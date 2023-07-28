package controllers

import (
	"github.com/cristopherecruz/cardapio-online-backend/database"
	"github.com/cristopherecruz/cardapio-online-backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Home(c *gin.Context) {
	c.String(http.StatusOK, "It works!")
}

func BuscarTodosProdutos(c *gin.Context) {

	var p []models.Produto
	dados := database.Db
	dados.Preload("Ingredientes").Find(&p)

	c.JSON(http.StatusOK, p)
}
