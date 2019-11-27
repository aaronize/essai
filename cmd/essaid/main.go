package main

import (
    "github.com/gin-gonic/gin"
    "log"
    "time"
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
            "message": "success!",
        })
    })
    r.GET("/info", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "code": 200,
            "message": "success",
            "data": map[string]interface{} {
                "service": "essai-api",
                "create_at": time.Now(),
                "update_at": time.Now(),
                //"admin": "aaron.test",
                "name": "aaron.chen",
                "location": "shanghai",
                "region": "East Asia",
                "company": "None",
            },
        })
    })
    if err := r.Run(":13030"); err != nil {
        log.Fatal(err)
    }
}
