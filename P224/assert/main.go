package main

import "fmt"

type Point struct {
	x int
	y int
}

func main() {
	// var a interface{}  // fixme 这两行是完全等价的
	var a any
	point := Point{1, 2}
	a = point // fixme 任何值都可以赋值给空接口!!!!!!! 所以a声明那一行说可以替换interface为any
	// 如何将a赋值给一个Point变量
	var b Point
	// b = a // 可以吗？ 报错是，不能使用一个空接口类型作为Point类型
	b = a.(Point)
	// fixme 有人会觉得这不就是强转吗，但这里的逻辑只是帮判断了一下里面这个到底是不是一个Point，不能的话还是会报错
	fmt.Println(b)

	// fixme 现在类型断言的默认返回值已经是两个了，天然附带bool判断，不接收这个bool会panic
	y, _ := a.(float32)
	// 当然也可以通过这个bool写一些判断过程
	fmt.Printf("y的类型是%T，y的值是%v", y, y) // fixme 即使是转换失败，类型任然会变，但是值为零值

	// 直接写在if判断条件上，但是这样y的作用域就只在if里，如果后续要使用的话还是在外面写把
	if y, ok := a.(float64); ok {
		fmt.Printf("y的类型是%T，y的值是%v", y, y)
		fmt.Println("断言成功")
	} else {
		fmt.Println("断言失败")
	}
}
