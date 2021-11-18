package service

import (
	"github.com/gin-gonic/gin"
)

type Repository interface {
	Create(c *gin.Context) (interface{}, error)
	Delete(c *gin.Context) (interface{}, error)
	Update(c *gin.Context) (interface{}, interface{})
	GetById(c *gin.Context) interface{}
	GetAll() []interface{}
}
