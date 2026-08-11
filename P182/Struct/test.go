package main

// 8.11

import "fmt"

func test() {
	var p1 Person
	p1.age = 10
	p1.name = "cjj"
	var p2 = &p1

	fmt.Println((*p2).age)
	fmt.Println(p2.age)
	p2.name = "cjj~~"
	fmt.Println(p2.name)
	fmt.Println((*p2).name)

}
