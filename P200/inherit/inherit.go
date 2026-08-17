package main

import "fmt"

/*
8.13
继承
我们现在需要设计一个学生考试系统
现在针对大学生、小学生、研究生三个系统，完全是可以复用的，所以需要继承

go 实现继承使用的是在结构体里再包一个共有的匿名结构体
*/

type goods struct {
	name  string
	price float64
}

type book struct {
	goods // fixme 用一个没有名字的结构体嵌入在里面
	pages int
}

type fish struct {
	goods
	num int
}

// todo 因为这两个结构体里都有这个匿名结构体，后面只要加这个匿名结构体的方法，那他们两都能使用

// GetInfo 的方式都是一样的所以没必要写两个
func (g goods) GetInfo() {
	fmt.Printf("该商品为%s,价格为%f \n", g.name, g.price)
}
func (b book) GetPage() {
	fmt.Printf("这本书有%d页 \n", b.pages)
}
func (f fish) GetNum() {
	fmt.Printf("鱼还有%d条 \n", f.num)
}

func main() {
	book1 := &book{}
	book1.pages = 326
	book1.goods.name = "book" // 层级关系也保留
	book1.goods.price = 25.5
	book1.GetInfo() // 可以看到，能够直接使用匿名结构体下的方法
	book1.GetPage()

	f1 := &fish{}
	f1.num = 10
	f1.goods.name = "fish"
	f1.goods.price = 15.3
	f1.GetInfo()
	f1.GetNum()
}
