//mockgen -source mockTest/mockDemo.go -destination ./mockTest/mockTest.go -package mockTest

package mockTest

import (
	"strconv"
)

type mockDemo interface {
	GetString() string
}

type intValue struct {
	a int
}

func (val intValue) GetString() string {
	return strconv.Itoa(val.a)
}

func GoMockDemo(a mockDemo) string {
	return a.GetString() + "45"
}
