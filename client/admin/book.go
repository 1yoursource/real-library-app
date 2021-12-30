package admin

import (
	"errors"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"lib-client-server/client"
	"lib-client-server/client/models"
	"lib-client-server/client/type_getter"
	"lib-client-server/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	Storage
}

func CreateBookModule() client.BookInterface {
	return &Book{Storage{collection:"adminBookDBc",Database:database.Connect("localhost", "libDB", "libraryDatabase")}}
}

func (b *Book) GetAll(c *gin.Context) {
	// has rights?
	c.JSON(http.StatusOK, models.Obj{"result": b.Storage.GetAll()})
}

func (b *Book) Create(c *gin.Context) {
	var book = client.Book{}
	if err := c.Bind(&book); err != nil {
		fmt.Println("auth.go -> Registration -> Bind: err = ", err)
		c.JSON(http.StatusInternalServerError, models.Obj{"error": "wrong"})
		return
	}

	if err := b.checkUniq(book); err != nil {
		fmt.Println("auth.go -> Registration -> Bind: err = ", err)
		c.JSON(http.StatusInternalServerError, models.Obj{"error": "book already exist"})
		return
	}

	book.Id = bson.NewObjectId()

	b.Storage.Set(book)

	c.JSON(http.StatusOK, models.Obj{"result": book})
} // a:cr,u:get

func (b *Book) Read(c *gin.Context) {
	var book = client.Book{}
	if err := c.Bind(&book); err != nil {
		fmt.Println("auth.go -> Registration -> Bind: err = ", err)
		c.JSON(http.StatusInternalServerError, models.Obj{"error": "wrong"})
		return
	}

	//book, exist := b.Storage.Get(book.Id)
	//if !exist {
	//	c.JSON(http.StatusInternalServerError, models.Obj{"error": "not found"})
	//	return
	//}
	c.JSON(http.StatusOK, models.Obj{"result": book})

} // a:see,u:see info //todo NEED ?????

func (b *Book) Update(c *gin.Context) {
	var book = client.Book{}
	if err := c.Bind(&book); err != nil {
		fmt.Println("auth.go -> Registration -> Bind: err = ", err)
		c.JSON(http.StatusInternalServerError, models.Obj{"error": "wrong"})
		return
	}

	// todo ???
	// add checking for new data in struct => by marshaling and compare 2 string

	b.Storage.Update(book)

} // a:edit,u:

func (b *Book) Delete(c *gin.Context) {

} // a:del,u:return

func (b *Book) checkUniq(newBook client.Book) error {
	record, err := b.GetByQuery(models.Obj{"author":newBook.Author,"name":newBook.Name,"publishYear":newBook.PublishYear})
	if err != nil {
		return err
	}

	book,isIt:=type_getter.GetTypeBook(record)

	switch {
	case !isIt:
		return errors.New("smth error")
	case book == nil:
		return errors.New("does not exist")
	case book.Author != newBook.Author:
		return nil
	case book.Name != newBook.Name:
		return nil
	case book.PublishYear != newBook.PublishYear:
		return nil
	default:
		return errors.New("already not exist")
	}
}