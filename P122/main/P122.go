package main

// 8.3 8.4

import (
	"fmt"
	"learn/P122/utils"
)

// 所以整体的执行顺序是 全局变量先初始化 -> init -> main

var age = test()

func test() int {
	fmt.Println("test function()...")
	return 23
}

// init 主要完成一些初始化的工作
func init() {
	fmt.Println("init function()...")
}

// Fun3 把匿名函数绑定到全局变量，使得全局可用
var (
	Fun3 = func(n1 int, n2 int) int {
		return n1 * n2
	}
)

func main() {
	fmt.Println("main function()...")
	fmt.Printf("年龄为%d,姓名为%s \n", utils.Age, utils.Name)

	// 匿名函数 test ==========================================================

	// 求两个数的和，用匿名函数完成
	res1 := func(n1 int, n2 int) int { // 我们发现这个函数是没有函数名的，所以他只能立即执行不能跨包执行
		return n1 + n2
	}(10, 20)
	fmt.Println(res1)

	// 匿名函数使用方法二
	a := func(n1 int, n2 int) int { // 把这个匿名函数对象给a
		return n1 - n2
	}
	res2 := a(10, 20) // 就可以直接通过变量a去调用函数
	fmt.Println(res2)

	res4 := Fun3(4, 6)
	fmt.Println(res4)

	// closure test ===========================================================
	// 测试闭包
	f := utils.AddUpper() // 注意是有括号的！！这里相当于 a 接收的是AddUpper()运行后返回的结果匿名函数
	// 所以闭包的外层函数是不需要形参的

	// 多次调用这个闭包函数
	fmt.Println(f(1)) // 输出: 11 (10 + 1)
	fmt.Println(f(2)) // 输出: 13 (11 + 2)
	fmt.Println(f(3)) // 输出: 16 (13 + 3)

	// 下面是正常的案例
	t := utils.TestClosure(".jpg") // 可以看到这里只需要一次指定就可以一直判断，也是个好处
	fmt.Println(t("english"))      // english.jpg
	fmt.Println(t("china.jpg"))    // china.jpg
	fmt.Println(t("japan"))        // japan.jpg

	t2 := utils.TestClosure2()
	fmt.Println(t2("english"))   // english.jpg
	fmt.Println(t2("china.jpg")) // china.jpg
	fmt.Println(t2("japan"))     // japan.jpg

	// defer test ==============================================================
	utils.Sum(10, 20)
	utils.Sum2(10, 20)

	// 变量作用域 ================================================================
	fmt.Println(utils.Name2) // 属于同一个包就不需要带前缀，但是这里不在同一个包里
	utils.Test01()
	utils.Test02()

	fmt.Println()

	fmt.Println(utils.Name2)
	utils.Test04()
	utils.Test03()

	utils.TestStr()
	fmt.Println()
	utils.TestStr2()

	utils.TestAtoi()

	utils.ConvByte()
	utils.ConvStr()

	utils.Contain()
	utils.Contain2()
	utils.Equal()
	utils.JudgeExist()
	utils.LsatSee()
	utils.Replace()
	utils.Split()
	utils.UpLow()
	utils.Clear()

}
