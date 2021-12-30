package main

import (
	"fmt"
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
	rGroup := r.Group("/api/v1")

	rGroup.Group("/ajax/")
	//rGroup.GET("/books/set", handlerModule.CreateItem)
	//rGroup.GET("/books/get/by/user/:userId", handlerModule.GetBooksByUser)
	//rGroup.GET("/books/get/by/author/:authorName", handlerModule.GetBooksByAuthor)
	//rGroup.GET("/readers/get/expired", handlerModule.GetReadersWithExpiredBookTerm)

	fmt.Println("Application started on port 2200")

	if err := r.Run(":2200"); err != nil {
		fmt.Println("main.go -> main: err = ", err)
	}
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
