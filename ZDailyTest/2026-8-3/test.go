package main

import (
	"fmt"
)

func main() {
	// fixme 匿名函数的直接创建并直接使用，但是这样只能调用一次
	res := func(a, b int) int {
		return a + b
	}(10, 20)
	fmt.Println(res)

	// fixme 第二种就是把匿名函数赋值给一个变量，这样才可以多次调用
	fc := func(a, b int) int { return a * b }
	fmt.Println(fc(2, 6))
	fmt.Println(fc(4, 6))

	// strconv.Atoi("123") 和 strconv.Atoi("abc") 分别返回什么？为什么不 panic？
	// Atoi("123") → (123, nil)
	// Atoi("abc") → (0, error)
	// Go 的设计哲学是不 panic 而是返回 error，让调用者自己决定如何处理错误。Atoi 返回 (int, error)，不会 panic。

	// fixme a 是返回出来的匿名函数
	a := AddUpper()
	// fixme 这里用的是（），不是直接用AddUpper
	// fixme 所以a不是一个实例而是执行AddUpper返回的匿名函数
	fmt.Println(a(1)) // fixme 故而可以直接传参
	fmt.Println(a(2))
	fmt.Println(a(3))
}

// AddUpper 累加器,函数的返回类型是一个 func(int) int
func AddUpper() func(int) int {
	n := 0
	return func(x int) int {
		// 调用AddUpper产生的是一个 加法运算机 func(int) int
		n += x
		return n // func 函数再返回结果，func会带着n逃逸，才能实现累加
		// 因为 func 把 n 返回到了add外部，所以编译器会认为n的生命周期应该比add长而保留了它
	}
}
