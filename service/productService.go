package service

import (
	"github.com/gin-gonic/gin"
	"github.com/rest-api-market/connection"
	"github.com/rest-api-market/model"
)

type ProductRepository interface {
	Create(c *gin.Context) (model.Product, error)
	Delete(c *gin.Context) (model.Product, error)
	GetById(c *gin.Context) model.Product
	GetAll() []model.Product
}

type productService struct {
}

func NewProductService() ProductRepository {
	return &productService{}
}

func (ps *productService) Create(c *gin.Context) (model.Product, error) {
	var product model.Product
	err := c.ShouldBindJSON(&product)
	if err != nil {
		return product, err
	}
	connection.GetConnection().Create(&product)
	return product, nil
}

func (ps *productService) Delete(c *gin.Context) (model.Product, error) {
	var product model.Product
	connection.GetConnection().Find(&product, c.Param("id"))
	err := c.ShouldBindJSON(&product)
	if err != nil {
		return product, err
	}
	connection.GetConnection().Unscoped().Delete(&product, product.ID)
	return product, nil
}

func (ps *productService) GetById(c *gin.Context) model.Product {
	var product model.Product
	connection.GetConnection().Find(&product, c.Param("id"))
	return product
}

func (ps *productService) GetAll() []model.Product {
	var products []model.Product
	connection.GetConnection().Find(&products)
	return products
}
