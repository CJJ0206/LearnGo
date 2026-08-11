package main

import "fmt"

/*
8.11
结构体与结构体变量（实例/对象）的关系
							  1.字段、属性 Name、Age、Color
struct( 包含猫的左右特征 ) ————> 2.猫会做的行为（这就是后续的方法）
							  3.(可以通过这个结构体创建多个变量)


结构体与结构体变量（实例/对象）的区别和联系
	1.结构体是自定义的数据类型，代表一类事物
	2.结构体变量（实例）是具体的，代表一个具体变量

结构体实例在内存中的布局
	新建一个实例时：内存中会产生一个和struct密切相关的内存空间（按照属性的定义开辟内存取空值）
    实例是直接指向这块内存空间的，所以struct是值类型的

todo 结构体声明
type name struct{
	name1 type
	name2 type
}

*/

// Cat : 先定义一个Cat结构体，把各个属性放进去进行管理
type Cat struct { // 结构体名称是大写表示可以在其他空间被使用
	Name  string
	Age   int
	Color string
	Food  string

	// fixme 但是slice和指针初始值都是nil,如果使用需要先make
	Ability map[string]int
	Ptr     *int
	slice   []int
}

func catStruct() {
	cat1 := Cat{Name: "小白"} // 既可以直接赋值 fixme 也可以不写属性名，直接赋值，系统会提示，但是这样必须一次性填满
	cat1.Age = 3            // 也可以通过 . 赋值
	cat1.Color = "white"
	cat1.Food = "fish"
	cat1.Ability = make(map[string]int, 3) // todo 一定要make哦

	// 用一个map来统计猫的各个能力的分数
	cat1.Ability["run"] = 84
	cat1.Ability["eat"] = 99
	cat1.Ability["sleep"] = 99

	// todo 后续不知可以绑定属性，还可以绑定方法

	fmt.Println(cat1)

	// fixme 结构体是值类型的验证
	cat2 := Cat{Name: "牛魔王", Age: 600, Color: "red"}

	cat4 := cat2 // 这样其实就是把一个结构体变量直接声明并创建给了一个新变量
	// todo 结构体是值类型，默认值拷贝
	fmt.Println("覆盖过来的值是：", cat4)
	cat2.Ability = make(map[string]int, 1)
	cat2.Ability["run"] = 78
	cat4.Ability = make(map[string]int, 1)
	cat4.Ability["eat"] = 99
	fmt.Println("修改后的cat2为：", cat2)
	fmt.Println("修改后的cat3为：", cat4) // todo 可以发现，这两是胡不影响的，所以是值类型，值传递

	cat5 := &cat2 // 这样相当于 cat5 是个指针，这样是可以修改结构体变量的值的
	fmt.Println(cat5)
}
