package main

import (
	"github.com/wanmii/TestServer/iInterface"
	"github.com/wanmii/TestServer/jwt"
)

//使用
func main() {

	jwt.Test()
	iInterface.Test()

	//router.MakeRouter()
}
