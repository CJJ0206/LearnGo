package main

import "fmt"

/*
8.17
结构体里面还可以嵌入匿名基本数据类型，那如何访问呢

多重继承
如一个struct嵌套了多个匿名结构体，那么该结构体可以直接访问嵌套的匿名结构体的字段和方法，从而实现多重继承
前面那个 tv 继承 grand 和 goods 的例子就是多重继承
但是为了保持代码简洁呢，尽量还是不要使用多重继承
但是入宫多重继承的两根匿名结构体含有同名属性，那么访问时就一定要指定哪个匿名结构体

*/

type A struct {
	name string
	age  int
}
type B struct {
	A
	float64
}

func main() {
	var b B
	fmt.Println(b.name, b.float64) // fixme 可以看到数据类型由于匿名可以直接通过实例.访问
}
