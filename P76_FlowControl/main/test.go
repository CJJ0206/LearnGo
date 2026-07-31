package main

import (
	"fmt"
	"math"
)

func test1() {
	var x int = 4
	var y int = 1

	// 无论多复杂，双分支永远只执行其中的一支，永远只有一直要看
	if x > 2 {
		if y > 2 {
			fmt.Println(x + y)
		}
		fmt.Println("cjj ")

	} else {
		fmt.Println(x)
	}
}

func test2(a int32, b int32) {
	if a+b > 50 {
		fmt.Println("over over over ")
	}
}

func test3(a float32, b float32) {
	if a > 10.0 && b < 20.0 {
		fmt.Printf("两数和为：%f \n", a+b)
	}
}

func test4(a int32, b int32) {
	sum := a + b
	if sum%5 == 0 && sum%3 == 0 {
		fmt.Println("该两数的和能够被3、5整除 ")
	}
}

func test5(a int32) {
	if (a%4 == 0 && a%100 != 0) || a%400 == 0 {
		fmt.Printf("%d年是闰年\n", a)
	} else {
		fmt.Printf("%d年不是闰年\n", a)
	}
}

func test6(a float64, b float64, c float64) {
	delta := b*b - 4*a*c
	if delta > 0 {
		ans1 := (-b + math.Sqrt(delta)) / (2 * a)
		ans2 := (-b - math.Sqrt(delta)) / (2 * a)
		fmt.Printf("此方程有两个根。分别为：%v 和 %v\n", ans1, ans2)
	} else if delta == 0 {
		ans1 := (-b + math.Sqrt(delta)) / 2 * a
		fmt.Printf("此方程有一个根。分别为：%v \n", ans1)
	} else {
		fmt.Printf("此方程有无解。")
	}
}
