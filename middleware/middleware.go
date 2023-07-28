package middleware

import "github.com/gin-gonic/gin"

func HeaderMiddleware(c *gin.Context) {
	//c.Header("Teste", "Teste")
	c.Next()
}
