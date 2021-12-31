package models

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
)

type(
	StorageInterface interface {
		Set(data interface{})
		Get(key uint64) (interface{}, error)
		GetByQuery(query interface{}) *mgo.Query
		GetAll() *mgo.Query
		Update(data interface{}, query ...Obj) error
		Delete(key uint64) error
	}
	BookInterface interface {
		Handler(c *gin.Context)
		GetAll(c *gin.Context)
		Create(c *gin.Context) // a:cr,u:get
		Read(c *gin.Context)   // a:see,u:see info //todo NEED ?????
		Update(c *gin.Context) // a:edit,u:
		Delete(c *gin.Context) // a:del,u:return
	}

	Obj map[string]interface{}
)

const (
	CoockieMaxAge = 32000000
)