package main

// 8.10
// 多维数组我们只介绍二维数组
/*
二维数组使用和内存分布
使用方式：先申明/定义，3再赋值
1. var name [][]type
	比如：var a2 [3][3]int
2. 内存布局
	(意义：两个包含三个元素的一维数组)
	arr [2][3]  -->  arr[*ptr1 | *ptr2]  // todo 存的是指向两个一维数组的地址指针
				     |2|3|4|      |5|6|7|  --> 这两个指针分别指向对应的一维数组
3. 同样也是支持直接初始化的
	var arr2 = [2][2]int{{1,2},{2,3}}
4. 另外几种方式
	var arr2 = [size][size]int{{},{},{},...}
	var arr2 = [...][size]int{{},{},{}...}  // 不指定大小，根据赋值自动计算大小

	arr2 := [size][size]int{{},{},{},...}
	...
*/

import (
	"fmt"
)

func main() {
	// 输出一个矩形数字矩阵
	// 先声明一个二维数组
	var arr [4][6]int // todo 意思就是4个数组他们的元素是有6个元素的一维数组
	arr[2][4] = 25
	fmt.Println(arr[2])
	for i := range arr { // 对于数组，如果只有一个值的话那返回的就是索引
		fmt.Println(arr[i])
		// fixme arr[i]这样获取的是一个一维数组里面含6个元素，所以要换行输出直接一行行换行就行了
	}

	// --------------------- 内存测试 ----------------------
	var arr2 [2][3]int
	fmt.Printf("arr2[0]的地址是%p \n", &arr2[0])
	fmt.Printf("arr2[1]的地址是%p \n ", &arr2[1]) // 这里的输出是刚好相差24个字节，也就是一个一维数组的3个元素

	fmt.Printf("arr2[0][0]的地址是%p \n", &arr2[0][0]) // 确实这个首元素的地址就直接等于去一维数组的地址
	fmt.Printf("arr2[1][0]的地址是%p \n ", &arr2[1][0])

	var arr3 = [2][3]int{{1, 2, 3}, {2, 3, 4}}
	// 等价于 var arr3 = [...][3]int{{1, 2, 3}, {2, 3, 4}}
	fmt.Println(arr3)

	iteration1()
	iteration2()
	iteration3()

	// classScore()
	// classScore2()

	test2()
	test3()
	test5()

}
