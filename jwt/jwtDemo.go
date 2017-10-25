package jwt

import (
	"fmt"
)

func Test() {
	fmt.Println("jwt test run \n")
}

func init() {
	jwtStr := ""
	if token, err := genClairToken(131415926, "qazwsxEDCfgh", "abcdefghijklmnopqrstuvwxyz"); err != nil {
		fmt.Println(err)
	} else {
		jwtStr = token
		fmt.Println(token)
	}

	if id, str1, str2, times, err := parseClairToken(jwtStr); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(id, str1, str2, times)
	}
}
