package service

import (
	"github.com/gin-gonic/gin"
	"github.com/rest-api-market/connection"
	"github.com/rest-api-market/model"
)

type ProductRepository interface {
	Create(c *gin.Context) (model.Product, error)
	Delete(c *gin.Context) (model.Product, error)
	Update(c *gin.Context) (model.Product, interface{})
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

func (ps *productService) Update(c *gin.Context) (model.Product, interface{}) {
	var dinamicAtributes map[string]interface{}
	var product model.Product
	c.Bind(&dinamicAtributes)

	if dinamicAtributes == nil {
		//case 1: there's nothing in the request body
		return product, 1
	}
	connection.GetConnection().Find(&product, c.Param("id"))
	if product.ID == 0 {
		//case 2: inexistent ID
		return product, 2
	}

	statusUpdate := connection.GetConnection().Model(&model.Product{}).Where(c.Param("id")).Updates(dinamicAtributes)

	if statusUpdate.RowsAffected == 0 {
		//case 3: there's no matching data from the request body with the database attributes
		return product, 3
	} else {
		connection.GetConnection().Find(&product, c.Param("id"))
	}
	//case 4: Correct operation
	return product, 4

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
