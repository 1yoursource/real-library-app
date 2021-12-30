package main

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	AuthModule struct {
	}

	Registration struct {
		Email          string `form:"email"`
		Password       string `form:"password"`
		PasswordSubmit string `form:"passwordSubmit"`
		FirstName      string `form:"firstName"`
		LastName       string `form:"lastName"`
		SurName        string `form:"surName"`
		Faculty        string `form:"faculty"`
	}
)

func (a AuthModule) Create() *AuthModule {
	return &a
}

func (a *AuthModule) Ajax(c *gin.Context) {
	switch c.Param("method") {
	case "registration":
		a.Registration(c)
	case "authorization":
		a.Authorization(c)
	default:
		c.String(http.StatusBadRequest, "Method not found in module \"PRO<O\"!")
	}
}

func (a *AuthModule) Registration(c *gin.Context) {
	inputData := Registration{}
	if err := c.Bind(&inputData); err != nil {
		fmt.Println("auth.go -> Registration -> Bind: err = ", err)
		c.JSON(http.StatusInternalServerError, obj{"error": "wrong"})
		return
	}

	if !inputData.passwordCompare() {
		c.JSON(http.StatusInternalServerError, obj{"error": "please, use equals passwords"})
		return
	}

	fmt.Printf("inputData = %+v\n", inputData)

	var id = bson.NewObjectId()

	if err := user.CreateUser(inputData,id); err != nil {
		fmt.Println("auth.go -> Registration -> CreateUser: err = ", err)
		c.JSON(http.StatusInternalServerError, obj{"error": err.Error()})
		return
	}

	a.Login(c, inputData.Email,id)
}

func (a *AuthModule) Authorization(c *gin.Context) {
	inputData := AuthData{}
	if err := c.Bind(&inputData); err != nil {
		fmt.Println("auth.go -> Authorization -> Bind: err = ", err)
		c.JSON(http.StatusInternalServerError, obj{"error": "wrong"})
		return
	}

	var user User
	if err := storage.C("users").Find(obj{"email": inputData.Login}).One(&user); err != nil {
		fmt.Println("auth.go -> Authorization -> Find: err = ", err)
		c.JSON(http.StatusInternalServerError, obj{"error": "not found"})
		return
		//return err
	}

	if err := inputData.checkPassword(user.Password); err != nil {
		fmt.Println("auth.go -> Authorization -> checkPassword: err = ", err)
		c.JSON(http.StatusInternalServerError, obj{"error": err.Error()})
		return
	}

	a.Login(c, inputData.Login, user.Id)
}

func (ad *AuthData) checkPassword(userPassword []byte) error {
	return hashedPasswordCompare(ad.Password, userPassword)
}

func (r *Registration) passwordCompare() bool {
	return r.Password == r.PasswordSubmit
}

func (a *AuthModule) Login(c *gin.Context, email string, id bson.ObjectId) {
	fmt.Println("login")
	setCookie(c, "lib-login", fmt.Sprint(email, "*lib"))
	setCookie(c, "lib-id", fmt.Sprint(id, "*lib"))
	c.JSON(http.StatusOK, obj{"error": nil, "url": "/"})
}

func (a *AuthModule) Logout(c *gin.Context) {
	deleteCookie(c, "lib-login")
	deleteCookie(c, "lib-id")
	pages.Auth(c)
}
