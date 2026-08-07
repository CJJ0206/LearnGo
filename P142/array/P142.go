package main

import "fmt"

/*
引入：
一个养鸡场有6只鸡，体重分别为3kg,5kg,1kg,3.4kg,2kg,50kg
请问这六只鸡的总体中是多少？平均体重是多少？


数组可以存放做个同一类型的数据，数组也是一种数据类型，是值类型的
*/

var chic [6]float64 // 一个float64的数组
// 全局作用域中只允许声明变量，不允许执行语句
// chic[0] = 3.0  // 所以这里会报错的

func main() {
	chicken()
	TestArray()
	TestInput()
	test()
	arr := [3]int{1, 2, 3} // Fixme go里面认为长度也是类型的一部分，所以长度为4的数组和长度为3的数组不是一个类型
	test2(arr)             // go 会把arr复印一份到函数的栈里去，这种操作其实是很浪费资源的，这也会引出后面的切片
	fmt.Println(arr)       // 可以看到函数里的修改对外部是没有用的，所以一定是值传递

	test4(&arr)
	fmt.Println(arr)

	ab()
	fmt.Println()

	arr2 := [5]float64{0.1, 0.2, 0.3, 0.4, 0.5}
	sum, avg := avgArr(arr2)
	fmt.Println(sum, avg)

	divert()
}

func chicken() {
	chic[1] = 5.0
	chic[2] = 1.0
	chic[3] = 3.4
	chic[4] = 2.0
	chic[5] = 50.0

	var sum float64
	for _, v := range chic {
		sum += v
	}
	fmt.Println(sum)
	fmt.Printf("%.2f\n", sum/6)                  // 普通的没有指定类型的数，go会当作无类型常量
	fmt.Printf("%.2f\n", sum/float64(len(chic))) // 指定为两位小数  // 但 len 返回值一定是int，需要转一下
	// 对于处理大量数据时非常方便
}
