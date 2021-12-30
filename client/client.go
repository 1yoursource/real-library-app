package client

import (
	"time"
)

type (

	Book struct {
		Id          string `form:"id" bson:"_id"`
		Name        string        `form:"name" bson:"name"`
		Author      string        `form:"author" bson:"author"`
		PublishYear uint64        `form:"publishYear" bson:"publishYear"`
		Publisher   string        `form:"publisher" bson:"publisher"`
		PagesCount  uint64        `form:"pagesCount" bson:"pagesCount"`
		ReturnDate  time.Time     `form:"returnDate" bson:"returnDate"`
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

)