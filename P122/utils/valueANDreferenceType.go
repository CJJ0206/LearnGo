package utils

import "fmt"

// 8.4 函数形参的值类型和引用类型

// TODO 因为值类型和引用类型是重难点，之恶理性参数默认是值传递，引用类型参数默认是引用传递
// 不管是值传递和引用传递，传递的都是变量的副本(都是拷贝)，值传递是值的拷贝，引用传递是地址的拷贝;一般地址拷贝效率更高因为数据量小

// 之前讲过如果希望函数内的参数可以修改函数外的变量，我们是传递的变量的地址

// SCADA 项目里传的就是结构体的引用类型
// func initCommunicationService(app *runtimeApp) { // *runtimeApp 这个参数类型是run.go里定义的结构体
// TODO 这里传的引用类型形参，传的结构体地址，因为传地址效率更高，相比结构体里那么多数据，地址很小

// 变量作用域 ==========================================================
// 当全局变量 和 局部变量同名是 会使用就近原则

var Name2 = "cjj"

func Test01() {
	fmt.Println(Name2)
}
func Test02() {
	Name2 := "jack"
	fmt.Println(Name2)
}

func Test04() {
	Name2 = "jack" // 这里不会报错，但是会修改原来的值
	fmt.Println(Name2)
}

func Test03() {
	fmt.Println(Name2)
}

// 在函数体外部的任何代码语句，都必须以 Go 语言的关键字开头
// a := 55  // 这句是会报错的
var a int = 55
