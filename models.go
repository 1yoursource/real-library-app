package main

import (
	"time"

	"gorm.io/gorm"
)

type (
	obj map[string]interface{}

	Book struct {
		Id          uint64    `json:"id" bson:"_id"`
		Name        string    `json:"name" bson:"name"`
		Author      string    `json:"author" bson:"author"`
		PublishYear uint64    `json:"publishYear" bson:"publishYear"`
		Publisher   string    `json:"publisher" bson:"publisher"`
		PagesCount  uint64    `json:"pagesCount" bson:"pagesCount"`
		ReturnDate  time.Time `json:"returnDate" bson:"returnDate"`
	}

	//Test struct {
	//	Id   uuid.UUID `sql:"id"`
	//	Name string    `sql:"name"`
	//}

	Test struct {
		gorm.Model
		//Id   int
		Name string
		//DataId in data_id"`
		Data Data
	}

	Data struct {
		gorm.Model
		Test1ID int
		Info    string
	}
)

const (
	standartTimeFmt = "2006-01-02 15:04:05"

	zero = 0
	six  = 6

	coockieMaxAge = 32000000
)
