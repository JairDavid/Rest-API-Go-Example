package routes

import "github.com/gin-gonic/gin"

func CustomerRouter(api *gin.RouterGroup) {

	customer := *api.Group("/customer")

	customer.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
}
