package models

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type(
	StorageInterface interface {
		Set(data interface{})
		Get(key bson.ObjectId) (interface{}, error)
		GetByQuery(query interface{}) (interface{}, error)
		GetAll() *mgo.Query
		Update(data interface{}, query ...Obj) error
		Delete(key bson.ObjectId)
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