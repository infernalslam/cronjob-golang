package health

import (
	"github.com/gin-gonic/gin"
)

// Register *
func Register(ctx *gin.Engine) {
	ctx.GET("/", fetchTodos)
}

func fetchTodos(c *gin.Context) {
	todosList := LoopTodos(20)
	c.JSON(200, gin.H{
		"data": todosList,
	})
}
