package service

import (
	"github.com/gin-gonic/gin"
)

type Repository interface {
	Create(c *gin.Context) (interface{}, error)
	Delete(c *gin.Context) (interface{}, error)
	Update(c *gin.Context) (interface{}, interface{})
	GetById(c *gin.Context) interface{}
	GetAll() interface{}
}

func Create(r Repository, c *gin.Context) {
	r.Create(c)
}

func Delete(r Repository, c *gin.Context) {
	r.Delete(c)
}

func Update(r Repository, c *gin.Context) {
	r.Update(c)
}

func GetById(r Repository, c *gin.Context) {
	r.GetById(c)
}

func GetAll(r Repository) {
	r.GetAll()
}
