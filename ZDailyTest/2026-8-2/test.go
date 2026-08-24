package main

import "fmt"

func main() {
	test(4)
	fmt.Println(fbn(5))
	fmt.Println(monkey(10))
}

// 这个测试要注意
func test(n int) {
	if n > 2 {
		n--
		test(n)
	}
	fmt.Println(n)
}

func fbn(n int) int {
	if n <= 2 {
		return 1
	}
	return fbn(n-1) + fbn(n-2)
}

// fixme 主要就是得到公式，然后这题已知的停止参数是第十天，所以可以反过来计算
// d1/2 -1 = d2 ; d1 = 2(d2 + 1 )
func monkey(i int) int {
	if i == 1 {
		return 1
	}
	return 2 * (monkey(i-1) + 1)
}

/*
递归调用时，Go 的调用栈是怎样的？"先入后出"是什么意思？
答案
每次递归调用会在栈顶压入一个新的函数帧，挂起当前层。最深层先执行完毕出栈，然后逐层返回。这就是"先入后出"——最先调用的函数最后执行完。
类比：叠盘子，最后放上去的最先拿下来。
*/
