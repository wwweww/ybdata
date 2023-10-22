package controller

import (
	"github.com/gin-gonic/gin"
	"log"
)

func Test(c *gin.Context) {
	token := c.Request.Header.Get("x-token")

	log.Println("Token: " + token)

	c.JSON(200, gin.H{
		"token": token,
		"if":    "" == token,
	})
}
