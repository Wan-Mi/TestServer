package goTest

import (
	"fmt"
	"testing"
)

func helloWorld() string {
	return "helloWorld"
}

// normal test
func TestUnit(t *testing.T) {
	fmt.Println("test begin:")
	if helloWorld() != "helloWorld" {
		t.Errorf("test result: %v", helloWorld())
	} else {
		fmt.Println("test PASS")
	}
}

// benchmark test
func BenchmarkUnit(b *testing.B) {
	b.N = 100

	for i := 0; i < b.N; i++ {
		helloWorld()
	}
}
