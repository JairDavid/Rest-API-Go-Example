package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rest-api-market/model"
	"github.com/rest-api-market/service"
)

func CustomerRouter(api *gin.RouterGroup) {

	customer := *api.Group("/customer")
	var customerService service.CustomerService

	customer.GET("/", func(c *gin.Context) {
		T := service.Repository.GetAll(&customerService)
		customers := T.([]model.Customer)
		c.JSON(http.StatusOK, gin.H{"data": customers})
	})

	customer.GET("/:id", func(c *gin.Context) {
		T := service.Repository.GetById(&customerService, c)
		customer := T.(model.Customer)

		if customer.Name == "" {
			c.JSON(http.StatusNotFound, gin.H{"data": "Not found"})
		} else {
			c.JSON(http.StatusOK, gin.H{"data": customer})
		}
	})

	customer.POST("/", func(c *gin.Context) {
		T, err := service.Repository.Create(&customerService, c)
		customer := T.(model.Customer)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
		} else {
			c.JSON(http.StatusCreated, gin.H{"data": customer})
		}
	})

	customer.POST("/:id", func(c *gin.Context) {
		T, err := customerService.Buy(c)
		customer := T.(model.Customer)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
		} else {
			c.JSON(http.StatusCreated, gin.H{"data": customer})
		}
	})

	customer.PATCH("/:id", func(c *gin.Context) {
		T, status := service.Repository.Update(&customerService, c)
		customer := T.(model.Customer)
		switch status {
		case 1:
			c.JSON(http.StatusNoContent, gin.H{"data": "Empty request body"})
		case 2:
			c.JSON(http.StatusNotFound, gin.H{"data": "Not found"})
		case 3:
			c.JSON(http.StatusExpectationFailed, gin.H{"data": "Attributes do not match"})
		case 4:
			c.JSON(http.StatusOK, gin.H{"data": customer})
		}
	})

	customer.DELETE("/:id", func(c *gin.Context) {
		T, err := service.Repository.Delete(&customerService, c)
		customer := T.(model.Customer)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"data": "Not found"})
		} else {
			c.JSON(http.StatusOK, gin.H{"data": customer})
		}
	})
}
