package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rest-api-market/model"
	"github.com/rest-api-market/service"
)

func ProductRouter(api *gin.RouterGroup) {

	product := *api.Group("/product")
	var productService service.ProductService

	product.GET("/", func(c *gin.Context) {
		T := service.Repository.GetAll(&productService)
		products := T.([]model.Product)
		c.JSON(http.StatusOK, gin.H{"data": products})
	})

	product.GET("/:id", func(c *gin.Context) {
		T := service.Repository.GetById(&productService, c)
		product := T.(model.Product)
		if product.Name == "" {
			c.JSON(http.StatusNotFound, gin.H{"data": "Not found"})
		} else {
			c.JSON(http.StatusOK, gin.H{"data": product})
		}
	})

	product.POST("/", func(c *gin.Context) {
		T, err := service.Repository.Create(&productService, c)
		product := T.(model.Product)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
		} else {
			c.JSON(http.StatusCreated, gin.H{"data": product})
		}
	})

	product.PATCH("/:id", func(c *gin.Context) {
		T, status := service.Repository.Update(&productService, c)
		products := T.([]model.Product)
		switch status {
		case 1:
			c.JSON(http.StatusNoContent, gin.H{"data": "Empty request body"})
		case 2:
			c.JSON(http.StatusNotFound, gin.H{"data": "Not found"})
		case 3:
			c.JSON(http.StatusExpectationFailed, gin.H{"data": "Attributes do not match"})
		case 4:
			c.JSON(http.StatusOK, gin.H{"data": products})
		}
	})

	product.DELETE("/:id", func(c *gin.Context) {
		T, err := service.Repository.Delete(&productService, c)
		product := T.(model.Product)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"data": product})
		}

	})
}
