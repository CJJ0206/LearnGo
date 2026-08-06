package main

import "strconv"

// 实践案例：编写一段代码来统计函数test的执行时间

func test() {
	str := ""
	// go 更新后 对于range的使用如果为int则只返回一个值了
	for value := range 100000 {
		str += "hello" + strconv.Itoa(value) // 把int转string
	}
}
