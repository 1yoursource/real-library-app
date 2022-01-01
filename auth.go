package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"lib-client-server/auto_incrementer"
	"lib-client-server/client/models"
	"net/http"
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

	if inputData.checkEmpty() {
		c.JSON(http.StatusInternalServerError, obj{"error": "please, fill all field"})
		return
	}

	if !inputData.passwordCompare() {
		c.JSON(http.StatusInternalServerError, obj{"error": "please, use equals passwords"})
		return
	}

	fmt.Printf("inputData = %+v\n", inputData)

	var id = auto_incrementer.AI.Next("users")

	if err := user.CreateUser(inputData, id); err != nil {
		fmt.Println("auth.go -> Registration -> CreateUser: err = ", err)
		c.JSON(http.StatusInternalServerError, obj{"error": err.Error()})
		return
	}

	a.Login(c, inputData.Email, id)
}

func (a *AuthModule) Authorization(c *gin.Context) {
	inputData := AuthData{}
	if err := c.Bind(&inputData); err != nil {
		fmt.Println("auth.go -> Authorization -> Bind: err = ", err)
		c.JSON(http.StatusInternalServerError, obj{"error": "wrong"})
		return
	}

	var user models.User
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

func (r *Registration) checkEmpty() bool {
	if len(r.Email) == 0 {
		return true
	}
	if len(r.Password) == 0 {
		return true
	}
	if len(r.PasswordSubmit) == 0 {
		return true
	}
	if len(r.FirstName) == 0 {
		return true
	}
	if len(r.LastName) == 0 {
		return true
	}
	if len(r.Faculty) == 0 {
		return true
	}
	return false
}

func (a *AuthModule) Login(c *gin.Context, email string, id uint64) {
	fmt.Println("login")
	//setCookie(c, "lib-customer", fmt.Sprint("user", "*lib"))
	setCookie(c, "lib-login", fmt.Sprint(email, "*lib"))
	setCookie(c, "lib-id", fmt.Sprint(id))
	c.JSON(http.StatusOK, obj{"error": nil, "url": "/usr/main"})
}

func (a *AuthModule) Logout(c *gin.Context) {
	deleteCookie(c, "lib-login")
	//deleteCookie(c, "lib-customer")
	deleteCookie(c, "lib-id")
	//pages.Auth(c)
	userPages.Auth(c)
}
