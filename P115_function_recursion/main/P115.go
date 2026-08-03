package main

import (
	"fmt"
	"learn/P115_function_recursion"
)

// 函数递归：在函数内部调用函数本身
func main() {
	P115_function_recursion.Test1(4)
	fmt.Println()
	P115_function_recursion.Test2(4)
	fmt.Println()
	fmt.Println(P115_function_recursion.Test3(7))
	// 1 1 2 3 5 8 13 21
	fmt.Println()
	fmt.Println(P115_function_recursion.Test4(3))
	fmt.Println(P115_function_recursion.Test5(10))

	a := P115_function_recursion.Sum2
	// b := P115_function_recursion.Sum3
	fmt.Println(P115_function_recursion.MyFunc(a, 1, 2)) // a 是没问题的
	// fmt.Println(P115_function_recursion.Myfunc(b, 1, 1)) // b 并不是我们定义的 type 的两个输入的类型的函数，所以这里不支持

	c := 13
	d := 24
	P115_function_recursion.Swap(&c, &d)
	fmt.Printf("c:%d, d:%d\n", c, d) // 交换成功的

}
