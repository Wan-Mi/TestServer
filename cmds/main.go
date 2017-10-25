package main

import (
	"github.com/wanmii/TestServer/jwt"
	"github.com/wanmii/TestServer/router"
)

//使用
func main() {

	jwt.Test()
	router.MakeRouter()
}
