package controllers

import (
	"github.com/cristopherecruz/cardapio-online-backend/database"
	"github.com/cristopherecruz/cardapio-online-backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func BuscarTodosCombos(c *gin.Context) {

	var cb []models.Combo
	dados := database.Db
	dados.Preload("Produtos").Find(&cb)

	c.JSON(http.StatusOK, cb)
}

func CriarCombo(c *gin.Context) {

	var combo models.Combo

	err := c.ShouldBindJSON(&combo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	database.Db.Create(&combo)

	c.JSON(http.StatusOK, combo)
}
