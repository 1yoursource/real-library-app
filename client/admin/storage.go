package admin

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"lib-client-server/client/models"
	"lib-client-server/client/type_getter"
	"lib-client-server/database"
)

type (
	Storage struct {
		collection string
		*mgo.Database
	}
)

func CreateConnect(host, name, cname string) models.StorageInterface {
	return &Storage{collection:cname,Database:database.Connect(host, name, cname)}
}

func (s *Storage) GetAll() *mgo.Query {
	return s.C(s.collection).Find(nil)
}

func (s *Storage) Set(data interface{}) {
	if book, isIt:=type_getter.GetTypeBook(data); isIt {
		//todo save to bd
		if err := s.C(s.collection).Insert(book); err != nil {
			fmt.Printf("bookStorage Insert error = %s; data = %+v\n", err,book)
		}
	}
}

func (s *Storage) Get(key uint64) (interface{}, error) {
	return nil, nil
}

func (s *Storage) GetByQuery(query interface{}) *mgo.Query {
	return s.C(s.collection).Find(query)
}

func (s *Storage) Update(selector models.Obj, updater models.Obj) error {
	return nil
}

func (s *Storage) Delete(key uint64) error {
	return s.C(s.collection).RemoveId(key)
}