package type_getter

import "lib-client-server/client"

func GetTypeBook(data interface{}) (*client.Book,bool) {
	switch v := data.(type) {
	case client.Book:
		return &v, true
	default:
		return nil, false
	}
}
