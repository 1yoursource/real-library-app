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

	//r.GET("/aaa", func(context *gin.Context) {
	//	id, _ := uuid.NewGen().NewV4()
	//	db := storage.DB
	//
	//	//db.AutoMigrate(&Test{}) // миграция (создание таблицы)
	//
	//	//db.Create(&Test{Name: "name"})
	//
	//	a := Test{}
	//	arr := []Test{}
	//	r := db.Find(&a)
	//
	//	// CReate
	//	fmt.Println("id : ", id)
	//	db.Create(&Test{Id: id, Name: "ppp"})
	//
	//	//Find
	//	r = db.Find(&a)
	//
	//	fmt.Println("here: ", r.Error)
	//
	//	fmt.Println("here 1: ", r.RowsAffected)
	//
	//	fmt.Println("here 2: ", a)
	//
	//	// delete
	//	//db.Where("name = ?", "ppp").Delete(&Test{})
	//	//db.Delete(&Test{Id: id}) // работает только с primary
	//	//db.Delete(&Test{}, id) // работает только по праймари ключу
	//
	//	//update
	//	db.Model(&Test{}).Where("name = ? AND id != ?", "ppp", id).Update("name", "polkijuyhtg")
	//	//db.Where("id = ?", id).Update("name", "")
	//
	//	//Find
	//	r = db.Find(&arr)
	//	fmt.Println("here: ", r.Error)
	//
	//	fmt.Println("here 1: ", r.RowsAffected)
	//
	//	fmt.Println("here 2: ", arr)
	//
	//	// &Test{id, "name"}
	//})

	r.GET("/drop", func(context *gin.Context) {
		db := storage.DB
		db.Migrator().DropTable(&Test{}, &Data{})
	})

	r.GET("/a1", func(context *gin.Context) {
		//id, _ := uuid.NewGen().NewV4()
		db := storage.DB

		//db.Migrator().DropTable(&Test1{}, Data{}) // миграция (создание таблицы)
		db.Migrator().CreateTable(&Test{}) // миграция (создание таблицы)
		db.Migrator().CreateTable(&Data{}) // миграция (создание таблицы)
		//db.Migrator().CreateConstraint(&Test{}, "Data")
		//db.Migrator().CreateConstraint(&Test{}, "fk_tests_data")

		db.Create(&Test{Name: "name", Data: Data{Info: "Info"}})
		//db.Create(&Data{Info: "Info"})

		//a := Test1{}
		arr := []Test{}
		r := db.Find(&arr)

		// CReate
		//fmt.Println("id : ", id)
		//db.Create(&Test1{Id: id, Name: "ppp"})

		//Find
		//r = db.Find(&a)

		fmt.Println("here: ", r.Error)

		fmt.Println("here 1: ", r.RowsAffected)

		fmt.Printf("here 2: %+v\n", arr)

		// delete
		//db.Where("name = ?", "ppp").Delete(&Test{})
		//db.Delete(&Test{Id: id}) // работает только с primary
		//db.Delete(&Test{}, id) // работает только по праймари ключу

		//update
		//db.Model(&Test1{}).Where("name = ? AND id != ?", "ppp", id).Update("name", "polkijuyhtg")
		//db.Where("id = ?", id).Update("name", "")

		//Find
		//r = db.Find(&arr)
		//fmt.Println("here: ", r.Error)
		//
		//fmt.Println("here 1: ", r.RowsAffected)
		//
		//fmt.Println("here 2: ", arr)

		// &Test{id, "name"}
	})

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
