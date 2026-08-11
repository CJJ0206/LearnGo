package main

// 8.10

import "fmt"

/*
map有以下几种使用方法
 	方式1：
		先声明一个map,之后再make（初始为nil）
	方式2：
		var map2 = make(map[string]string,10)  // 声明并直接make
	方式3：
		var map2 = map[string]string {"no1":"北京",..}  // 声明并直接赋值
*/

func use1() {
	var map1 map[string]string
	map1 = make(map[string]string, 10)
	map1["no1"] = "宋江"
}

func use2() {
	var map2 = make(map[string]string, 10)
	fmt.Println(map2)
}

func use3() {
	var map3 = map[string]string{
		"no1": "cjj",
		"no2": "qg",
	}
	fmt.Println(map3)
}

// 演示一个 k-v,v 是map的例子。三个学生，每个都有name and sex
func test1() {
	studentMap := make(map[string]map[string]string)
	studentMap["no1"] = make(map[string]string, 3) // 第二层也需要先make一下才行
	studentMap["no1"]["name"] = "tom"              // no1 的 name
	studentMap["no1"]["sex"] = "man"               // no1 的 sex
	studentMap["no1"]["address"] = "beijing"       // no1 的 address

	studentMap["no2"] = make(map[string]string) // fixme make 一定是不能少的 ,不加长度也可以，会自动跟着扩容, 但是只是map不需要加长度，切片是一定要加长度的
	studentMap["no2"]["name"] = "mary"          // no1 的 name
	studentMap["no2"]["sex"] = "woman"          // no1 的 sex
	studentMap["no2"]["address"] = "nanjing"    // no1 的 address

	fmt.Println(studentMap)
	fmt.Println(studentMap["no1"]["address"]) // 灵活地取舍
}

// ---------------------------struct 作为 value-----------------------------------

// Stu 学生结构体
type Stu struct {
	Name  string
	Age   int
	Grade float64
}

func student() {
	students := make(map[string]Stu, 10) // 这个map的value是结构体
	stu1 := Stu{"tom", 10, 99.4}
	stu2 := Stu{"jerry", 18, 95.4}
	stu3 := Stu{"cjj", 23, 99}
	students["no1"] = stu1 // 直接把定义好的结构体赋值过来就完事了
	students["no2"] = stu2
	students["no3"] = stu3
	fmt.Println(students)

	for k, v := range students {
		fmt.Println("编号", k)
		fmt.Println("名字", v.Name)
		fmt.Println("年龄", v.Age)
		fmt.Println("成绩", v.Grade)
	}
}

/*-----------------------------------map test ------------------------------------
使用map[string]map[string]string的数据类型
key是用户名，是唯一的
如果用户名存在。密码改为888888，如果不存在就增加这个用户（昵称 + code）
写一个modifyUser(users map[string]map[string]string,name string)*/
// todo 这个回头要复习再做
func modifyUser(users map[string]map[string]string, name string) {
	if users[name] != nil {
		users[name]["code"] = "888888"
	} else {
		users[name] = make(map[string]string, 2) // 要想新增也要make一下才行
		users[name]["code"] = "888888"
		users[name]["nickname"] = "tom"
	}

	//for k, v := range users {
	//	if k == name {
	//		v["code"] = "888888"
	//	} else {
	//		v[name] = name
	//		v[name]["nicname"] = name
	//		v["code"] = "123456"
	//	}
	//
	//}
	fmt.Println(users)
}
