package main // 7.31

import (
	"fmt"
	"math"
)

// 单分支
func isAdult(age int) {
	if age > 18 {
		fmt.Println("你是成年人，要对自己的行为负责！")
	}
}

// 双分支
func isAdult2(age int) {
	if age > 18 {
		fmt.Println("你是成年人，要对自己的行为负责！")
	} else {
		fmt.Println("你是未成年，回去吃奶")
	}
}

// golang支持在if中，直接定义变量
//if age := 20; age > 18 {un
//	fmt.Println("成年了")
//}

// 多分枝
func reward(score float32) {
	if score == 100 {
		fmt.Println("奖励一台车")
	} else if score > 80 && score < 90 {
		fmt.Println("奖励一部手机")
	} else if score > 60 && score < 80 {
		fmt.Println("奖励一个手表")
	} else {
		fmt.Println("nothing")
	}
}

// 同样只会执行一个入口，一旦某个分支被执行了，就不会继续往下了

//func wrong(){
//	var b bool = true  // go里面 if 语句后面是不能跟赋值语句的，只能是条件表达式
//	if b = true{
//	}
//}

//---------------------------------------------------------------------------
/*
switch分支结构（用于根据不同条件执行不同命令）
go的switch的语句后面是不用加break的

switch的执行流程是，先执行表达式，得到值，然后和case表达式进行比较，如果相等就执行该case的下的操作，然后会自动break
如果所有case都没有匹配，执行最后的default块的内容退出
go的一个case中的表达式可以有多个，用逗号隔开
*/

func testSwitch(c byte) {
	switch c {
	case 'a':
		fmt.Println("周一")
	case 'b':
		fmt.Println("周二")
	case 'c':
		fmt.Println("周三")
	case 'd':
		fmt.Println("周四")
	default:
		fmt.Println("输入无效")
	}
}

func os(s string) {
	switch os := "windows"; os { // 允许在语句中直接申明定义变量
	case "macOS":
		fmt.Println("Mac")
	case "windows":
		fmt.Println("Windows")
	default:
		fmt.Println("Other OS")
	}
}

func fallt() {
	var i int = 10
	switch i {
	case 10:
		fmt.Println("10")
		fallthrough // switch穿透，只穿透下一层直接执行输出
	case 20:
		fmt.Println("穿透一层")
	case 30:
		fmt.Println("穿透两层")
	}
}

/* switch 细节
case 后可以是一个表达式
case 后可以带多个表达式，用逗号隔开
case 甚至可以跟一个又返回值的函数，或者常量（只要有值就行）
switch 和 case 的值的数据类型要一致，否则报错
switch 后可以不带表达式，直接当做if else 语句用
Type switch:switch语句还可以用于 type-switch 来判断某个 interface（接口） 变量中实际指向的变量类型




switch n1{
	case n2 , 10 , 5:  // 是允许这样做的去和多个选项比较
		fmt.Print()
}



switch {  // 这种switch后面不加表达式就相当于是一个if分支语句
	case age == 10 :
		fmt.Pringt（"为成年"）
	case age == 20 :
		fmt.Print("成年了")
}



switch {  // case后面的语句也可以是一些判断表达式
	case age < 10 :
		fmt.Pringt（"为成年"）
	case age > 20 :
		fmt.Print("成年了")
}



func inter(i int, x interface{}) {
	switch i := x.(type) { // i 会接收 x 的数据类型
	case nil:
		fmt.Printf("x的类型是%T")
	case int:
		fmt.Println("x的类型是int")
	case float64:
		fmt.Println("类型是float")
	}
}
*/

// test ----------------------------------------------------------------------------
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
