package main

// 8.11

import "fmt"

/*
Go 在设计上抛弃了传统的 class 的概念，它的面向对象体系是非常轻量和独特的
在 Go 里，struct 承担了传统面向对象语言中 class 的核心数据承载功能

张老太太养了一只猫，名字叫小白，今年三岁，白色。还有一只叫小花，今年20岁，花色。请编写一个程序
当用户输入小猫的名字时，就显示该猫的名字、年龄、毛色。如果用户输入的名字错误，则显示张老太太没有这只猫

todo 我的第一个想法是用mpa套map的结构来表示 map[string]map[string]string
使用单独变量解决（麻烦）
使用数组解决（二维数组吗：数组是做不了的，二维不够，三个一维数组又很麻烦）
*/

func main() {
	// mapMethod1()
	catStruct()
	createStruct2()
	createStruct3()
	createStruct4()

	test()

	detail1()
	detail2()
	detail3()
}

// 方式一：使用map嵌套(todo Details)
// 细致看一下，有助于理解很多细节
func mapMethod1() {
	var inputName string
	catMap := map[string]map[string]string{
		"小白": {"age": "3", "color": "白色"},
		"小花": {"age": "20", "color": "花色"}}
	fmt.Print("请输入猫名：")
	_, err := fmt.Scanln(&inputName)
	if err != nil {
		fmt.Println(err)
	}

	// todo 记到DailyTest里
	// cat, exists := catMap[inputName] 这行语句在做的是
	//用catMap[inputName]里的值去声明创建cat, exists，如果没有，那自然就是false
	if cat, exists := catMap[inputName]; exists { // 这里的句式是 if := ; bool判别式/bool{}
		fmt.Printf("找到啦！名字: %s, 年龄: %s岁, 毛色: %s\n", inputName, cat["age"], cat["color"])
	} else {
		fmt.Println("张老太太没有这只猫")
	}
}

// pass
func method2() {
	//var vat1name string = "小白"
	//var ...
	// 显然光在定义阶段就很麻烦了
}

// 方式三：定义二维切片 (类似二维数组)
func method3() {
	// 每一行的格式约定为: [名字, 年龄, 毛色] 只有约定好，才能取出相对应的值，不如map有key
	cats := [][]string{
		{"小白", "3", "白色"},
		{"小花", "20", "花色"},
	}

	var inputName string
	fmt.Print("请输入小猫的名字: ")
	_, err := fmt.Scanln(&inputName)
	if err != nil {
		fmt.Println(err)
	}

	isFound := false
	// 遍历二维切片查找
	for _, cat := range cats {
		if cat[0] == inputName { // cat[0] 就是名字
			fmt.Printf("找到啦！名字: %s, 年龄: %s岁, 毛色: %s\n", cat[0], cat[1], cat[2])
			isFound = true
			break // 找到了就立刻停止循环
		}
	}
	// 如果循环结束还没找到
	if !isFound {
		fmt.Println("张老太太没有这只猫")
	}
}
