package admin

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"lib-client-server/auto_incrementer"
	"lib-client-server/client"
	"lib-client-server/client/models"
	"lib-client-server/database"
	"net/http"
	"strings"
)

type Book struct {
	Storage
}

var bookModule *Book

func CreateAdminBookModule(host, dbName string) models.BookInterface {
	bookModule := &Book{Storage{collection:"books",Database:database.Connect(host,dbName, "libraryDatabase")}}
	return bookModule
}

func (b *Book) Handler(c *gin.Context) {
	switch c.Param("method") {
	case "add":
		b.Create(c)
	case "search":
		b.GetAll(c)
	case "delete":
		b.Delete(c)
	default:
		c.String(http.StatusBadRequest, "Module not found!")
	}
}

func (b *Book) GetAll(c *gin.Context) {
	// has rights?
	var input = struct {
		FilterId string `form:"filter"`
		Value string `form:"value"`
	}{}

	if err := c.Bind(&input); err != nil {
		fmt.Println("admin book GetAll -> Bind: err = ", err)
		c.JSON(http.StatusInternalServerError, models.Obj{"error": "wrong"})
		return
	}

	var (
		books []client.Book
		err error
	)

	switch input.FilterId {
	case "1":
		err = b.Storage.GetAll().All(&books)
	case "2":
		err = b.Storage.GetByQuery(models.Obj{"author":input.Value}).All(&books)
	case "3": // все книги пользователя
		err = b.Storage.GetByQuery(models.Obj{"takenBy":input.Value}).All(&books)
	case "4": // должники
	default:
		c.JSON(http.StatusBadRequest,models.Obj{"error":"unknown filter","result":[]client.Book{}})
		return
	}

	c.JSON(http.StatusOK, models.Obj{"error": err, "result": books})
}

func (b *Book) Create(c *gin.Context) { // add book
	var book = client.Book{}
	if err := c.Bind(&book); err != nil {
		fmt.Println("admin book Create -> Bind: err = ", err)
		c.JSON(http.StatusInternalServerError, models.Obj{"error": "wrong"})
		return
	}

	if err := b.checkUniq(book); err != nil {
		fmt.Println("admin book Create -> Registration -> checkUniq: err = ", err)
		c.JSON(http.StatusInternalServerError, models.Obj{"error": "book already exist"})
		return
	}

	book.Id = auto_incrementer.AI.Next(b.Storage.collection)

	b.Storage.Set(book)

	c.JSON(http.StatusOK, models.Obj{"error": nil, "result": book})
} // a:cr,u:get

func (b *Book) Read(c *gin.Context) {
	var book = client.Book{}
	if err := c.Bind(&book); err != nil {
		fmt.Println("admin book Read -> Bind: err = ", err)
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

func (b *Book) Update(c *gin.Context) { // todo unused
	var book = client.Book{}
	if err := c.Bind(&book); err != nil {
		fmt.Println("admin book Update -> Bind: err = ", err)
		c.JSON(http.StatusInternalServerError, models.Obj{"error": "wrong"})
		return
	}

	// todo ???
	// add checking for new data in struct => by marshaling and compare 2 string

} // a:edit,u:

func (b *Book) Delete(c *gin.Context) {
	var data = struct{
		BookId uint64 `form:"bookId"`
	}{}

	if err := c.Bind(&data); err != nil {
		fmt.Println("admin book Update -> Bind: err = ", err)
		c.JSON(http.StatusInternalServerError, models.Obj{"error": "wrong"})
		return
	}

	err := b.Storage.Delete(data.BookId)

	c.JSON(http.StatusOK,models.Obj{"error":err,"result":data.BookId})
} // a:del,u:return

func (b *Book) checkUniq(newBook client.Book) error {
	var book client.Book
	err := b.Storage.GetByQuery(models.Obj{"author":newBook.Author,"name":newBook.Name,"publishYear":newBook.PublishYear}).One(&book)
	switch {
	case err == nil:
		break
	case strings.Contains(err.Error(),"not found"):
		return nil
	default:
		return err
	}

	switch {
	case book.Id == 0:
		return errors.New("does not exist")
	case book.Author != newBook.Author:
		return nil
	case book.Name != newBook.Name:
		return nil
	case book.Publisher != newBook.Publisher:
		return nil
	case book.PublishYear != newBook.PublishYear:
		return nil
	default:
		return errors.New("already not exist")
	}
}