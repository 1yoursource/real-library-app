package main

import (
	"time"
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
		TakenBy     uint64    `json:"takenBy" bson:"takenBy"` // кем взято, если 0 - книга доступна для взятия
	}
	Debtor struct {
		User  uint64
		Books []Book
	}
	DebtorsList struct {
		Date string `json:"date" bson:"_id"`
		Debtors []Debtor `bson:"debtors"`
	}
)

const (
	standartTimeFmt = "2006-01-02 15:04:05"
	standartDateFmt = "2006-01-02"

	zero = 0
	six  = 6

	coockieMaxAge = 32000000
)
