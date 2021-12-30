package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type (
	PagesModule struct {
	}
)

func (p *PagesModule) Index(c *gin.Context) {
	templateData := gin.H{
		"notLogin": !p.checkIsLogin(c),
	}
	c.HTML(http.StatusOK, "index.html", templateData)
}

func (p *PagesModule) Search(c *gin.Context) {
	templateData := gin.H{
		"notLogin": !p.checkIsLogin(c),
	}
	c.HTML(http.StatusOK, "search.html", templateData)
}

func (p *PagesModule) About(c *gin.Context) {
	templateData := gin.H{
		"notLogin": !p.checkIsLogin(c),
	}
	c.HTML(http.StatusOK, "about.html", templateData)
}

func (p *PagesModule) Shell(c *gin.Context) {
	templateData := gin.H{
		"notLogin": !p.checkIsLogin(c),
	}
	c.HTML(http.StatusOK, "shell.html", templateData)
}

func (p *PagesModule) D(c *gin.Context) {
	templateData := gin.H{
		"notLogin": !p.checkIsLogin(c),
	}
	c.HTML(http.StatusOK, "distribution-map.html", templateData)
}

func (p *PagesModule) Auth(c *gin.Context) {
	templateData := gin.H{
		"notLogin": !p.checkIsLogin(c),
	}
	c.HTML(http.StatusOK, "registration.html", templateData)
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

func (p *PagesModule) checkValidEmail(email string) bool {
	data := &Registration{Email: email}
	err := data.CheckEmail()
	fmt.Println("DATA: ", data)
	fmt.Println("err =  ", err)
	if err != nil && strings.Contains(err.Error(), "already registered") {
		fmt.Println("1111111111111")
		return true
	}
	fmt.Println("22222222222222222")
	return false
}
