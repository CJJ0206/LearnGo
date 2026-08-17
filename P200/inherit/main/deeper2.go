package main

import "fmt"

/*
8.17
结构体嵌入两个或者多个匿名结构体，如两个匿名结构体有相同的字段和方法
（同时结构体本身没有同名的字段和方法），再访问时，就必须明确指定匿名结构体的名字

如果在一个struct嵌套一个有名结构体，这种模式就是组合，如果是组合关系
那么在访问组合的结构体字段或方法时，必须带上结构体的名字

*/

type AA struct {
	Name string
	age  int
}
type BB struct {
	Name string
	age  int
}
type C struct {
	AA
	BB
}
type D struct {
	a AA
}

type goods struct {
	name  string
	price int
}

type brand struct {
	name    string
	address string
}

type tv struct {
	*goods
	*brand // 嵌套地址效率更高
}

func main() {
	var c C
	// fmt.Println(c.Name) // 如果出现多重继承里含有相同参数时，需要指定具体匿名结构体,否则报错
	fmt.Println(c.AA.Name) // 这样才对
	fmt.Println(c)

	var d D
	//d.Name
	//fmt.Println(d.AA.Name)  // 这两都是错的
	fmt.Println(d.a.Name) // 一定是通过a这个名字去调用（不是匿名结构体）
	// fixme 往上找对应的属性时，首先看的就是自己的结构体里有没有对应的，有就直接用，没有再去匿名结构体里找

	// todo 如果用指针类型的话，需要传地址，输出要解引用
	tv1 := &tv{
		goods: &goods{"tv", 3000},
		brand: &brand{"samsung", "korea"},
	}
	fmt.Println(*tv1.brand)
}
