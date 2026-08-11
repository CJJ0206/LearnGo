package main

// 8.10

import "fmt"

/*
map 的增删改查操作
map["key"] = value // 如果key还没有，就是增加，如果有了，就是修改
delete(map,"key"),delete是一个内置函数，如果key存在，就删除该键值对，如果不存在也不会报错
	go 里是没有办法一次性删除所有的key的，除非遍历删除
	或者直接把现有 map = make(...),用空的覆盖掉就行
todo map查找
	val,ok := map2("no1")  // 如果存在，会返回一个 值 和 true

map 的遍历，因为普通的for循环是用一个int的i去遍历，但是map可能两个值都不是int，也没有索引，所以要用range
*/

func iteration() {
	a := make(map[string]string)
	a["no1"] = "宋江"
	a["no2"] = "吴用"
	a["no1"] = "武松"
	a["no3"] = "吴用"

	for k, v := range a {
		fmt.Println(k, v)
	}
	for i := range a { // fixme range 出来的 i 是key，不是什么索引
		// fmt.Println(i)
		fmt.Println(a[i]) // 直接只要一个索引也是可以的
	}
}

func complexIteration() {
	studentMap := make(map[string]map[string]string)
	studentMap["no1"] = make(map[string]string, 3)
	studentMap["no1"]["name"] = "tom"
	studentMap["no1"]["sex"] = "man"
	studentMap["no1"]["address"] = "beijing"

	studentMap["no2"] = make(map[string]string)
	studentMap["no2"]["name"] = "mary"
	studentMap["no2"]["sex"] = "woman"
	studentMap["no2"]["address"] = "nanjing"

	// 那这个怎么遍历呢
	for _, v := range studentMap {
		for k1, v1 := range v { // 用index或者value都一样
			fmt.Println(k1, v1)
		}
	}
	fmt.Println(len(studentMap)) // todo map 的长度同样是用len获取
}
