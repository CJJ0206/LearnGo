package utils

// 8.3  8.4

import (
	"strconv"
	"strings"
)

// 闭包就是一个函数和其相关的引用环境组合的一个整体（实体）

// AddUpper 是函数声明，空的表示不用传任何东西；func(int) int 是 AddUpper 的返回值类型 （用了个匿名函数）
func AddUpper() func(int) (int, string) { // 所以这里意思是返回类型是一个函数
	var n int = 10
	var s string = "hello"
	return func(x int) (int, string) { // 调用AddUpper产生的是一个 加法运算机 func(int) int
		n += x
		s += strconv.Itoa(x)
		return n, s // func 函数再返回结果，func会带着n逃逸，才能实现累加
	}
}

// return 的是一个匿名函数，但这个函数引用到了函数外的 n,所以这个匿名函数就和n形成了一个整体，构成闭包
// 在 Go 语言中，闭包的构成可以用一个非常简单的“公式”来概括：闭包 = 匿名函数 + 外部环境变量

/*
里面用了外面的：var n int = 10
问题来了：AddUpper()执行结束后：
		AddUpper()结束，按理说 n 应该消失，但是返回出去的函数还需要 n
所以 Go 发现：这个变量还被外部函数引用，不能释放。是 Go 会把：变量 n + 匿名函数 打包保存。
这就是：closure = 函数 + 它引用的外部变量
*/

// TestClosure 函数要实现的是可以接收一个文件后缀名，并返回一个闭包，调用闭包可以传入一个文件名，如该文件有后缀则直接返回，若没有则加上
func TestClosure(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

// TestClosure2 这个和上面本质是一致的
func TestClosure2() func(string) string {
	suffix := ".jpg" // 这里只是显得通用性没那么强
	return func(name string) string {
		if strings.HasSuffix(name, suffix) {
			return name
		}
		return name + suffix
	}
}

// FIXME : 闭包的灵魂就是绑定外部变量，他一定会绑架一个外部变量否则就不是闭包。
// TODO : 闭包的作用在于

// NewCounter 优雅的计数器，状态 i 被安全地隔离在闭包内部
func NewCounter() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// MakeDiscount 再如可以传不同折扣率的闭包函数
func MakeDiscount(rate float64) func(price float64) float64 {
	return func(price float64) float64 {
		return price * rate
	}
}
