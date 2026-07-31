package main

import "fmt"

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
