package main

/*
func scanln:读取一行数据
func Scanf
func Fscanf
func Sscanf
func Scan
func Fscan

案例：要求可以从控制台接收用户信息[姓名、年龄、薪水、是否通过考试]
*/

import "fmt"

func main() {
	// 方式一：fmt.scanln-----------------------------------------------------------------
	var name string
	var age int
	var salary int
	var isPassed bool
	//fmt.Println("请输入姓名")
	//fmt.Scanln(&name) // 接收的是地址，因为只有地址里的真是值改变了，值才会改变
	////当程序执行到这里是会听在这里等待输入
	//
	//fmt.Println("请输入年龄")
	//fmt.Scanln(&age) // 标黄是因为这个函数有两个返回值，编译器在提醒我们是否遗漏了err
	//
	//fmt.Println("请输入薪水")
	//fmt.Scanln(&salary)
	//
	//fmt.Println("请输入是否通过")
	//fmt.Scanln(&isPassed)
	//
	//fmt.Printf("姓名%s年龄%d薪水%d是否通过%t", name, age, salary, isPassed)

	// 方式二：Scanf-----------------------------------------------------------------------
	fmt.Println("请输入姓名、年龄、薪水、是否通过考试，用空格隔开")
	fmt.Scanf("%s %d %d %v", &name, &age, &salary, &isPassed)
	fmt.Printf("姓名%s年龄%d薪水%d是否通过%t", name, age, salary, isPassed)

}
