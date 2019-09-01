package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"weixin/src/db"
)

func main() {
	r := gin.Default()
	r.Use(Cors())
	r.POST("/wexin/call", func(c *gin.Context) {
		var message db.Message
		c.ShouldBind(&message)
		message.Save()
	})
	r.GET("/wexin/getAllNames", func(c *gin.Context) {
		names := db.SelectAllNames()
		c.JSON(200, gin.H{
			"status": "0",
			"result": names,
		})
	})
	r.GET("/wexin/getNameLine", func(c *gin.Context) {
		name := c.Query("name")
		line := db.SelectNameLine(name)
		fmt.Printf("sdf", line)
		c.JSON(200, gin.H{
			"status": "0",
			"result": line,
		})
	})
	r.GET("/wexin/test", func(c *gin.Context) {
		db.Test()
		c.JSON(200, gin.H{
			"status": "0",
			"result": 123,
		})
	})
	r.GET("/wexin/getAllGroups", func(c *gin.Context) {

		name := c.Query("name")
		names := strings.Split(name, ",")
		groups := db.SelectAllGroups(names)
		c.JSON(200, gin.H{
			"status": "0",
			"result": groups,
		})
	})
	r.Run("0.0.0.0:8080") // listen and serve on 0.0.0.0:8080
}

// 处理跨域请求,支持options访问
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}
