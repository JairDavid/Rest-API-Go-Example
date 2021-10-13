package routes

import "github.com/gin-gonic/gin"

func CustomerProductRouter(api *gin.RouterGroup) {

	customerProduct := *api.Group("/customerProduct")

	customerProduct.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
}
