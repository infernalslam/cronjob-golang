package main

import (
	"github.com/gin-gonic/gin"
	"github.com/infernalslam/template-golang/health"
)

func main() {
	r := gin.Default()

	health.Register(r)
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	r.Run()
}
