package main // 7.31

/* for 循环基本语法
for 循环变量初始化；循环条件；循环变量迭代{
	循环操作语句
}
并且这个变量的作用域仅在循环内
*/

import "fmt"

// 循环控制就是要让代码可以循环执行
// 嵌套分支不要过深，深了就很复杂了，建议控制在3层

func loopTest1(score float32, gender string) {
	if score < 8.0 {
		if gender == "man" {
			fmt.Println("进入男子组决赛。")
		} else if gender == "woman" {
			fmt.Println("进入女子组决赛。")
		}
	}
}

// ---------------------------------------------------------------------------------
// for 循环的第一种写法
func fo1() {
	for i := 1; i <= 10; i++ { // i = 1 ； i <= 10 ; i自增
		fmt.Println("你好")
	}
}

// for 循环的第二种写法
func fo2() {
	j := 1
	for j <= 10 {
		fmt.Println(j)
		j++
	}
}

// for 循环第三种写法(死循环)，通常配合 break 使用
func fo3(k int) { // 配合达到特定条件再终止
	for {
		if k < 10 {
			fmt.Println("1")
		} else {
			break
		}
	}
}

// for-range
// golang 提供for-range方法用来方便的遍历字符串和数组
func oldFor() {
	var str string = "hello"
	for i := 0; i < len(str); i++ {
		fmt.Printf("第%d个字符是%c", i, str[i]) // 传统方式访问字符串每个char
	}
}

// range 是普遍使用的，可以获取各种数据结构的索引和对应值
func newFor(str string) {
	for index, value := range str {
		fmt.Printf("index = % d,value = %v", index, value)
	}
}

//------------------------------------------------------------------------------------

// 传统 for 会把中文截断
func forTest() {
	str := "hello 上海" // range是按rune字符遍历的，value是完整的rune(int32)能完美还原中文
	for index, value := range str {
		fmt.Printf("index = %c,value = %v", index, value)
	}
}

func LoopTest2() {
	var sum int = 0
	for i := 1; i <= 100; i++ {
		if i%9 == 0 {
			sum += i
		}
	}
	fmt.Printf("结果为：%d", sum)
}

// 把 for 训话的步长直接设为9，直接做求和
func betterLoopTest2() {
	var sum int = 0
	// 直接从 9 开始，每次累加 9，出来的全是 9 的倍数
	for i := 9; i <= 100; i += 9 {
		sum += i
	}
	fmt.Printf("结果为：%d", sum)
}

func LoopTest3(res int) {
	for i := 1; i <= res; i++ {
		fmt.Printf("%d + %d = %d \n ", i, res-i, res)
	}
}

// ------------------------------------------------------------------------------------
// go 语言是没有 while 和 do while 的 ， 但可以通过 for 来实现对应效果
// while 先判断再执行
// do while 先执行再判断（必定会执行一次）

// ForWhile 相当于 while (i <= 10) 逻辑
func ForWhile() {
	var i int = 1
	for {
		if i > 10 {
			break
		} else {
			fmt.Println(i)
			i++
		}
	}
}

// ForWhile2 是简洁写法
func ForWhile2() {
	var i int = 1
	// 条件写在 for 后面，当 i <= 10 时继续循环，超出则自动退出
	for i <= 10 {
		fmt.Println(i)
		i++
	}
}

// DoWhileTest 相当于 do while 逻辑
func DoWhileTest() {
	var i int = 1
	for {
		// 1. 先无条件执行逻辑
		fmt.Println(i)
		i++

		// 2. 后判断条件：如果满足退出条件则 break
		if i > 10 {
			break
		}
	}
}

// 开头大写的函数的注释也要以其名字开始

// --------------------------------------------------------------------------------------
// 多重嵌套循环控制（不要超过两层）
// 其实就是把内层循环当作外层循环的循环体
// 这个得多做两编

func score() {
	var totalSum float64 = 0 // 用于记录所有班级的总成绩
	classCount := 3          // 班级数量
	studentCount := 5        // 每个班的学生数量

	// 外层循环：控制班级（从第 1 班到第 3 班）
	for i := 1; i <= classCount; i++ {
		var classSum float64 = 0 // 每次进入新班级，班级总分清零   // 把整个方法忘了
		fmt.Printf("--- 开始输入第 %d 个班级的成绩 ---\n", i)

		// 内层循环：控制学生（从第 1 个到第 5 个）
		for j := 1; j <= studentCount; j++ {
			var score float64
			fmt.Printf("请输入第 %d 个学生的成绩: ", j)
			_, err := fmt.Scanln(&score)
			if err != nil {
				fmt.Println("输入格式错误，请重新输入")
			}
			classSum += score // 累加当前班级的总分
		}
		// 计算并输出当前班级的平均分
		classAvg := classSum / float64(studentCount)
		fmt.Printf(">> 第 %d 班的平均成绩为：%.2f\n\n", i, classAvg)
		// 将当前班级的总分累加到年级总池子里
		totalSum += classSum
	}

	// 计算所有班级的总平均成绩
	totalStudents := classCount * studentCount
	totalAvg := totalSum / float64(totalStudents)

	fmt.Printf("=== 最终统计结果 ===\n")
	fmt.Printf("所有班级的总平均成绩为：%.2f\n", totalAvg)
}

/*
     *       1层 1*  2 * 层数 - 1  空格 2 规律 总层数减当前层数
    ***      2层 2*
   *****     3层 3*
*/
func tower() {
	var totalevel = 9
	for i := 1; i <= 9; i++ {
		for k := 1; k <= totalevel-i; k++ {
			fmt.Print(" ")
		}

		for j := 1; j <= 2*i-1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

// 进阶：空心金字塔

// 九九乘法表
func table() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d x %d = %d \t", j, i, i*j)
		}
		fmt.Println()
	}
}
