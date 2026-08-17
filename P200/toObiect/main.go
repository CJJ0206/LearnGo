package main

import (
	"fmt"
	"strings"
)

/*
8.13
步骤:
	1.声明（定义）结构体，确定结构体名
	2.编写结构体的字段
	3.编写结构体的方法

案例：
	1.编写一个Student结构体，包含name,gender,age,id,score字段，分别为string,string,int,int,float64
	2.结构体中声明一个say方法，返回string类型，方法返回信息中包含所有字段值
	3.再main方法中，创建Student结构体实例，并风闻say方法，并将调用结果输出

盒子案例：
	1.编程创建一个Box结构体，在其中声明三个字段来表示一个立方体的长宽高，且要从终端获取
	2.声明一个方法获取立方体体积
	3.创建一个Box结构体实例，打印给定尺寸的立方体的体积

景区门票案例：
	1.景区根据人的年龄收取不同价格，大于18收20，其他免费
	2.编写Visitor结构体，根据年龄决定门票价格并输出
	3.要求出入姓名年龄，当输入name为q时，结束
*/

type Student struct {
	Name   string
	Age    int
	Gender string
	Id     int
	Score  float64
}

func (s *Student) say() {
	fmt.Printf("my Name is %s ,age is %d,id is %d,gender is %s,score is %f", s.Name, s.Age, s.Id, s.Gender, s.Score)
	fmt.Println()
}

type Box struct {
	Height float64
	Width  float64
	Length float64
}

func (b *Box) area() float64 {
	fmt.Print("请输入对应的长宽高尺寸：")
	_, err := fmt.Scanln(&b.Height, &b.Width, &b.Length)
	if err != nil {
		fmt.Println(err)
	}
	return b.Height * b.Width * b.Length
}

type Visitor struct {
	Age   int
	Price float64
	Name  string
}

func (v *Visitor) visit() {
	// 通过这种指针结构去修改对应的票价
	if v.Age > 18 {
		v.Price = 20
	} else {
		v.Price = 0
	}
}

func main() {
	stu1 := Student{"cjj", 23, "man", 001, 93.4}
	stu1.say()

	var box1 Box
	fmt.Println(box1.area())

	var vis Visitor
	for {
		fmt.Printf("请输入年龄和姓名：")
		_, err := fmt.Scanln(&vis.Age, &vis.Name)
		if err != nil {
			fmt.Println(err)
		}
		if strings.EqualFold(vis.Name, "q") {
			break
		}

		fmt.Println("进函数前:", vis.Price)
		vis.visit()
		fmt.Println("进函数后:", vis.Price)
	}

	// fixme new 其实用的不多，除非真的需要0值
	vis2 := &Visitor{23, 200, "cjj"} // fixme 和 new 类似，但是new有初始0值，这个没有
	fmt.Println(*vis2)
	fmt.Printf("指向的地址是%p，类型是%T \n", vis2, vis) // 直接就是个指针类型
	fmt.Println(vis2.Name)

}
