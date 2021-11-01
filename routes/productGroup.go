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
		productobj := productService.GetById(c)
		if productobj.Name == "" {
			c.JSON(http.StatusNotFound, gin.H{"data": "Not found"})
		} else {
			c.JSON(http.StatusOK, gin.H{"data": productobj})
		}
	})

	product.POST("/", func(c *gin.Context) {
		productobj, err := productService.Create(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
		} else {
			c.JSON(http.StatusCreated, gin.H{"data": productobj})
		}
	})

	product.PATCH("/:id", func(c *gin.Context) {
		productobj, status := productService.Update(c)
		switch status {
		case 1:
			c.JSON(http.StatusNoContent, gin.H{"data": "Empty request body"})
		case 2:
			c.JSON(http.StatusNotFound, gin.H{"data": "Not found"})
		case 3:
			c.JSON(http.StatusExpectationFailed, gin.H{"data": "Attributes do not match"})
		case 4:
			c.JSON(http.StatusOK, gin.H{"data": productobj})
		}
	})

	product.DELETE("/:id", func(c *gin.Context) {
		productobj, err := productService.Delete(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"data": productobj})
		}

	})
}
