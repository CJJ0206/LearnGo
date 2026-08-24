package main

import (
	"fmt"
)

/*
基本介绍
变量（实例）具有多种形态。面向对象的第三大特征，在go中，多态是通过接口实现的。
可以按照同一个接口来调用不同的实现。这时接口变量就呈现不同的多态。

接口体现多态特性的两种形式
1.多态参数
	在前面的Usb接口案例，Usb usb,即可以接收手机变量，又可以接收相机变量，就体现了该接口的多态
2.多态数组
	演示一个案例：给Usb数组中，存放phone结构体和camera结构体变量，
	todo 下面这部分后面将类型断言时再做
	phone还有一个特有的接口call，请遍历Usb数组，如果是phone变量，除了调用Usb接口声明的方法外，还需要调用phone特有的call
*/

type Usb interface {
	Start()
	end()
}
type phone struct {
	name string
}
type camera struct {
	name string
}

func (c camera) Start() {}
func (c camera) end()   {}
func (p phone) Start()  {}
func (p phone) end()    {}

func (p phone) call() { fmt.Println("it can call sb.") } // phone 的 call 方法

func main() {
	// 给usb数组中既存放phone结构体，又存放camera结构体（问题就是数组只能放一个类型，怎么办）
	// todo 定义一个usb接口数组
	var usbArr [4]Usb
	usbArr[0] = phone{"iphone"}
	usbArr[1] = camera{"canon"}
	// fixme 这他妈说白了就是数组类型时接口类型，所以只要是实现了这个接口的就都可以放入

	fmt.Println(usbArr)
	// 现在我们的数组里有两个结构体了，怎么去调用phone的call函数呢
	// 我们通过usbArr[0]. 只能调用到Usb里的Start方法
	if pho, ok := usbArr[0].(phone); ok { // fixme pho返回的是值拷贝回来的phone（不在用一个内存空间），但是也可以通过判断指针类型实现同步
		pho.call()
	} else {
		fmt.Println("转换失败")
	}

	// fixme 使用switch实现（x.（type）方法）
	for _, item := range usbArr {
		if item == nil { // 如果 item 是空的，跳过
			continue
		}
		// 这里就是 x.(type) 的用武之地！  item.(type) 会提取出真实的类型
		switch device := item.(type) { // fixme 这个语法只能跟 switch 语句配合使用
		case phone:
			fmt.Println("发现手机！") // 如果匹配到 phone，在这个 case 块里：
			device.call()        // 编译器会自动把 device 当作一个真正的 phone 结构体！
		case camera:
			// 如果匹配到 camera，在这个 case 块里：
			fmt.Println("发现相机！", device.name) // device 就自动变成了 camera 结构体
		default:
			fmt.Printf("未知的设备类型: %T\n", device)
		}
	}
}

/*
fixme 类型断言
type Point struct{
	x int
	y int
}
func main(){
	var a interface {}
	var point Point = Point(1,2)
	a = point
	// 如何将a赋值给一个Point变量
	var b Point
	b = a // 可以吗？
	fmt.Println(b)
}

*/
