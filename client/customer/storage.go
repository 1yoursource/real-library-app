package customer

import (
	"errors"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"lib-client-server/client"
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

func (s *Storage) GetByQuery(query interface{}) (interface{}, error) {
	if query == nil {
		return nil, errors.New("NIL QUERY")
	}
	var book client.Book
	if err := s.C(s.collection).Find(query).One(&book); err != nil {
		fmt.Printf("bookStorage Insert error = %s; data = %+v\n", err,book)
		return nil, err
	}
	return book, nil
}

func (s *Storage) Get(key bson.ObjectId) (interface{}, error) {
	return nil, nil
}

func (s *Storage) Update(data interface{}, query ...models.Obj) error {
	return nil
}

func (s *Storage) Delete(key bson.ObjectId) {

}
