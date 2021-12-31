package helper

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/sha3"
	"io"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

const (
	coockieMaxAge = 32000000

	zero = 0
	six  = 6
	)

// createHashPassword - создает хэшированный пароль
func CreateHashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func HashedPasswordCompare(password string, hashedPassword []byte) error {
	return bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
}

func GetCookie(c *gin.Context, name string) (string, error) {
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

func SetCookie(c *gin.Context, name string, value string) {
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     name,
		Value:    value,
		Path:     "/",
		SameSite: http.SameSiteLaxMode,
		MaxAge:   coockieMaxAge,
		Expires:  time.Now().AddDate(zero, six, zero),
	})
}

func DeleteCookie(c *gin.Context, name string) {
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
func ToLower(name string) string {
	return strings.ToLower(name)
}

// func for temlate
func ToUpper(name string) string {
	return strings.ToUpper(name)
}

func GetRand(count float64) int {
	rand.Seed(time.Now().UnixNano())
	return int(rand.Float64() * count)
}