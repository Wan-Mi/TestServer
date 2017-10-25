package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var res string

func MakeRouter() {
	fmt.Println("gin is ready")
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		res = res + "a"
		c.Next() //处理请求

		res = res + "c   "

	}
}

func secondLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		res = res + "d" + c.Request.URL.Path
		s := "s"
		c.Next() //处理请求

		res = res + "e " + s + fmt.Sprintf("%d", c.Writer.Status)
	}
}

func init() {
	r := gin.Default()
	r.Use(Logger())
	r.Use(secondLogger())

	r.GET("/ginTest", func(c *gin.Context) {
		res = res + "b"
		c.String(200, res)
	})

	fmt.Println(res)
	r.Run(":8080")
}
