package goTest_test

import (
	"fmt"
	"testing"
)

func helloWorld() string {
	return "helloWorld"
}

func Test_HelloWord(t *testing.T) {
	fmt.Println("test begin:\n")
	t.Errorf("test result %s", helloWorld())
}
