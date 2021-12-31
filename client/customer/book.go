package customer

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/night-codes/types.v1"
	"lib-client-server/client/models"
	"lib-client-server/database"
	"net/http"
	"time"
)

type Book struct {
	Storage
}

var bookModule *Book

func CreateUserBookModule(host, dbName string) models.BookInterface {
	bookModule := &Book{Storage{collection: "books", Database: database.Connect(host, dbName, "libraryDatabase")}}
	return bookModule
}

func (b *Book) Handler(c *gin.Context) {
	switch c.Param("method") {
	case "search":
		b.GetAll(c)
	case "take":
		b.Create(c)
	case "return":
		b.Delete(c)
	default:
		c.String(http.StatusBadRequest, "Module not found!")
	}
}

func (b *Book) GetAll(c *gin.Context) {
	var input = struct {
		FilterId string `form:"filter"`
		Value    string `form:"value"`
	}{}

	if err := c.Bind(&input); err != nil {
		fmt.Println("usr book GetAll -> Bind: err = ", err)
		c.JSON(http.StatusInternalServerError, models.Obj{"error": "wrong"})
		return
	}

	var (
		books []models.Book
		err   error
	)

	switch input.FilterId {
	case "1":
		err = b.Storage.GetAll().All(&books)
	case "2": // за назвою
		err = b.Storage.GetByQuery(models.Obj{"name": input.Value}).All(&books)
	case "3": // за автором
		err = b.Storage.GetByQuery(models.Obj{"author": input.Value}).All(&books)
	case "4": // усі взяті користувачем книги
		fmt.Println("Aaaaaaaaaaaaaaaaaaaaaaaaa",input)
		books, err  = b.getAllUserBooks(types.Uint64(input.Value))
	case "5": // взяті користувачем книги, які підлягають поверненню
		books, err  = b.getUserBooksForReturning(types.Uint64(input.Value))
	default:
		c.JSON(http.StatusBadRequest, models.Obj{"error": "unknown filter", "result": []models.Book{}})
		return
	}

	c.JSON(http.StatusOK, models.Obj{"error": err, "result": books})
}

func (b *Book) Create(c *gin.Context) { // користувач бере книгу собі
	var data = struct {
		BookId uint64 `form:"bookId"`
		UserId uint64 `form:"userId"`
	}{}

	if err := c.Bind(&data); err != nil {
		fmt.Println("usr book Create -> Bind: err = ", err)
		c.JSON(http.StatusInternalServerError, models.Obj{"error": "wrong"})
		return
	}

	// перевіряємо чи книжка взята
	if err := b.checkAvailable(data.BookId); err != nil {
		fmt.Println("usr book Create -> checkAvailable: err = ", err)
		c.JSON(http.StatusInternalServerError, models.Obj{"error": err})
		return
	}

	var (
		selector = models.Obj{"_id": data.BookId}
		updater  = models.Obj{"$set": models.Obj{"takenBy": fmt.Sprint(data.UserId), "returnDate": time.Now().AddDate(0, 1, 0)}}
	)

	if err := b.Storage.Update(selector, updater); err != nil {
		fmt.Println("usr book Create -> Update: err = ", err)
		c.JSON(http.StatusInternalServerError, models.Obj{"error": err})
		return
	}

	var book models.Book
	if err := b.Storage.GetByQuery(selector).One(&book); err != nil {
		fmt.Println("usr book Create -> Update: err = ", err)
		c.JSON(http.StatusInternalServerError, models.Obj{"error": err})
		return
	}

	selector = models.Obj{"_id": data.UserId}
	updater = models.Obj{"$push": models.Obj{"books": book}}

	if err := b.Storage.C("users").Update(selector, updater); err != nil {
		fmt.Println("usr book Create -> users.Update: err = ", err)
		c.JSON(http.StatusInternalServerError, models.Obj{"error": err})
		return
	}

	c.JSON(http.StatusOK, models.Obj{"error": nil, "result": data.BookId})
} // a:cr,u:get

func (b *Book) Read(c *gin.Context) {
	var book = models.Book{}
	if err := c.Bind(&book); err != nil {
		fmt.Println("usr book Read -> Bind: err = ", err)
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
	var book = models.Book{}
	if err := c.Bind(&book); err != nil {
		fmt.Println("usr book Update -> Bind: err = ", err)
		c.JSON(http.StatusInternalServerError, models.Obj{"error": "wrong"})
		return
	}

	// todo ???
	// add checking for new data in struct => by marshaling and compare 2 string

} // a:edit,u:

func (b *Book) Delete(c *gin.Context) {
	var data = struct {
		BookId uint64 `form:"bookId"`
		UserId uint64 `form:"userId"`
	}{}

	if err := c.Bind(&data); err != nil {
		fmt.Println("usr book Update -> Bind: err = ", err)
		c.JSON(http.StatusInternalServerError, models.Obj{"error": "wrong"})
		return
	}

	var (
		selector = models.Obj{"_id": data.BookId}
		updater  = models.Obj{"$set": models.Obj{"takenBy": "", "returnDate": time.Time{}}}
	)

	if err := b.Storage.Update(selector, updater); err != nil {
		fmt.Println("usr book Create -> Update: err = ", err)
		c.JSON(http.StatusInternalServerError, models.Obj{"error": err})
		return
	}

	selector = models.Obj{"_id": data.UserId}
	updater = models.Obj{"$pull": models.Obj{"books": fmt.Sprint(data.BookId)}}

	if err := b.Storage.C("users").Update(selector, updater); err != nil {
		fmt.Println("usr book Create -> users.Update: err = ", err)
		c.JSON(http.StatusInternalServerError, models.Obj{"error": err})
		return
	}

	c.JSON(http.StatusOK, models.Obj{"error": nil, "result": data.BookId})
} // a:del,u:return

func (b *Book) checkAvailable(id uint64) error {
	var book models.Book
	if err := b.Storage.GetByQuery(models.Obj{"_id": id}).One(&book); err != nil {
		return err
	}

	if book.Id == 0 {
		return errors.New("does not exist")
	}

	if len(book.TakenBy) > 0 {
		return errors.New("already in use")
	}

	return nil
}

func (b *Book) getAllUserBooks(userId uint64) ([]models.Book, error) {
	var user = models.User{}
	if err := b.Storage.C("users").FindId(userId).One(&user); err != nil {
		return []models.Book{}, err
	}

	fmt.Println()

	return user.Books, nil
}

func (b *Book) getUserBooksForReturning(userId uint64) ([]models.Book, error) {
	var user = models.User{}
	if err := b.Storage.C("users").FindId(userId).One(&user); err != nil {
		return []models.Book{}, err
	}

	var result []models.Book
	for _,v := range user.Books {
		if v.ReturnDate.After(time.Now()) {
			continue
		}
		result = append(result,v)
	}

	return result, nil
}