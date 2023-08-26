package controllers

import (
	"github.com/cristopherecruz/cardapio-online-backend/database"
	"github.com/cristopherecruz/cardapio-online-backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func BuscarIngredientes(c *gin.Context) {

	var p []models.Produto
	dados := database.Db
	dados.Preload("Ingredientes").Find(&p)

	c.JSON(http.StatusOK, p)
}

func CriarIngrediente(c *gin.Context) {

	var ingrediente models.Ingrediente

	err := c.ShouldBindJSON(&ingrediente)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	database.Db.Create(&ingrediente)

	c.JSON(http.StatusOK, ingrediente)
}
