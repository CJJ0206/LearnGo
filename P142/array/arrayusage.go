package main

import (
	"fmt"
	"math/rand/v2" // v2不用指定随机种子了
)

// 创建一个byte类型的2个元素的数组，分别放置A-Z,使用 for 循环访问所有元素并打印出来
// 求数组的最大值以及下标
// 求数组的和与平均值 range

func ab() {
	var arr [26]byte
	for i := range 26 {
		arr[i] = byte(i + 'A')
	}
	fmt.Printf("%c", arr)
}

func maxArr(arr [5]int) {
	// 那就是一个个传统遍历了
}

func avgArr(arr [5]float64) (float64, float64) {
	var sum float64
	for _, v := range arr {
		sum += float64(v)
	}
	avg := sum / float64(len(arr))
	return sum, avg
}

// 随机生成5个数，并将其反转打印
func divert() {
	var arr [5]int
	for i := range arr {
		arr[i] = rand.IntN(100)
	}
	fmt.Println(arr)
	for i := len(arr) - 1; i >= 0; i-- {
		fmt.Print(arr[i], " ")
	}
	// 或者用临时变量来做左右交换

}
