package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rest-api-market/model"
	"github.com/rest-api-market/service"
)

func CategoryRouter(api *gin.RouterGroup) {

	category := *api.Group("/category")
	var categoryService service.CategoryService

	category.GET("/", func(c *gin.Context) {
		T := service.Repository.GetAll(&categoryService)
		categories := T.([]model.Category)
		c.JSON(http.StatusOK, gin.H{"data": categories})
	})

	category.GET("/:id", func(c *gin.Context) {
		T := service.Repository.GetById(&categoryService, c)
		category := T.(model.Category)
		if category.Name == "" {
			c.JSON(http.StatusNotFound, gin.H{"data": "Not found"})
		} else {
			c.JSON(http.StatusOK, gin.H{"data": category})
		}
	})

	category.POST("/", func(c *gin.Context) {
		T, err := service.Repository.Create(&categoryService, c)
		category := T.(model.Category)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
		} else {
			c.JSON(http.StatusCreated, gin.H{"data": category})
		}
	})

	//it works, but i need to improve this method
	category.PATCH("/:id", func(c *gin.Context) {
		T, status := service.Repository.Update(&categoryService, c)
		category := T.(model.Category)
		switch status {
		case 1:
			c.JSON(http.StatusNoContent, gin.H{"data": "Empty request body"})
		case 2:
			c.JSON(http.StatusNotFound, gin.H{"data": "Not found"})
		case 3:
			c.JSON(http.StatusExpectationFailed, gin.H{"data": "Attributes do not match"})
		case 4:
			c.JSON(http.StatusOK, gin.H{"data": category})
		}
	})

	category.DELETE("/:id", func(c *gin.Context) {
		T, err := service.Repository.Delete(&categoryService, c)
		category := T.(model.Category)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"data": "Not found"})
		} else {
			c.JSON(http.StatusOK, gin.H{"data": category})
		}
	})
}
