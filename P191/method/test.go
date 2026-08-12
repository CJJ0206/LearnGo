package main

import "fmt"

/*
编写结构体（MethodUtils）,写一个不需要参数的方法，在方法中打印一个10*8的矩形

编写一个方法，提供m和n两个参数，方法中打印一个m*n的矩形

编写一个方法计算该矩形的面积（接收长宽len,wid）,将其作为方法返回值，接收返回的面积并打印

编写方法判断一个数是奇数还是偶数

根据行、列、字符打印出 行*列 的 某符号

定义小小计算器结构体，实现加减乘除
	实现一：分四个方法
	实现二：用一个方法
*/

type MethodUtils struct{} // fixme 结构体可以是空的

func (mu MethodUtils) Method1() {
	for range 10 { // fixme 要么直接range
		for _ = range 8 { // fixme 不用 : 直接 _
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func (mu MethodUtils) Method2(m, n int) {
	for range m {
		for range n {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func (mu MethodUtils) area(len, wid int) int {
	return wid * len
}

func testMethod() {
	mu1 := MethodUtils{}
	mu1.Method1()
	mu1.Method2(10, 20)
	fmt.Println(mu1.area(10, 20))
}

// -----------------------------------------------------------------------

type Num struct {
	I int
}

func (num Num) oddEven() {
	if num.I%2 == 0 {
		fmt.Println(num.I, "is even")
	} else {
		fmt.Println(num.I, "is odd")
	}
}

func testOddEven() {
	num1 := Num{1}
	num1.oddEven()
}

// --------------------------------------------------------------------------

type Calculate struct {
	a, b float64
}

func (cal *Calculate) getRes(operator byte) float64 {
	res := 0.0
	switch operator {
	case '+':
		res = cal.a + cal.b
	case '-':
		res = cal.a - cal.b
	case '*':
		res = cal.a * cal.b
	case '/':
		res = cal.a / cal.b
	default:
		return 0
	}
	return res
}

// fixme 一个结构体下的方法，统一要么使用"值接收者"要么使用"指针接收者",80% 使用指针
func (cal *Calculate) prt(l, w int, c byte) {
	for range l {
		for range w {
			fmt.Printf("%c", c) // %c 是用来打印字符的，%v是原始值
		}
		fmt.Println()
	}
}

func cal() {
	f1 := new(Calculate)
	f1.a = 10
	f1.b = 20
	fmt.Println(f1.getRes('+'))
	fmt.Println(f1.getRes('-'))
	fmt.Println(f1.getRes('*'))
	fmt.Println(f1.getRes('/'))

	f1.prt(5, 10, '&')

}
