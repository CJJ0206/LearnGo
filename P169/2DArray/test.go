package main

// 8.10

import (
	"fmt"
	"math/rand/v2"
	"strings"
)

// 定义一个二维数组，用来保存三个班，每个班五名同学的成绩，并求出每个班的平均分，以及所有班的平均分
func classScore() {
	arr := [3][5]int{}
	var sum int
	for i, v := range arr {
		for j, v2 := range v {
			fmt.Printf("\n请输入%d班%d号的成绩:", i+1, j+1)
			_, err := fmt.Scanln(&v2) // fixme range的 &v2 是临时变量地址，无法修改原数组的值，所以这里只能用索引
			if err != nil {
				fmt.Println("输入有误")
			}
		}
	}
	fmt.Println(arr)
	for i, v := range arr {
		sumEvg := 0
		for _, v2 := range v {
			sumEvg += v2
			sum += v2
		}
		fmt.Printf("班级%d的平均分为：%d", i+1, sumEvg/len(arr[0]))
	}
	fmt.Printf("总平均分为：%d", sum/(len(arr[0])*len(arr)))
}

// 正确版本
func classScore2() {
	arr := [3][5]int{}
	var sum int

	// 输入成绩部分
	for i := range arr {
		for j := range arr[i] {
			fmt.Printf("\n请输入%d班%d号的成绩:", i+1, j+1)
			// 直接使用原数组的地址进行赋值
			_, err := fmt.Scanln(&arr[i][j])
			if err != nil {
				fmt.Println("输入有误")
			}
		}
	}

	fmt.Println("录入的成绩为:", arr)

	// 计算平均分部分
	for i, v := range arr {
		sumEvg := 0
		for _, v2 := range v {
			// 这里只是读取，不修改原数组，所以用 v2 这个拷贝是没有问题的
			sumEvg += v2
			sum += v2
		}
		// len(arr[0]) 或 len(v) 都是5
		fmt.Printf("班级%d的平均分为：%d\n", i+1, sumEvg/len(v))
	}
	fmt.Printf("总平均分为：%d\n", sum/(len(arr[0])*len(arr)))
}

// 1.随机生成10个整数（1~100）保存到数组，并倒序打印，以及求平均，最大，最小的下标，并查找表里有没有55

// 2.定义一个3*4的数组，逐个从键盘输入，编写程序将四周清零

// 3.4*4的,实现1 4行互换， 2 3 行互换

// 4.已知有个升序的数组，要求插入一个元素，最后打印该数组，依然是升序

// 5.写出；实现查找的核心代码，如已知数组arr[10]string 里面存了十个元素，查找AA是否在其中，打印提示，如果有多个打印其下标

// 6.函数接收数组，大小为5，找出最大最小对应的下标

// 7. 8 个整数的数组，求数组中大于和小于平均值的个数

/*
8.跳水比赛，8个评委评分，运动员成绩是8个中去掉一个最高分，去掉一个最低分，剩下的6个平均就是最后成绩
	1.把打最高和最低恶的评委找出来
	2.找出最佳和最差评委（与最终结果分差小/大）
*/

func test1() {
	var arr [10]int
	var sum int
	for i := range 10 {
		arr[i] = rand.IntN(100) + 1
		sum += arr[i]
	}
	fmt.Println(arr)

	for i := 9; i >= 0; i-- {
		fmt.Println(arr[i])
	}

}

func test2() {
	fmt.Println("Here is the test2")
	var arr [5][6]int
	for i := range arr {
		for j := range arr[i] {
			arr[i][j] = rand.IntN(100) + 1
		}
	}
	for i := range arr {
		fmt.Println(arr[i])
	}
	fmt.Println()

	for i := range arr {
		for j := range arr[i] {
			if i == 0 || i == len(arr)-1 || j == len(arr[i])-1 || j == 0 {
				arr[i][j] = 0
			}
		}
	}
	for i := range arr {
		fmt.Println(arr[i])
	}
}

func test3() {
	fmt.Println("Here is the test3")
	var arr [4][4]int
	for i := range arr {
		for j := range arr[i] {
			arr[i][j] = rand.IntN(100) + 1
		}
	}
	for i := range arr {
		fmt.Println(arr[i])
	}
	fmt.Println()

	var temp1 [4]int
	temp1 = arr[2]
	arr[2] = arr[0]
	arr[0] = temp1
	temp1 = arr[3]
	arr[3] = arr[1]
	arr[1] = temp1
	for i := range arr {
		fmt.Println(arr[i])
	}
}

func test4() {
	// todo 问题是要插入一个新值保持升序
	// 最大的问题是数组是不允许动态变化的，那么我能想到的就是 slice := arr[:]
	// 用切片 append 进去
	arr := [8]float64{1, 2, 3, 4, 5, 6, 7, 8}
	slice := arr[:]
	slice = append(slice, 2.5)
	// 排序插进去就行

}

func test5() {
	arr := [10]string{
		"AApple",
		"banana",
		"cherry",
		"dAAte",
		"elderberry",
		"fig",
		"grape",
		"honeydew",
		"kiwi",
		"lemon",
	}
	for i := range arr {
		bo := strings.Contains(arr[i], "AA")
		// index := strings.Index(arr[i], "AA")
		if bo == true {
			fmt.Printf("arr[%d]中存在AA,坐标是%d \n", i, i)
		}
	}

}

func test6() {
	// 找到最大或者最小
	// 用index获取对应数字的位置就行
}

func test7() {
	// 求均值然后比较
	// 两个count用来记录
}

func test8() {
	// 主要还是排序，找出最大最小
	// 单组直接计算，多组求和均值的方差即可
}
