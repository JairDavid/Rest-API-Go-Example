package service

import (
	"github.com/gin-gonic/gin"
	"github.com/rest-api-market/connection"
	"github.com/rest-api-market/model"
)

type CustomerService struct {
}

func NewCustomerService() Repository {
	return &CustomerService{}
}

func (cs *CustomerService) Create(c *gin.Context) (interface{}, error) {
	var customer model.Customer
	err := c.ShouldBindJSON(&customer)
	if err != nil {
		return customer, err
	}
	connection.GetConnection().Create(&customer)
	return customer, nil
}

func (cs *CustomerService) Delete(c *gin.Context) (interface{}, error) {
	var customer model.Customer
	connection.GetConnection().Find(&customer, c.Param("id"))
	err := c.ShouldBindJSON(&customer)
	if err != nil {
		return customer, err
	}
	//Soft Delete
	//connection.GetConnection().Delete(&customer, customer.ID)
	connection.GetConnection().Unscoped().Delete(&customer, customer.ID)
	return customer, err
}

func (cs *CustomerService) Update(c *gin.Context) (interface{}, interface{}) {
	var dinamicAtributes map[string]interface{}
	var customer model.Customer
	c.Bind(&dinamicAtributes)

	if dinamicAtributes == nil {
		//case 1: there's nothing in the request body
		return customer, 1
	}
	connection.GetConnection().Find(&customer, c.Param("id"))
	if customer.ID == 0 {
		//case 2: inexistent ID
		return customer, 2
	}

	statusUpdate := connection.GetConnection().Model(&model.Customer{}).Where(c.Param("id")).Updates(dinamicAtributes)

	if statusUpdate.RowsAffected == 0 {
		//case 3: there's no matching data from the request body with the database attributes
		return customer, 3
	} else {
		connection.GetConnection().Find(&customer, c.Param("id"))
	}
	//case 4: Correct operation
	return customer, 4

}

func (cs *CustomerService) GetById(c *gin.Context) interface{} {
	var customer model.Customer
	connection.GetConnection().Find(&customer, c.Param("id"))
	return customer
}

func (cs *CustomerService) GetAll() interface{} {
	var customers []model.Customer
	connection.GetConnection().Preload("CustomerProducts").Find(&customers)
	return customers
}
