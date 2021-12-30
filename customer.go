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
	inputData := struct {
		UserId string `json:"userId"`
		BookId string `json:"bookId"`
	}{}

	if err := c.Bind(&inputData); err != nil {
		fmt.Println("customer.go -> GetBook -> Bind: err = ", err)
		c.JSON(http.StatusInternalServerError, obj{"error": "wrong"})
		return
	}
	user := User{}
	book := Book{}
	err := storage.C("users").FindId(bson.ObjectId(inputData.UserId)).One(&user)
	if err != nil {
		fmt.Println("customer.go -> GetBook -> user not found, err:", err)
		c.JSON(http.StatusNotFound, obj{"error": "user not found"})
		return
	}

	err = storage.C("books").FindId(bson.ObjectId(inputData.BookId)).One(&book)
	if err != nil {
		fmt.Println("customer.go -> GetBook -> book not found, err:", err)
		c.JSON(http.StatusNotFound, obj{"error": "book not found"})
		return
	}
	// проверяем на занятость
	if book.TakenBy != "" {
		c.JSON(http.StatusBadRequest, obj{"error": "Book already taken"})
		return
	}
	// добавляем в список
	user.Books = append(user.Books, book)
	//сохраняем в базу юзеров изменения.
	storage.C("users").Update(user.Id, user)

	book.TakenBy = user.TicketNumber
	book.ReturnDate = time.Now().AddDate(0, 0, 30) // 30 дней с момента взятия книги
	// сохраняем в базу книг
	storage.C("books").Update(book.Id, book)

	c.JSON(http.StatusOK, obj{"answer": "ok"})

}

func (u *UserModule) ReturnBook(c *gin.Context) {
	inputData := struct {
		UserId string `json:"userId"`
		BookId string `json:"bookId"`
	}{}
	if err := c.Bind(&inputData); err != nil {
		fmt.Println("customer.go -> GetBook -> Bind: err = ", err)
		c.JSON(http.StatusInternalServerError, obj{"error": "wrong"})
		return
	}

	user := User{}
	book := Book{}
	err := storage.C("users").FindId(bson.ObjectId(inputData.UserId)).One(&user)
	if err != nil {
		fmt.Println("customer.go -> GetBook -> user not found, err:", err)
		c.JSON(http.StatusNotFound, obj{"error": "user not found"})
		return
	}
	err = storage.C("books").FindId(bson.ObjectId(inputData.BookId)).One(&book)
	if err != nil {
		fmt.Println("customer.go -> GetBook -> book not found, err:", err)
		c.JSON(http.StatusNotFound, obj{"error": "book not found"})
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
	storage.C("users").Update(user.Id, user)

	book.TakenBy = ""
	//сохраняем в базу книг
	storage.C("books").Update(book.Id, book)

	c.JSON(http.StatusOK, obj{"answer": "ok"})
}

// получить список книг у читателя
func (u *UserModule) GetAllTakenBooks(c *gin.Context) {
	inputData := struct {
		UserId string `json:"userId"`
	}{}
	if err := c.Bind(&inputData); err != nil {
		fmt.Println("customer.go -> GetAllTakenBooks -> Bind: err = ", err)
		c.JSON(http.StatusInternalServerError, obj{"error": "wrong"})
		return
	}
	user := User{}
	err := storage.C("users").FindId(bson.ObjectId(inputData.UserId)).One(&user)
	if err != nil {
		fmt.Println("customer.go -> GetBook -> user not found, err:", err)
		c.JSON(http.StatusNotFound, obj{"error": "user not found"})
		return
	}
	booksList := []string{}
	for _, book := range user.Books {
		// записываем список книг с авторами, которые на руках у данного пользователя
		booksList = append(booksList, book.Name+", "+book.Author)
	}

	c.JSON(http.StatusOK, obj{"books": booksList})
}

func (u *UserModule) CreateUser(data Registration, id bson.ObjectId) error {
	if err := data.CheckEmail(); err != nil {
		return err
	}

	password, err := createHashPassword(data.Password)
	if err != nil {
		return err
	}

	user := User{
		Id:           id,
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
