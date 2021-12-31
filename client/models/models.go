package models

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"time"
)

type (
	StorageInterface interface {
		Set(data interface{})
		Get(key uint64) (interface{}, error)
		GetByQuery(query interface{}) *mgo.Query
		GetAll() *mgo.Query
		Update(selector Obj, updater Obj) error
		Delete(key uint64) error
	}
	BookInterface interface {
		Handler(c *gin.Context)
		GetAll(c *gin.Context)
		Create(c *gin.Context) // a:cr,u:get
		Read(c *gin.Context)   // a:see,u:see info //todo NEED ?????
		Update(c *gin.Context) // a:edit,u:
		Delete(c *gin.Context) // a:del,u:return
	}

	Obj map[string]interface{}

	User struct {
		Id           uint64 `bson:"_id"`
		TicketNumber string `bson:"ticketNumber"`
		Email        string `bson:"email"`
		Password     []byte `bson:"password"`
		FirstName    string `bson:"firstName"`
		LastName     string `bson:"lastName"`
		SurName      string `bson:"surName"`
		Faculty      string `bson:"faculty"`
		LastLogin    string `bson:"lastLogin"`
		Books        []Book `bson:"books"`
	}

	Book struct {
		Id          uint64    `form:"id" bson:"_id"`
		Name        string    `form:"name" bson:"name"`
		Author      string    `form:"author" bson:"author"`
		PublishYear uint64    `form:"publishYear" bson:"publishYear"`
		Publisher   string    `form:"publisher" bson:"publisher"`
		PagesCount  uint64    `form:"pagesCount" bson:"pagesCount"`
		ReturnDate  time.Time `form:"returnDate" bson:"returnDate"`
		TakenBy     string    `json:"takenBy" bson:"takenBy"` // кем взято, если "" - книга доступна для взятия
	}

	Dept struct {
		TicketNumber string `form:"ticketNumber"`
		Book         string `form:"book"`
		ReturnDate   string `form:"returnDate"`
	}
)

const (
	CoockieMaxAge = 32000000
)
