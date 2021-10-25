package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rest-api-market/service"
)

func CustomerRouter(api *gin.RouterGroup) {

	customer := *api.Group("/customer")

	customerService := service.NewCustomerService()

	customer.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": customerService.GetAll()})
	})

	customer.GET("/:id", func(c *gin.Context) {
		customer := customerService.GetById(c)
		if customer.Name == "" {
			c.JSON(http.StatusNotFound, gin.H{"data": "Not found"})
		} else {
			c.JSON(http.StatusOK, gin.H{"data": customer})
		}
	})

	customer.POST("/", func(c *gin.Context) {
		customerObj, err := customerService.Create(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"data": customerObj})
		}
	})

	customer.DELETE("/:id", func(c *gin.Context) {
		customerObj, err := customerService.Delete(c)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"data": "Not found"})
		} else {
			c.JSON(http.StatusOK, gin.H{"data": customerObj})
		}
	})
}
