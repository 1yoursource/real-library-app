package admin

import (
	main_opt "lib-client-server/client/main"
	"sync"

	"gopkg.in/mgo.v2/bson"

	"lib-client-server/client"
)

type Storage struct {
	//KeyType      string
	//FreeBooksMut sync.RWMutex
	//FreeBooks    map[bson.ObjectId]client.Book
}

func createStorage() models.StorageInterface {
	return &Storage{}
}

func (s *Storage) Set(data interface{}) {
	//s.FreeBooksMut.Lock()
	//s.FreeBooks[data.Id] = data
	//s.FreeBooksMut.Unlock()
}

func (s *Storage) GetAll() []interface{} {
	//var result []client.Book
	//s.FreeBooksMut.RLock()
	//for _, v := range s.FreeBooks {
	//	result = append(result, v)
	//}
	//s.FreeBooksMut.RUnlock()
	return nil
}

func (s *Storage) Get(key bson.ObjectId) (interface{}, bool) {
	return nil, true
}

func (s *Storage) Update(book client.Book, query ...models.Obj) error {
	return nil
}

func (s *Storage) Delete(key bson.ObjectId) {

}
