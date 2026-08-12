package main

import "fmt"

/*
方法和函数的区别
1.调用方式不一样
	函数调用方式 ： 函数名（实参列表）
	方法调用方式 ： 变量.方法名（实参列表）
2.对于普通函数，接收者为值类型时，只能传值，反之亦然
3.对于方法（如struct）接收者为值类型时，可以直接用指针类型的变量调用方法，反之亦然
*/

type Cat struct {
	Color string
}

func (c Cat) col() {
	fmt.Printf("this is a cat of %v color", c)
}

func testCat() {
	cat1 := new(Cat)
	cat1.Color = "red"
	cat1.col()
}

// ---------------------------------------------------------
func (c Cat) test03() {
	c.Color = "black"
	fmt.Println("猫的颜色是：", c.Color)
}
func testTest03() {
	c := Cat{Color: "blue"}
	c.test03()
	(&c).test03() // 当这样时，其实做的还是值传递，主要看方法里面的类型
	// fixme 语法上支持这种写法，但是底层会安全地转为方法需要的类型
	fmt.Println(c.Color)
}
