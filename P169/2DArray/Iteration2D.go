package main

// 8.10

import "fmt"

// 1.双层for循环遍历
// 2.for range 遍历
// 二维数组的len拿到的是第一个长度

var arr = [3][4]int{{1, 2, 3, 4}, {2, 3, 4, 5}, {5, 6, 7, 8}}

func iteration1() {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Print(" ", arr[i][j])
		}
	}
	fmt.Println()
}

// 这两个是完全等价的
// todo 这个是只取索引的方式
func iteration2() {
	for i := range arr {
		for j := range arr[i] {
			fmt.Print(" ", arr[i][j])
		}
	}
	fmt.Println()
}

// todo 这个是只取值的方式
func iteration3() {
	for _, v := range arr { // 因为这个v是一维数组，所以可以继续range
		for _, v2 := range v {
			fmt.Print(" ", v2)
		}
	}
}
