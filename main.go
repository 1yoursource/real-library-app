package main

import (
	"fmt"
	"github.com/carlescere/scheduler"
	"net/http"
	"os"
	"os/signal"
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

	r.GET("/", pages.Index)
	r.GET("/search", pages.Search)
	r.GET("/about", pages.About)
	r.GET("/shell", pages.Shell)
	//r.GET("/d", pages.D)
	r.GET("/logout", auth.Logout)
	r.GET("/auth", pages.Auth)
	r.POST("/ajax/:module/:method", ajax)

	r.POST("/ajax2/:customer/:module/:method", ajax2)
	r.GET("/adm/:page", adminPages.Handler)
	r.GET("/usr/:page", userPages.Handler)

	fmt.Println("Application started on port 2200")

	if err := r.Run(":2200"); err != nil {
		fmt.Println("main.go -> main: err = ", err)
	}
	scheduler.Every().Day().At("00:30:00").Run(MakeDebtorsList)
}
func ajax(c *gin.Context) {
	c.Header("Expires", time.Now().String())
	c.Header("Cache-Control", "no-cache")
	switch c.Param("module") {
	case "auth":
		auth.Ajax(c)
	case "booking":
		user.Ajax(c)
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
