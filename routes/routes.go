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
	router.GET("/api/personalidades", controllers.BuscarTodosProdutos)
	/*router.GET("/api/personalidades/:id", controllers.BuscarPersonalidadePorId)
	router.POST("/api/personalidades", controllers.CriarPersonalidade)
	router.DELETE("/api/personalidades/:id", controllers.DeletarPersonalidade)
	router.PUT("/api/personalidades/:id", controllers.EditarPersonalidade)*/

	log.Fatal(router.Run())

}
