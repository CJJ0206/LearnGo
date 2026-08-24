package main

import "fmt"

// 类型断言 ：由于接口是一般类型，不知道具体类型，如果要转换成具体类型，就需要使用类型断言
// 断言的前提条件是你这个变量原来确实指向对应的类型

/*
当执行 usbArr[0].(phone) 时，这就是一个类型断言。它其实是给 Go 的运行时（Runtime）下达了一个指令：“把这个盒子打开，检查里面装的到底是不是 phone。”
下面为你拆解它在底层到底做了什么：
	当你写下 pho, ok := usbArr[0].(phone) 时，Go 运行时会按照以下步骤飞速执行：
	核对类型（比对第一个指针）： 运行时首先去看 usbArr[0] 的第一个指针，发现里面记录的实际类型是 phone。然后拿它和你断言的目标类型（也就是括号里的 phone）进行对比。
	提取数据（读取第二个指针）： 发现类型完全一致（比对成功），运行时就会顺着第二个指针，找到内存里那个真实的 {"iphone"} 数据。
	拷贝赋值： 将这份数据拷贝一份，赋值给左边的变量 pho。此时，pho 变成了一个真正的、拥有全部属性和方法的 phone 结构体。
	返回状态： 将状态标志 ok 设置为 true。
*/

type student struct {
	name string
}

// TypeJudge 写一个函数循环判断传入参数的类型
func TypeJudge(items ...any) {
	for _, x := range items {
		switch x.(type) { // fixme 这个x.(type)只能在switch中使用，外部只能使用%T
		case bool:
			fmt.Println("bool")
		case int:
			fmt.Println("int")
		case string:
			fmt.Println("string")
		case float32:
			fmt.Println("float32")
		case float64:
			fmt.Println("float64")
		case student:
			fmt.Println("student")
		case *student: // 包括这种自定义的类型都是可以识别到的
			fmt.Println("*student")
		default:
			fmt.Println("类型错误")
		}
	}
}
