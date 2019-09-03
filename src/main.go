package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"weixin/src/db"
	"weixin/src/redis"
)

type Person struct {
	Name string `form:"name"`
}

func ToAscll(str string) string {
	textQuoted := strconv.QuoteToASCII(str)
	textUnquoted := textQuoted[1 : len(textQuoted)-1]
	fmt.Println(textUnquoted)
	re3, _ := regexp.Compile(`\\u....`)
	rep := re3.ReplaceAllStringFunc(textUnquoted, strings.ToUpper)
	re4, _ := regexp.Compile(`\\u`)
	rep = re4.ReplaceAllStringFunc(textUnquoted, strings.ToLower)
	return rep
}

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
		name := ToAscll(redis.GET("teacher"))
		names := strings.Split(name, ",")
		groups := db.SelectAllGroups(names)
		c.JSON(200, gin.H{
			"status": "0",
			"result": groups,
		})

	})
	r.GET("/wexin/getTeacher", func(c *gin.Context) {
		teacher := redis.GET("teacher")
		c.JSON(200, gin.H{
			"status": "0",
			"result": teacher,
		})
	})
	r.POST("/wexin/setTeacher", func(c *gin.Context) {
		var teacher Person
		c.ShouldBind(&teacher)
		redis.SET("teacher", teacher.Name)
		c.JSON(200, gin.H{
			"status": "0",
			"result": "ok",
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
