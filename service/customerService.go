package service

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rest-api-market/connection"
	"github.com/rest-api-market/model"
)

type CustomerRepository interface {
	Create(c *gin.Context) (model.Customer, error)
	Delete(c *gin.Context) (model.Customer, error)
	// Update(c *gin.Context)
	GetById(c *gin.Context) model.Customer
	GetAll() []model.Customer
	MultilCreate(*gin.Context)
}

type customerService struct {
}

func NewCustomerService() CustomerRepository {
	return &customerService{}
}

func (cs *customerService) Create(c *gin.Context) (model.Customer, error) {
	var customer model.Customer
	err := c.ShouldBindJSON(&customer)
	if err != nil {
		return customer, err
	}
	connection.GetConnection().Create(&customer)
	return customer, nil
}

func (ps *customerService) MultilCreate(c *gin.Context) {
	var customer []model.Customer
	err := c.ShouldBindJSON(&customer)
	log.Print(customer)
	if err != nil {
		log.Print(err)
	}

}

func (cs *customerService) Delete(c *gin.Context) (model.Customer, error) {
	var customer model.Customer
	connection.GetConnection().Find(&customer, c.Param("id"))
	err := c.ShouldBindJSON(&customer)
	if err != nil {
		return customer, err
	}
	connection.GetConnection().Delete(&customer, customer.ID)
	return customer, err
}

func (cs *customerService) GetById(c *gin.Context) model.Customer {
	var customer model.Customer
	connection.GetConnection().Find(&customer, c.Param("id"))
	return customer
}

func (cs *customerService) GetAll() []model.Customer {
	var customers []model.Customer
	connection.GetConnection().Find(&customers)
	return customers
}
