package main

import (
	"github.com/wanmii/TestServer/jwt"
	"github.com/wanmii/TestServer/router"
)

//使用
func main() {
	router.MakeRouter()
	jwt.Test()
}
