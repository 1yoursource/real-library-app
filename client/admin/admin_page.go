package admin

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"lib-client-server/client/models"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
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
	default:
		c.JSON(http.StatusNotFound,models.Obj{"error":"page does't exist"})
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

	c.HTML(http.StatusOK, fmt.Sprint(p.pagePrefix,"_","index.html"), templateData)
}

func (p *PagesModule) Search(c *gin.Context) {
	var isLogin =  p.checkIsLogin(c)

	templateData := gin.H{
		"isLogin": isLogin,
	}

	c.HTML(http.StatusOK, fmt.Sprint(p.pagePrefix,"_","search.html"), templateData)
}

func (p *PagesModule) About(c *gin.Context) {
	templateData := gin.H{
		"isLogin": !p.checkIsLogin(c),
	}
	c.HTML(http.StatusOK, fmt.Sprint(p.pagePrefix,"_","about.html"), templateData)
}

func (p *PagesModule) Shell(c *gin.Context) {
	templateData := gin.H{
		"isLogin": !p.checkIsLogin(c),
	}
	c.HTML(http.StatusOK, fmt.Sprint(p.pagePrefix,"_","shell.html"), templateData)
}

func (p *PagesModule) D(c *gin.Context) {
	templateData := gin.H{
		"isLogin": !p.checkIsLogin(c),
	}
	c.HTML(http.StatusOK, fmt.Sprint(p.pagePrefix,"_","distribution-map.html"), templateData)
}

func (p *PagesModule) Auth(c *gin.Context) {
	templateData := gin.H{
		"isLogin": !p.checkIsLogin(c),
	}
	c.HTML(http.StatusOK, fmt.Sprint(p.pagePrefix,"_","registration.html"), templateData)
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

/////////////////
/////////////////
/////////////////
/////////////////
// todo ВРЕМЕННО


// createHashPassword - создает хэшированный пароль
func createHashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func hashedPasswordCompare(password string, hashedPassword []byte) error {
	return bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
}

func getCookie(c *gin.Context, name string) (string, error) {
	data, err := c.Cookie(name)
	if err != nil {
		return "", err
	}

	if len(data) == 0 {
		return "", errors.New("")
	}

	if !strings.Contains(data, "lib") {
		return "", errors.New("")
	}

	return strings.Split(data, "*")[0], nil // при обращении к [0] никогда не паникнёт, из-за проверки выше (пример куки - abrakadabra18@gmail.com*lib)
}

func setCookie(c *gin.Context, name string, value string) {
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     name,
		Value:    value,
		Path:     "/",
		SameSite: http.SameSiteLaxMode,
		MaxAge:   models.CoockieMaxAge,
		Expires:  time.Now().AddDate(0, 0, 1),
	})
}

func deleteCookie(c *gin.Context, name string) {
	http.SetCookie(c.Writer, &http.Cookie{
		Name:    name,
		Value:   "",
		Path:    "/",
		MaxAge:  1,
		Expires: time.Now(),
	})
}