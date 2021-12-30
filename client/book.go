package client

import (
	"github.com/gin-gonic/gin"
)

type (
	BookInterface interface {
		GetAll(c *gin.Context)
		Create(c *gin.Context) // a:cr,u:get
		Read(c *gin.Context)   // a:see,u:see info //todo NEED ?????
		Update(c *gin.Context) // a:edit,u:
		Delete(c *gin.Context) // a:del,u:return
	}
)