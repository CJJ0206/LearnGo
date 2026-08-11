package main

// 8.10

import "fmt"

/*
map切片
那其实切出来就是动态map呗和数组一样  【】map
*/

// 案例：使用一个map来记录monster的信息，name和sex，妖怪数可以动态增长
func monster() {
	var slice []map[string]string
	slice = make([]map[string]string, 2) // fixme 这里是单独对切片空间进行make,没有这一步后面都是错的

	slice[0] = make(map[string]string) // fixme 这些都是对map进行make，可以不用带size
	slice[0]["name"] = "牛魔王"
	slice[0]["age"] = "458"

	slice[1] = make(map[string]string)
	slice[1]["name"] = "玉兔"
	slice[1]["age"] = "778"

	// 由于上面的容量为2，所以这里报错了
	// fixme 切片的扩容机制是在使用append的前提下的
	monster3 := map[string]string{
		"name": "猴子",
		"age":  "458",
	}
	slice = append(slice, monster3) // 通过先创建一个map元素，然后直接append到slice里

	fmt.Println(slice)

}
