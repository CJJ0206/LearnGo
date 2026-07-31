package main

// test + 关系运算符（比较运算符） + 逻辑运算符 + 赋值运算符

import "fmt"

/*
假如还有97天放假，问：xx个星期零x天
定义一个变量保存华氏温度，转换公式为：5/9*（华氏温度-32），请求出华氏温度对应的摄氏温度
*/

func main() {
	// test1
	var restDay int32 = 97
	var week int32
	var day int32
	week = restDay / 7
	day = restDay % 7
	fmt.Printf("剩余%d周%d天\n", week, day)

	// test2
	var huaShi float32 = 99.25
	var sheShi float32
	sheShi = 5.0 / 9 * (huaShi - 32) // 5要用5.0 因为顺序执行的原因 5/9 会得到结果 0
	fmt.Println(sheShi)

	// 关系运算符 == != < > <= >=
	var n1 int = 1
	var n2 int = 2
	fmt.Println(n1 == n2)
	fmt.Println(n1 != n2)
	fmt.Println(n1 >= n2)
	fmt.Println(n1 <= n2)

	// 逻辑运算符
	/*
		&& 两边均为真才为true
		|| 有一个true就为true
		! 逻辑非，取反
	*/
	// 案例演示
	var age int = 32
	if age < 30 && age > 10 {
		fmt.Println("青年")
	}
	if age > 30 {
		fmt.Println("中年")
	}

	// 逻辑与的两个条件均为真才返回true，当第一个条件为false是就直接输出了，不判断后面了
	// 逻辑或的两个条件均为假才返回false,当第一个条件为true是也直接输出，不判断后面
	var i int = 10
	if i < 9 && test() { // 从输出是可以直接看到没有调用test，直接跳过了
		fmt.Println("ok")
	}

	// 赋值运算符 + += -= *= /= %=
	var a = 10
	var b = 20
	a += 10
	b *= 5
	fmt.Println(a, b)

	//test3--------------------------------------------------------------------
	// 有两个变量，a、b，要将其进行交换，但是不允许使用中间变量
	var a1 int = 10
	var b1 int = 20
	a1 = a1 + b1
	b1 = a1 - b1
	a1 -= b1
	fmt.Println(a1, b1)
	//--------------------------------------------------------------------------

	// 单目运算个赋值运算都是从右向左的
	var a3, b3, c int
	fmt.Println(a3, b3, c)

	// 取地址和解引用 & *

	// test---------------------------------------------------------------------
	// 计算两个数的最大值已经三个数的最大值
	var a4 int = 10
	var b4 int = 20
	var c1 int = 30

	ma := max2(a4, b4)
	fmt.Printf("最大值为：%d\n", ma)
	ma2 := max3(a4, b4, c1)
	fmt.Printf("最大值为：%d\n", ma2)
	//--------------------------------------------------------------------------

}

func test() bool {
	fmt.Println("test")
	return true
}

func max2(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func max3(a int, b int, c int) int {
	max := a // 假设 a 最大
	if b > max {
		max = b // 如果 b 更大，替换 max
	}
	if c > max {
		max = c // 如果 c 更大，替换 max
	}
	return max // 唯一的出口，编译器绝对放心
}
