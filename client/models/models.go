package models

import "gopkg.in/mgo.v2/bson"

type(
	StorageInterface interface {
		Set(data interface{})
		Get(key bson.ObjectId) (interface{}, bool)
		GetAll() []interface{}
		Update(data interface{}, query ...Obj) error
		Delete(key bson.ObjectId)
	}

	Obj map[string]interface{}
)
