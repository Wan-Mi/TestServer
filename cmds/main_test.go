package main

//gotests -all -w ./cmds/main.go

import (
	"testing"
)

func Test_getTest(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{"my"}, {"name"}, {"is"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			getTest(tt.name)
		})
	}
}

// func Test_main(t *testing.T) {
// 	tests := []struct {
// 		name string
// 	}{
// 	// TODO: Add test cases.
// 	}
// 	for range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			main()
// 		})
// 	}
// }
