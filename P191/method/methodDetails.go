package main

import (
	"fmt"
	"math"
)

/*
方法调用和传参机制原理
说明：方法的调用和传参机制和函数的基本一样，不一样的地方是方法调用时，会将调用方法的变量，当作实参也传递给方法。
	也就是，用p去调用方法是，p自身这个结构体也会被拷贝到方法空间里

func （receiver type） methodName(参数列表)（返回值列表）{
		方法体
		return
}

注意事项和细节
1.结构体类型是值类型，在方法调用中，遵守值类型的传递机制，是值拷贝传递方式
2.如程序员希望在方法中，修改结构体变量的值，可通过结构体指针的方式来处理
3.go 中的方法是作用在指定的数据类型上的，所以自定义类型也可以绑定方法，而不仅仅是struct，如 int float
4.方法的访问范围控制的规则，和函数一样。方法名首字母大写才能跨包使用。
5.如果一个变量实现了String()这个方法，那么fmt.Println默认会调用这个变量的String()进行输出

fixme 一个结构体下的方法，统一要么使用"值接收者"要么使用"指针接收者",80% 使用指针
todo 当数据很小或者要保证它只读时，可以使用值类型
*/

type Circle struct {
	radius float64
}
type Dog struct {
	Name string
}

func (c *Circle) area() float64 {
	fmt.Printf("c 的地址是 %p \n", c)
	c.radius = 20 // fixme 这样传地址的话，外部的值也会被改变
	return math.Pi * c.radius * c.radius
	// 等价于 return math.Pi * (*c).radius * (*c).radius
	// go 底层做了优化，所以不需要我们再去取址
}

func (d Dog) bark() {
	fmt.Printf("方法里的d指向的地址是：%p \n", &d)
	// fixme 可以看到，如果不是指针类型的话，这里的值传递会给d一个新的内存空间
	fmt.Println(d)
	fmt.Println("wang wang wang !")
}

func testCircle() {
	circle := new(Circle)
	fmt.Printf("circle的地址是：%p \n", circle)
	circle.radius = 10 // 方法的类型是可以为指针类型的，那么你调用的对象就也要是指针类型的struct
	area := circle.area()
	fmt.Println(area)
	fmt.Println(circle.radius)

	dog1 := new(Dog)
	dog1.Name = "aaa"
	fmt.Printf("dog1指向地址是%p \n", dog1)
	dog1.bark()
}

type Integer int

func (i Integer) test() {
	fmt.Println("这不是普通的int哦")
}

func testInt() {
	inte := Integer(1)
	inte.test()
	fmt.Printf("类型是%T，值为%d \n", inte, inte)
}

type Student struct {
	Name string
	Age  int
}

// 给 student 实现String 方法
func (s *Student) String() string {
	str := fmt.Sprintf("这里是[%v],年龄[%d]，yo yo yo,whats up. \n", s.Name, s.Age)
	return str
}

func testString() {
	stu1 := Student{"cjj", 23}
	fmt.Println(stu1)
	fmt.Println(&stu1)
}
