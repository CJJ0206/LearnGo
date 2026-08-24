package main

import (
	"fmt"
	"strings"
)

func main() {
	f := MakeSuffix() // fixme 这个我写的很鸡肋
	fmt.Println(f("english"))
	fmt.Println(f("korea.png"))
	fmt.Println(f("chinese.jpg"))

	// fixme 制造专门的具体闭包函数，就像是打包好一样
	makeJpg := MakeSuffix2(".jpg") // 制造一个专门加 .jpg 的函数
	makePng := MakeSuffix2(".png") // 制造一个专门加 .png 的函数
	fmt.Println(makeJpg("avatar")) // 输出: avatar.jpg
	fmt.Println(makePng("avatar")) // 输出: avatar.png

	// fixme 创建的是两个独立的闭包实例所以不会互相影响
	n := NewCounter()
	n2 := NewCounter()
	fmt.Println(n())
	fmt.Println(n2())
	fmt.Println(n())
	fmt.Println(n2())

	Sum(10, 20)
	// fixme defer 将语句入栈时，同时拷贝了相关变量的值（值快照）
	// fixme 即使后续 n1++、n2++，defer 中输出的仍然是入栈时的值 10 和 20
	// fixme 但 res 的计算用的是修改后的值：(10+1) + (20+1) = 32
	Sum2(10, 20)

}

// MakeSuffix 函数，接收文件后缀名（如 .jpg），返回一个闭包。闭包接收文件名，如果文件名已有该后缀则直接返回，没有则加上后缀。
func MakeSuffix() func(file string) string {
	suffix := ".jpg"
	return func(file string) string {
		ok := strings.HasSuffix(file, suffix)
		if ok {
			return file
		}
		return file + suffix
	}
}

// MakeSuffix2 接收一个参数作为闭包的“环境配置”
func MakeSuffix2(suffix string) func(file string) string {
	return func(file string) string {
		if strings.HasSuffix(file, suffix) {
			return file
		}
		return file + suffix
	}
}

// NewCounter 闭包计数器 ，每次调用返回递增的值（从 1 开始）
func NewCounter() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// Sum fixme defer 延迟机制 (也是 栈 先进后出)
func Sum(n1, n2 int) int {
	defer fmt.Println("ok1 n1 =", n1)
	defer fmt.Println("ok2 n2 =", n2)
	res := n1 + n2
	fmt.Println("ok3 res =", res)
	return res
}

// Sum2 defer 将语句入栈时，参数的值是"快照"还是"最新值"？
func Sum2(n1, n2 int) int {
	defer fmt.Println("ok1 n1 =", n1) // 输出什么？
	defer fmt.Println("ok2 n2 =", n2) // 输出什么？
	n1++
	n2++
	res := n1 + n2
	fmt.Println("ok3 res =", res) // 输出什么？
	return res
}

// ok3 res = 32
// ok2 n2 = 20
// ok1 n1 = 10
