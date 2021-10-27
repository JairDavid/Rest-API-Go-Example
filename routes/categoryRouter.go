package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rest-api-market/service"
)

func CategoryRouter(api *gin.RouterGroup) {
	category := *api.Group("/category")
	categoryService := service.NewCategoryRepository()

	category.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": categoryService.GetAll()})
	})

	category.GET("/:id", func(c *gin.Context) {
		categoryObj := categoryService.GetbyId(c)
		if categoryObj.Name == "" {
			c.JSON(http.StatusNotFound, gin.H{"data": "Not found"})
		} else {
			c.JSON(http.StatusOK, gin.H{"data": categoryObj})
		}
	})

	category.POST("/", func(c *gin.Context) {
		categoryObj, err := categoryService.Create(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
		} else {
			c.JSON(http.StatusCreated, gin.H{"data": categoryObj})
		}
	})

	//testing it does not work yet
	category.PATCH("/:id", func(c *gin.Context) {
		categoryService.Update(c)
	})

	category.DELETE("/:id", func(c *gin.Context) {
		categoryObj, err := categoryService.Delete(c)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"data": "Not found"})
		} else {
			c.JSON(http.StatusOK, gin.H{"data": categoryObj})
		}
	})
}
