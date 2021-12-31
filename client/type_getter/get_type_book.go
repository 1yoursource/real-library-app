package type_getter

import (
	"lib-client-server/client/models"
)

func GetTypeBook(data interface{}) (*models.Book, bool) {
	switch v := data.(type) {
	case models.Book:
		return &v, true
	default:
		return nil, false
	}
}
