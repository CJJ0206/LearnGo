package main

import (
	"fmt"
	"strings"
)

func testCut() {
	before, after, found := strings.Cut("I had a delicious buttered bread roll this morning.", "li")
	fmt.Println("Before:", before)
	fmt.Println("After:", after)
	fmt.Println("Found:", found)
}

func testCutPrefix() {
	name := "https://gemini.google.com"
	before, ok := strings.CutPrefix(name, "http://")
	if !ok {
		fmt.Println("格式错误，没有找到对应前缀")
	} else {
		fmt.Println("before:", before)
	}
}
