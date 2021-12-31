package main

import (
	"lib-client-server/client/models"
)

type (
	obj map[string]interface{}

	Debtor struct {
		User  uint64
		Books []models.Book
	}
	DebtorsList struct {
		Date    string   `json:"date" bson:"_id"`
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
