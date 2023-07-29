package routes

import (
	"github.com/cristopherecruz/cardapio-online-backend/controllers"
	"github.com/cristopherecruz/cardapio-online-backend/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
)

func HandleRequest() {

	router := gin.Default()

	router.Use(cors.Default())

	router.Use(middleware.HeaderMiddleware)

	//Produtos
	router.GET("/", controllers.Home)
	router.GET("/api/produtos", controllers.BuscarTodosProdutos)

	//Combos
	router.GET("/api/combos", controllers.BuscarTodosCombos)
	router.POST("/api/combos", controllers.CriarCombo)
	/*router.POST("/api/personalidades", controllers.CriarPersonalidade)
	router.DELETE("/api/personalidades/:id", controllers.DeletarPersonalidade)
	router.PUT("/api/personalidades/:id", controllers.EditarPersonalidade)*/

	log.Fatal(router.Run())

}
