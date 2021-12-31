package customer

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/night-codes/types.v1"
	"lib-client-server/client/helper"
	"lib-client-server/client/models"
	"lib-client-server/database"
	"net/http"
	"strings"
)

type (
	PagesModule struct {
		pagePrefix string
	}
)

func CreateUserPagesModule(pagePrefix string) *PagesModule {
	return &PagesModule{pagePrefix: pagePrefix}
}

func (p *PagesModule) Handler(c *gin.Context) {
	pg := c.Param("page")
	switch pg {
	case "main":
		p.Index(c)
	case "auth":
		p.Auth(c)
	case "search":
		p.Search(c)
	case "about":
		p.About(c)
	case "shell":
		p.Shell(c)
	default:
		c.JSON(http.StatusNotFound, models.Obj{"error": "page does't exist"})
	}
}

func (p *PagesModule) Index(c *gin.Context) {
	var (
		isLogin = p.checkIsLogin(c)
	)
	templateData := gin.H{
		"isLogin": isLogin,
	}

	if isLogin {
		templateData["userId"] = p.getUserId(c)
	}

	c.HTML(http.StatusOK, fmt.Sprint(p.pagePrefix, "_", "index.html"), templateData)
}

func (p *PagesModule) Search(c *gin.Context) {
	var isLogin = p.checkIsLogin(c)

	templateData := gin.H{
		"isLogin": isLogin,
	}

	c.HTML(http.StatusOK, fmt.Sprint(p.pagePrefix, "_", "search.html"), templateData)
}

func (p *PagesModule) About(c *gin.Context) {
	templateData := gin.H{
		"isLogin": p.checkIsLogin(c),
	}
	c.HTML(http.StatusOK, fmt.Sprint(p.pagePrefix, "_", "about.html"), templateData)
}

func (p *PagesModule) Shell(c *gin.Context) {
	var isLogin = p.checkIsLogin(c)
	templateData := gin.H{
		"isLogin": isLogin,
	}

	if isLogin {
		fmt.Println("isLogin: ", isLogin)
		// отримуємо користувача
		if idStr, err := getCookie(c, "lib-id"); err == nil {
			var user models.User
			if err := database.Connect("localhost",
				"libDB",
				"libraryDatabase",
				).C("users").FindId(types.Uint64(idStr)).One(&user); err == nil {
				if user.Id != 0 {
					templateData["user"] = user
				} else {
					isLogin = false
				}
			}
		}
	}

	fmt.Println("aaaaaaaaaaa: ", templateData["user"])

	c.HTML(http.StatusOK, fmt.Sprint(p.pagePrefix, "_", "shell.html"), templateData)
}

func (p *PagesModule) Auth(c *gin.Context) {
	templateData := gin.H{
		"isLogin": p.checkIsLogin(c),
	}
	c.HTML(http.StatusOK, fmt.Sprint(p.pagePrefix, "_", "registration.html"), templateData)
}

func (p *PagesModule) checkIsLogin(c *gin.Context) bool {
	var isLogin bool
	if cookie, err := getCookie(c, "lib-login"); err == nil {
		if p.checkValidEmail(cookie) {
			isLogin = true
		}
	}
	return isLogin
}

func (p *PagesModule) getUserId(c *gin.Context) string {
	if cookie, err := getCookie(c, "lib-id"); err == nil {
		if p.checkValidId(cookie) {
			// todo add validation 	for correct id

			return strings.Split(cookie, "*")[0]
		}
	}
	return ""
}

func (p *PagesModule) checkValidId(id string) bool {
	return len(id) == len(id)
	// todo добавить норм проверку админ мыла
}

func (p *PagesModule) checkValidEmail(email string) bool {
	return len(email) == len(email)
	// todo добавить норм проверку админ мыла
}

func getCookie(c *gin.Context, name string) (string, error) {
	return helper.GetCookie(c, name)
}

func setCookie(c *gin.Context, name string, value string) {
	helper.SetCookie(c, name, value)
}

func deleteCookie(c *gin.Context, name string) {
	helper.DeleteCookie(c, name)
}
