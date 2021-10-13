package routes

import "github.com/gin-gonic/gin"

func ProductRouter(api *gin.RouterGroup) {

	product := *api.Group("/product")

	product.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

}
