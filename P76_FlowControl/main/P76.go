package main // 7.31

import "fmt"

// 要import别的包的函数，只能引用整体的包，不能单独引用文件
// 同一个文件夹下的内容本身就是一个包里的，所以同一个文件夹里的函数是可以直接调用的

func main() {

	testCont()

	var age int
	fmt.Print("请输入年龄：")
	_, err := fmt.Scanln(&age) // Scanf接收的是一个变量的地址
	if err != nil {
		fmt.Println("格式输入有误，请重新出入。")
		return
	}
	isAdult(age)

	// isAdult2(age)

	test1()

	var a, b int32 = 33, 29
	test2(a, b)

	var c, d float32 = 11.0, 3.0
	test3(c, d)

	var e, f int32 = 7, 8
	test4(e, f)

	var year int32 = 2020
	test5(year)

	var score float32
	fmt.Print("请输入成绩；")
	_, err1 := fmt.Scanln(&score)
	if err1 != nil {
		fmt.Println("输入格式有误，请重新输入。")
		return
	}
	reward(score)

	var a2, b2, c2 float64 = 2.0, 3.0, 1.0
	test6(a2, b2, c2)

	var racescore float32
	var gender string
	fmt.Print("请输入成绩和性别:")
	_, err2 := fmt.Scanln(&racescore, &gender)
	if err2 != nil {
		fmt.Println("输入格式错误，请重新输入。")
	}
	loopTest1(racescore, gender)

	var cin byte
	fmt.Print("请输入内容")
	_, err3 := fmt.Scanf("%c\n", &cin)
	if err3 != nil {
		fmt.Println("出入格式有误，重新输入。")
	}
	testSwitch(cin)

	fallt()
	fmt.Println()

	tower()
	fmt.Println()

	table()
	fmt.Println()

	num()
	fmt.Println()

	login()

}

// Scanf 函数扫描从标准输入读取的文本，并将连续的空格分隔值存储到指定的参数中
