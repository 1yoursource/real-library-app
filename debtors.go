package main

import (
	"fmt"
	"lib-client-server/client/models"
	"time"
)

// запускается каждый день после полуночи
func FindExpiredBooksByUser(ticketNumber uint64) []models.Book {
	books := []models.Book{}

	if err := storage.C("books").Find(obj{"takenBy": ticketNumber, "returnDate": obj{"$lte": time.Now().Format(standartDateFmt)}}).All(&books); err != nil {
		fmt.Println("debtors.go -> FindExpiredBooksByUser -> books not found, err:", err)
		return []models.Book{}
	}

	return books
}

func MakeDebtorsList() {
	debtors := []Debtor{}

	users := []models.User{}
	if err := storage.C("users").Find(nil).All(&users); err != nil {
		fmt.Println("debtors.go -> MakeDebtorsList -> users not found, err:", err)
		return
	}

	for _, user := range users {
		books := FindExpiredBooksByUser(user.Id)
		if len(books) > 0 {
			debtor := Debtor{
				User:  user.Id,
				Books: books,
			}
			debtors = append(debtors, debtor)
		}
	}

	debtorsList := DebtorsList{
		Date:    time.Now().Format(standartDateFmt),
		Debtors: debtors,
	}

	if err := storage.C("debtors").Insert(debtorsList); err != nil {
		fmt.Println("debtors.go -> MakeDebtorsList -> Insert err:", err)
	}
}
