package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"lib-client-server/auto_incrementer"
	"lib-client-server/client/helper"
	"lib-client-server/client/models"
	"lib-client-server/database"
	"net/http"
)

type (
	AuthModule struct{}

	Registration struct {
		Email          string `form:"email"`
		Password       string `form:"password"`
		PasswordSubmit string `form:"passwordSubmit"`
		Code           string `form:"code"`
	}

	AuthData struct {
		Login    string `form:"login"`
		Password string `form:"password"`
	}

	Admin struct {
		Id       uint64 `bson:"_id"`
		Email    string `bson:"email"`
		Password []byte `bson:"password"`
	}
)

func CreateAuthModule() *AuthModule {
	return &AuthModule{}
}

func (a *AuthModule) Ajax(c *gin.Context) {
	fmt.Println("HEHHHAJSKDHJASJD")
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
		c.JSON(http.StatusInternalServerError, models.Obj{"error": "wrong"})
		return
	}

	if inputData.checkEmpty() {
		c.JSON(http.StatusInternalServerError, models.Obj{"error": "please, fill all field"})
		return
	}

	if !inputData.passwordCompare() {
		c.JSON(http.StatusInternalServerError, models.Obj{"error": "please, use equals passwords"})
		return
	}

	var id = auto_incrementer.AI.Next("admins")

	if err := inputData.CreateAdmin(id); err != nil {
		c.JSON(http.StatusInternalServerError, models.Obj{"error": "doesn't create"})
		return
	}


	a.Login(c, inputData.Email, id)
}

func (a *AuthModule) Authorization(c *gin.Context) {
	inputData := AuthData{}
	if err := c.Bind(&inputData); err != nil {
		fmt.Println("auth.go -> Authorization -> Bind: err = ", err)
		c.JSON(http.StatusInternalServerError, models.Obj{"error": "wrong"})
		return
	}

	var adm Admin
	if err := database.Connect("localhost",
		"libDB",
		"libraryDatabase",
	).C("admins").Find(models.Obj{"email": inputData.Login}).One(&adm); err != nil {
		fmt.Println("auth.go -> Authorization -> Find: err = ", err)
		c.JSON(http.StatusInternalServerError, models.Obj{"error": "not found"})
		return
		//return err
	}

	if err := inputData.checkPassword(adm.Password); err != nil {
		fmt.Println("auth.go -> Authorization -> checkPassword: err = ", err)
		c.JSON(http.StatusInternalServerError, models.Obj{"error": err.Error()})
		return
	}

	a.Login(c, inputData.Login, adm.Id)
}

func (ad *AuthData) checkPassword(userPassword []byte) error {
	return helper.HashedPasswordCompare(ad.Password, userPassword)
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
	return false
}

func (a *AuthModule) Login(c *gin.Context, email string, id uint64) {
	setCookie(c, "lib-admin", fmt.Sprint("admin-",id))
	setCookie(c, "lib-login", fmt.Sprint(email, "*lib"))
	setCookie(c, "lib-id", fmt.Sprint(id))
	c.JSON(http.StatusOK, models.Obj{"error": nil, "url": "/adm/main"})
}

func (a *AuthModule) Logout(c *gin.Context) {
	deleteCookie(c, "lib-login")
	deleteCookie(c, "lib-admin")
	deleteCookie(c, "lib-id")
	var p = &PagesModule{pagePrefix: "admin"}
	p.Auth(c)
}

func (r *Registration) CheckEmail() error {
	count, err := database.Connect("localhost",
		"libDB",
		"libraryDatabase",
	).C("admins").Find(models.Obj{"email": r.Email}).Count()
	switch {
	case err != nil:
		break
	case count > 0:
		err = fmt.Errorf("already registered")
	}

	return err
}

func (r *Registration) CreateAdmin(id uint64) error {
	if err := r.CheckEmail(); err != nil {
		return err
	}

	password, err := helper.CreateHashPassword(r.Password)
	if err != nil {
		return err
	}

	adm := Admin{
		Id:           id,
		Email:        r.Email,
		Password:     password,
	}

	if err := database.Connect("localhost",
		"libDB",
		"libraryDatabase",
	).C("admins").Insert(adm); err != nil {
		return err
	}

	return nil
}