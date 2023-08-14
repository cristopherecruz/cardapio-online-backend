package controllers

import (
	"github.com/cristopherecruz/cardapio-online-backend/database"
	"github.com/cristopherecruz/cardapio-online-backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Registro(c *gin.Context) {

	var usuario models.Usuario

	if err := c.ShouldBindJSON(&usuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.Db.Create(&usuario)

	c.JSON(http.StatusOK, gin.H{"message": "Usuário registrado com sucesso!"})

}

func Login(c *gin.Context) {

	var usuario models.Usuario

	if err := c.ShouldBindJSON(&usuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.Usuario{}

	u.NomeUsuario = usuario.NomeUsuario
	u.Senha = usuario.Senha

	token, err := models.ChecarLogin(u.NomeUsuario, u.Senha, database.Db)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Usuário ou senha incorretos!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})

}
