package main

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/pkg/errors"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/sha3"
)

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
		return "", errors.New("empty cookies")
	}

	var loginCookie = name == "lib-login"

	if loginCookie && !strings.Contains(data, "lib") {
		return "", errors.New("")
	}

	if loginCookie {
		data = strings.Split(data, "*")[0] // при обращении к [0] никогда не паникнёт, из-за проверки выше (пример куки - abrakadabra18@gmail.com*lib)
	}

	return data,nil
}

func setCookie(c *gin.Context, name string, value string) {
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     name,
		Value:    value,
		Path:     "/",
		SameSite: http.SameSiteLaxMode,
		MaxAge:   coockieMaxAge,
		Expires:  time.Now().AddDate(zero, six, zero),
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

func H3hash(s string) string {
	h3 := sha3.New512()
	io.WriteString(h3, s)
	return fmt.Sprintf("%x", h3.Sum(nil))
}

// func for temlate
func toLower(name string) string {
	return strings.ToLower(name)
}

// func for temlate
func toUpper(name string) string {
	return strings.ToUpper(name)
}

func GetRand(count float64) int {
	rand.Seed(time.Now().UnixNano())
	return int(rand.Float64() * count)
}

//
/*

	clientStorage.Get(bson.NewObjectId())
	data, err := clientStorage.GetByQuery(models.Obj{}) // коллекция книг
	if err != nil {
		fmt.Println("err")
		return
	}
	book, isIt:=type_getter.GetTypeBook(data)
	if !isIt {
		return
	}
*/
