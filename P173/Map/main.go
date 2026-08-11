package main

// 8.10

import "fmt"

/*
map 是 key-value 数据结构（键值对）,还称为映射

声明：
var map name map[keytype]valuetype
	var a map[string]string
	var a map[string]int
	var a map[int]string
	var a map[string]map[string]string
	// todo 声明是不会分配内存的，初始化需要make后才能赋值和使用

fixme 这个 key 的数据类型通常是 int / string
value 通常是 int float string struct map(嵌套)

*/

func main() {
	var a map[string]string
	fmt.Println(a) // 此时输出是 map【】空的
	//a["no1"] = "cjj"
	// 而且此时对他赋值，是会报错的，因为压根没有空间
	a = make(map[string]string, 10) // fixme 一定要make,这里就是创建10个大小的空间给这个map
	// make 不指定个数的话，容量就为1
	a["no1"] = "宋江"
	a["no2"] = "吴用"
	a["no1"] = "武松" // 1 会被覆盖，所以map里key是不重复的
	a["no3"] = "吴用" // value 是可以重复的
	fmt.Println(a)  // map[no1:宋江 no2:吴用]
	// todo 但是我们发现golang里的map是无序的

	test1()

	iteration()
	complexIteration()
	monster()

	MySort()

	mapt := map[string]map[string]string{
		"张三": {
			"nickname": "god",
			"code":     "4535245",
		},
		"李四": {
			"nickname": "cat",
			"code":     "8755124",
		},
	}
	modifyUser(mapt, "cjj")
}
