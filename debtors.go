package main

import (
	"net/http"
	"time"
	"fmt"
)

// запускается каждый день после полуночи
func FindExpiredBooksByUser(ticketNumber string) []Book {
	books := []Book{}

	if err := storage.C("books").Find(obj{"takenBy": ticketNumber, "returnDate": obj{"$lte": time.Now().Format(standartDateFmt)}}).All(&books); err != nil {
		fmt.Println("debtors.go -> FindExpiredBooksByUser -> books not found, err:", err)
		return []Book{}
	}

	return books
}


func MakeDebtorsList()  {
	debtors := []Debtor{}

	users := []User{}
	if err := storage.C("users").Find(nil).All(&users); err != nil {
		fmt.Println("debtors.go -> MakeDebtorsList -> users not found, err:", err)
		return
	}

	for _, user := range users {
		books := FindExpiredBooksByUser(user.TicketNumber)
		if len(books)>0 {
			debtor := Debtor{
				User:  user.TicketNumber,
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
