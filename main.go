package main

import (
	"github.com/gin-gonic/gin"
	"github.com/infernalslam/template-golang/health"
)

func main() {
	r := gin.Default()

	health.Register(r)

	r.Run()
}
