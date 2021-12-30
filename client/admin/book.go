package admin

import (
	"fmt"
	"lib-client-server/client"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/gin-gonic/gin"
)

type Book struct {
	Storage
}

func CreateBookModule() client.BookInterface {
	return &Book{}
}

func (b *Book) GetAll(c *gin.Context) {
	// has rights?
	c.JSON(http.StatusOK, client.Obj{"result": b.Storage.GetAll()})
}

func (b *Book) Create(c *gin.Context) {
	var book = client.Book{}
	if err := c.Bind(&book); err != nil {
		fmt.Println("auth.go -> Registration -> Bind: err = ", err)
		c.JSON(http.StatusInternalServerError, client.Obj{"error": "wrong"})
		return
	}

	book.Id = bson.NewObjectId()

	b.Storage.Set(book)

	c.JSON(http.StatusOK, client.Obj{"result": book})
} // a:cr,u:get

func (b *Book) Read(c *gin.Context) {
	var book = client.Book{}
	if err := c.Bind(&book); err != nil {
		fmt.Println("auth.go -> Registration -> Bind: err = ", err)
		c.JSON(http.StatusInternalServerError, client.Obj{"error": "wrong"})
		return
	}

	book, exist := b.Storage.Get(book.Id)
	if !exist {
		c.JSON(http.StatusInternalServerError, client.Obj{"error": "not found"})
		return
	}
	c.JSON(http.StatusOK, client.Obj{"result": book})

} // a:see,u:see info //todo NEED ?????

func (b *Book) Update(c *gin.Context) {
	var book = client.Book{}
	if err := c.Bind(&book); err != nil {
		fmt.Println("auth.go -> Registration -> Bind: err = ", err)
		c.JSON(http.StatusInternalServerError, client.Obj{"error": "wrong"})
		return
	}

	b.Storage.Update(book)

} // a:edit,u:

func (b *Book) Delete(c *gin.Context) {

} // a:del,u:return
