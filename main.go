package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/bonjour", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code": 200,
			"message": "request success!",
		})
	})
	if err := r.Run(":13030"); err != nil { 
		log.Fatal(err)
	}
}
