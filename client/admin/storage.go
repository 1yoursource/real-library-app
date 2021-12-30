package admin

import (
	"sync"

	"gopkg.in/mgo.v2/bson"

	"lib-client-server/client"
)

type Storage struct {
	KeyType      string
	FreeBooksMut sync.RWMutex
	FreeBooks    map[bson.ObjectId]client.Book
}

func createStorage(keyType string) client.StorageInterface {
	return &Storage{KeyType: keyType, FreeBooks: map[bson.ObjectId]client.Book{}}
}

func (s *Storage) Set(data client.Book) {
	s.FreeBooksMut.Lock()
	s.FreeBooks[data.Id] = data
	s.FreeBooksMut.Unlock()
}

func (s *Storage) GetAll() []client.Book {
	var result []client.Book
	s.FreeBooksMut.RLock()
	for _, v := range s.FreeBooks {
		result = append(result, v)
	}
	s.FreeBooksMut.RUnlock()
	return result
}

func (s *Storage) Get(key bson.ObjectId) (client.Book, bool) {
	return client.Book{}, true
}

func (s *Storage) Update(book client.Book, query ...client.Obj) error {
	return nil
}

func (s *Storage) Delete(key bson.ObjectId) {

}
