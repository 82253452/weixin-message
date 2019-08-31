package main

import (
	"github.com/gin-gonic/gin"
	"weixin/src/db"
)

func main() {
	r := gin.Default()
	r.POST("/wexin/call", func(c *gin.Context) {
		var message db.Message
		c.ShouldBind(&message)
		message.Save()
	})
	r.Run("0.0.0.0:8001") // listen and serve on 0.0.0.0:8080
}
