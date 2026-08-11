package main

import (
	"fmt"
	"sort"
)

// MySort map 是无序的，那怎么对map进行排序呢
// fixme 利用切片对他的key排序后输出
func MySort() {
	map1 := make(map[int]int, 10)
	map1[0] = 10
	map1[1] = 16
	map1[8] = 17
	map1[3] = 14

	// 我们先把map的key放进切片中
	// 对切片排序
	// 遍历切片，然后按照key来输出map的值
	var keys []int
	for k := range map1 {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	fmt.Println("keys:", keys)
	// 然后用这个key顺序遍历map就行了
}
