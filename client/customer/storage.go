package customer

import (
	"sync"

	"lib-client-server/client"
)

type Storage struct {
	KeyType      string
	FreeBooksMut sync.RWMutex
	FreeBooks    map[string]client.Book
}

func createStorage(keyType string) client.StorageInterface {
	return &Storage{KeyType: keyType, FreeBooks: map[string]client.Book{}}
}

func (s *Storage) GetAll() []client.Book {
	return []client.Book{}
}

func (s *Storage) Set(data client.Book) {
}

func (s *Storage) Get(key interface{}) (client.Book, bool) {
	return client.Book{}, true
}

func (s *Storage) Update(key interface{}) error {
	return nil
}

func (s *Storage) Delete(key interface{}) {

}
