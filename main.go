package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var res string

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		res = res + "a"

		c.Next() //处理请求

		res = res + "c"

	}
}

func secondLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		res = res + "d"

		c.Next() //处理请求

		res = res + "e"
	}
}

//使用
func main() {
	r := gin.Default()
	r.Use(Logger())
	r.Use(secondLogger())

	r.GET("/", func(c *gin.Context) {
		res = res + "b"
		c.String(200, res)
	})

	fmt.Println(res)
	r.Run(":8080")
}