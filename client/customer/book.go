package customer

import (
	"fmt"
	"lib-client-server/client/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	BookMod struct {
		Storage
	}

	GetBook struct {
		Id string `form:"id"`
	}
)

func CreateBookModule() models.BookInterface {
	return &BookMod{}
}

func (b *BookMod) GetAll(c *gin.Context) {
	//var books []client.Book
	//b.FreeBooksMut.RLock()
	//for _, v := range b.FreeBooks {
	//	books = append(books, v)
	//}
	//b.FreeBooksMut.RUnlock()

	//c.JSON(http.StatusOK, client.Obj{"result": books})
}

func (b *BookMod) Create(c *gin.Context) { // взять
	inputData := GetBook{}
	if err := c.Bind(&inputData); err != nil {
		fmt.Println("auth.go -> Registration -> Bind: err = ", err)
		c.JSON(http.StatusInternalServerError, models.Obj{"error": "wrong"})
		return
	}

	c.JSON(http.StatusOK, models.Obj{"result": "books"})
}

func (b *BookMod) Read(c *gin.Context) {

} // a:see,u:see info //todo NEED ?????

func (b *BookMod) Update(c *gin.Context) {

} // a:edit,u:

func (b *BookMod) Delete(c *gin.Context) {

} // a:del,u:return
