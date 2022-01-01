package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"lib-client-server/client/helper"
	"lib-client-server/client/models"
	"net/http"
	"strings"
)

type (
	PagesModule struct {
		pagePrefix string
	}
)

func CreateAdminPagesModule(pagePrefix string) *PagesModule {
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
	case "logout":
		deleteCookie(c, "lib-login")
		deleteCookie(c, "lib-id")
		p.Auth(c)
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
		templateData["userId"] = p.getAdminId(c)
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
	templateData := gin.H{
		"isLogin": p.checkIsLogin(c),
	}
	c.HTML(http.StatusOK, fmt.Sprint(p.pagePrefix, "_", "shell.html"), templateData)
}

func (p *PagesModule) D(c *gin.Context) {
	templateData := gin.H{
		"isLogin": p.checkIsLogin(c),
	}
	c.HTML(http.StatusOK, fmt.Sprint(p.pagePrefix, "_", "distribution-map.html"), templateData)
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

func (p *PagesModule) getAdminId(c *gin.Context) string {
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
