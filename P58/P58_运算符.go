package main

import "fmt"

/*
算术运算符
赋值运算符
比较运算符
逻辑运算符
位运算符
其他 & *
*/

func main() {
	// 算数运算符 + - * / % ++ -- +

	//除法
	fmt.Println(10 / 4)   // 结果是2：两头运算的都是整数，强制做的是整数运算
	fmt.Println(10 / 4.0) // 编译器发现浮点数，自动把10转为浮点防止精度损失

	var n1 float32 = 10 / 4 // 这里是因为传的是整数运算后的结果所以输出是2但是这个2的类型是float
	fmt.Printf("输出的类型是%T，数据值为%v\n", n1, n1)

	// 取余
	// 看一个公式 a % b = a - a / b * b 是否满足就可取余
	fmt.Println(10 % 3)    // 1
	fmt.Println(-10 % 3)   // -1
	fmt.Println(-10 % -3)  // -1
	fmt.Println(-10 % -10) // 0

	var i int = 1
	i++ // 自增、自减的操作是不可以赋值的，只能独立使用
	// 且没有前 ++ / --
	fmt.Println(i)

	// go里的 ++ 其实是一个语句

	//var i int = 10
	//i = i ++  // 不允许出现在语句
	//fmt.Println(i)
	//
	//var i2 int = 10
	//if i++ > 10{  // 比较也不可以
	//	fmt.Println("ok")
	//}

}
