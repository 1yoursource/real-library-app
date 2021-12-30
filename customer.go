package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"

	"gopkg.in/mgo.v2/bson"
)

type (
	UserModule struct{}

	User struct {
		Id           bson.ObjectId `bson:"_id"`
		TicketNumber string        `bson:"ticketNumber"`
		Email        string        `bson:"email"`
		Password     []byte        `bson:"password"`
		FirstName    string        `bson:"firstName"`
		LastName     string        `bson:"lastName"`
		SurName      string        `bson:"surName"`
		Faculty      string        `bson:"faculty"`
		LastLogin    string        `bson:"lastLogin"`
		Books        []Book        `bson:"books"`
	}

	AuthData struct {
		Login    string `form:"login" bson:"login"`
		Password string `form:"password" bson:"password"`
	}
)

func (u UserModule) Create() *UserModule {
	return &u
}

func (u *UserModule) GetBook(c *gin.Context) {
	user := User{} // получать юзера
	book := Book{}
	if err := c.Bind(&book); err != nil {
		fmt.Println("customer.go -> GetBook -> Bind: err = ", err)
		c.JSON(http.StatusInternalServerError, obj{"error": "wrong"})
		return
	}
	// смотрим книгу в базе данных, проверяем на занятость
	if book.TakenBy != "" {
		c.JSON(http.StatusBadRequest, obj{"error": "Book already taken"})
		return
	}
	// добавляем в список
	user.Books = append(user.Books, book)
	//сохраняем в базу юзеров изменения.

	book.TakenBy = user.TicketNumber
	// сохраняем в базу книг

	c.JSON(http.StatusOK, obj{"answer": "ok"})

}

func (u *UserModule) ReturnBook(c *gin.Context) {
	user := User{} // получать юзера
	book := Book{}
	if err := c.Bind(&book); err != nil {
		fmt.Println("customer.go -> ReturnBook -> Bind: err = ", err)
		c.JSON(http.StatusInternalServerError, obj{"error": "wrong"})
		return
	}
	userBooksNew := []Book{}
	for _, v := range user.Books {
		if book.Id != v.Id {
			userBooksNew = append(userBooksNew, v)
		}
	}
	user.Books = userBooksNew
	// сохраняем в базу юзеров

	book.TakenBy = ""
	//сохраняем в базу книг

	c.JSON(http.StatusOK, obj{"answer": "ok"})
}

func (u *UserModule) GetAllTakenBooks() {

}

func (u *UserModule) CreateUser(data Registration) error {
	if err := data.CheckEmail(); err != nil {
		return err
	}

	password, err := createHashPassword(data.Password)
	if err != nil {
		return err
	}

	user := User{
		Id:           bson.NewObjectId(),
		TicketNumber: fmt.Sprint(len(data.LastName), data.Faculty, "-", time.Now().UnixNano()),
		FirstName:    data.FirstName,
		LastName:     data.LastName,
		SurName:      data.SurName,
		Faculty:      data.Faculty,
		Email:        data.Email,
		Password:     password,
	}

	if err := storage.C("users").Insert(user); err != nil {
		return err
	}

	return nil
}

func (r *Registration) CheckEmail() error {
	fmt.Println("email: ", r.Email)
	fmt.Println("storage: ", storage)
	fmt.Println("storage.C(users)", storage.C("users"))
	fmt.Println("email: ", r.Email)
	count, err := storage.C("users").Find(obj{"email": r.Email}).Count()
	fmt.Println("count: ", count)
	fmt.Println("err: ", err)

	switch {
	case err != nil:
		break
	case count > 0:
		fmt.Println("HETE: ", count, "--", err)
		err = fmt.Errorf("already registered")
	}

	return err
}

//func (r *Registration) checkFaculty() {
//	r.FirstName
//}

// todo
func (u *UserModule) BlockUser() {

}
