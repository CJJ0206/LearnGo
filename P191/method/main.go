package main

import "fmt"

/*
8.12
方法基本介绍
在某些情况下，我们需要声明（定义）方法。比如person结构体：除了有一些字段外，还可以有一些行为，说话 、跑步、吃饭等
此时就需要用方法完成

fixme go 中的方法是作用在指定的数据类型上（即和指定的数据类型绑定），因此自定义的数据类型也可以有方法，而不仅仅是struct

方法的声明和调用
type A struct{
	Num int
}

func (a A) test(){
	fmt.Println(a.Num)
}
fixme func (a A) test 表示 A 结构体有一个方法，名字叫 test
fixme （a A）体现 test 和 A 类型是绑定的,a 是后续随意一个实例来调用这个方法的

*/

type Person struct {
	Name string
}

// p 表示接收者变量，这个实例来调用方法
// fixme 这里的 p 接收的是传过来的值拷贝
func (p Person) speak() {
	p.Name = "jack" // 所以这的修改是不会影响原值的
	fmt.Println(p.Name, "can speak.")
}

// fixme 注意这里的 *Person，代表接收者是一个指针,想修改原值用指针
//func (p *Person) speak2() {
//	p.Name = "jack" // 这里修改的直接是“原件”
//	fmt.Println(p.Name, "is a man.")
//}
// 传的是地址，适合需要修改原实例数据的操作，todo 或者当结构体非常大时（避免拷贝带来的性能消耗）

// fixme func （p Person） test(){...}   p 表示哪个Person变量调用这个p就是谁的副本（值拷贝）

func main() {
	person1 := Person{"cjj"}
	person1.speak() // 不能直接无实例调用，也不能用其他类型调用
	fmt.Println(person1.Name)

	test1()
	testCircle()
	testInt()
	testString()

	testMethod()
	testOddEven()
	cal()

	testCat()
	testTest03()

}
