package main

import (
	"fmt"
	"testing"
)

var depth int = 0

func tabs() {
	for i := 0; i < depth; i++ {
		fmt.Printf("\t")
	}
}
func Test000(t *testing.T) {
	strings := parse("sample/sample.go")
	for k, v := range strings {
		fmt.Println("key = ", k, ", val = ", v)
	}
}
