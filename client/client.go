package client

import (
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/gin-gonic/gin"
)

type (
	Obj map[string]interface{}

	Book struct {
		Id          bson.ObjectId `form:"id" bson:"_id"`
		Name        string        `form:"name" bson:"name"`
		Author      string        `form:"author" bson:"author"`
		PublishYear uint64        `form:"publishYear" bson:"publishYear"`
		Publisher   string        `form:"publisher" bson:"publisher"`
		PagesCount  uint64        `form:"pagesCount" bson:"pagesCount"`
		ReturnDate  time.Time     `form:"returnDate" bson:"returnDate"`
	}

	BookInterface interface {
		GetAll(c *gin.Context)
		Create(c *gin.Context) // a:cr,u:get
		Read(c *gin.Context)   // a:see,u:see info //todo NEED ?????
		Update(c *gin.Context) // a:edit,u:
		Delete(c *gin.Context) // a:del,u:return
	}

	AuthInterface interface {
		Login()
		Logout()
		Register()
	}

	PageInterface interface {
		Main()
		Auth()
		Search()
		Shell()
		About()
	}

	DBInterface interface {
		Insert()
		Update()
		Find()
	}

	UserInterface interface {
		Create()
		Block()
		Info()
		Update() // недоступно для админа
	}

	StorageInterface interface {
		Set(data Book)
		Get(key bson.ObjectId) (Book, bool)
		GetAll() []Book
		Update(book Book, query ...Obj)
		Delete(key bson.ObjectId)
	}
)
