package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rest-api-market/service"
)

func ProductRouter(api *gin.RouterGroup) {

	product := *api.Group("/product")
	productService := service.NewProductService()
	product.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": productService.GetAll()})
	})

	product.GET("/:id", func(c *gin.Context) {
		product := productService.GetById(c)
		if product.Name == "" {
			c.JSON(http.StatusNotFound, gin.H{"data": "Not found"})
		} else {
			c.JSON(http.StatusOK, gin.H{"data": product})
		}
	})

	product.POST("/", func(c *gin.Context) {
		product, err := productService.Create(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
		} else {
			c.JSON(http.StatusCreated, gin.H{"data": product})
		}
	})

	product.DELETE("/:id", func(c *gin.Context) {
		product, err := productService.Delete(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"data": product})
		}

	})
}
