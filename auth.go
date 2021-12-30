package main

import (
	"fmt"
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

	if err := user.CreateUser(inputData); err != nil {
		fmt.Println("auth.go -> Registration -> CreateUser: err = ", err)
		c.JSON(http.StatusInternalServerError, obj{"error": err.Error()})
		return
	}

	a.Login(c, inputData.Email)
}

func (a *AuthModule) Authorization(c *gin.Context) {
	inputData := AuthData{}
	if err := c.Bind(&inputData); err != nil {
		fmt.Println("auth.go -> Authorization -> Bind: err = ", err)
		c.JSON(http.StatusInternalServerError, obj{"error": "wrong"})
		return
	}

	if err := inputData.checkPassword(); err != nil {
		fmt.Println("auth.go -> Authorization -> checkPassword: err = ", err)
		c.JSON(http.StatusInternalServerError, obj{"error": err.Error()})
		return
	}

	a.Login(c, inputData.Login)
}

func (ad *AuthData) checkPassword() error {
	var user User
	if err := storage.C("users").Find(obj{"email": ad.Login}).One(&user); err != nil {
		return err
	}

	return hashedPasswordCompare(ad.Password, user.Password)
}

func (r *Registration) passwordCompare() bool {
	return r.Password == r.PasswordSubmit
}

func (a *AuthModule) Login(c *gin.Context, email string) {
	fmt.Println("login")
	setCookie(c, "lib-login", fmt.Sprint(email, "*lib"))
	c.JSON(http.StatusOK, obj{"error": nil, "url": "/"})
}

func (a *AuthModule) Logout(c *gin.Context) {
	deleteCookie(c, "lib-login")
	pages.Auth(c)
}
