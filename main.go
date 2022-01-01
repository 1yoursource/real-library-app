package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	r = gin.New()
)

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, os.Kill, syscall.SIGTERM)
	go func() {
		// Запускается при закрытии программы, для сохранения данных
		<-sigs

		fmt.Println("application shutdown signal received")

		os.Exit(1)
	}()

	gin.SetMode(gin.ReleaseMode)

	r.LoadHTMLGlob("templates/*.html")
	r.Static("/static", "./static")

	CreateModules()

	//r.GET("/", pages.Index)
	//r.GET("/search", pages.Search)
	//r.GET("/about", pages.About)
	//r.GET("/shell", pages.Shell)
	//r.GET("/d", pages.D)
	//r.GET("/logout", auth.Logout)
	//r.GET("/auth", pages.Auth)
	r.POST("/ajax/:module/:method", ajax)

	r.POST("/ajax2/:customer/:module/:method", ajax2)
	r.GET("/adm/:page", checkIsLoginAdm, adminPages.Handler)
	r.GET("/usr/:page", checkIsLoginUsr, userPages.Handler)

	fmt.Println("Application started on port 2200")

	if err := r.Run(":2200"); err != nil {
		fmt.Println("main.go -> main: err = ", err)
	}
	//scheduler.Every().Day().At("00:30:00").Run(MakeDebtorsList)
}
func ajax(c *gin.Context) {
	c.Header("Expires", time.Now().String())
	c.Header("Cache-Control", "no-cache")
	switch c.Param("module") {
	case "auth":
		auth.Ajax(c)
	default:
		c.String(http.StatusBadRequest, "Module not found!")
	}
}
func ajax2(c *gin.Context) {
	c.Header("Expires", time.Now().String())
	c.Header("Cache-Control", "no-cache")
	switch c.Param("customer") {
	case "adm":
		adm(c)
	case "usr":
		usr(c)
	default:
		c.String(http.StatusBadRequest, "restricted")
	}
}

func adm(c *gin.Context) {
	switch c.Param("module") {
	case "auth":
		adminAuth.Ajax(c)
	case "book":
		adminBooks.Handler(c)
	case "user":
		adminBooks.Handler(c)
	default:
		c.String(http.StatusBadRequest, "Module not found!")
	}
}

func usr(c *gin.Context) {
	switch c.Param("module") {
	case "book":
		userBooks.Handler(c)
	default:
		c.String(http.StatusBadRequest, "Module not found!")
	}
}
func checkIsLoginAdm(c *gin.Context) {
	checkIsLogin(c, "lib-admin")
}

func checkIsLoginUsr(c *gin.Context) {
	checkIsLogin(c, "lib-login")
}

func checkIsLogin(c *gin.Context, cookieName string) {
	var page = c.Param("page")

	if !needCheckLogin(page) {
		return
	}

	coockie, err := getCookie(c, cookieName)

	var (
		emptyCookies = len(coockie) == 0
		code         = http.StatusMovedPermanently
		location     = "auth"
	)

	switch {
	case err != nil && !strings.Contains(err.Error(),"http: named cookie not present"):
		break
	case page == "logout" && emptyCookies: // якщо ти не залогінений, вихід повинен бути недоступний, редірект на головну
		location = "main"
		break
	case page == "auth" && emptyCookies:
		return
	case emptyCookies:
		break
	case page == "auth": // якщо ти вже залогінений, сторінка аутентифікації повинна бути недоступна, редірект на головну
		code, location = http.StatusFound, "main"
		break
	default:
		return
	}
	c.Redirect(code, location)

}

func needCheckLogin(page string) bool {
	switch page {
	case "about", "main":
		return false
	default: // search, shell, auth, logout
		return true
	}
}
