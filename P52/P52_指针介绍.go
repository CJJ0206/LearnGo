package main

import (
	"fmt"
)

/*
& 获取地址
%p：专门用来打印指针自身的内存地址。
%v：通用占位符
指针只能指向地址
*/

/*
指针使用类型
值类型:都有对应的指针类型，形式为 *数据类型
值类型包括:基本数据类型、int系列、float系列、bool、string、数组和结构体

内存主要分为堆区、栈区

值类型：变量直接存储其值，内存通常在栈中分配
引用类型:变量存储的是一个地址，这个地址对应的空间才是真正存储数据值，内存通常在堆上分配，
当没有任何变量引用这个地址时，该地址对应的数据空间就成了一个垃圾被回收。

主要就是一个是变量直接存储其值，一个变量存储的是指向其值的地址（如整数就是值类型，指针就是引用类型）
*/

func main() {
	var i int = 10
	fmt.Println("i的地址是", &i) //直接println取地址 0x6ab69bbc0a0

	var ptr *int = &i // ptr是一个指针变量  类型是*int  值为&i
	fmt.Printf("ptr的地址为%p,值为%v\n", &ptr, ptr)

	var ptr2 **int = &ptr // 变成二级指针才可以指向另一个指针的位置
	fmt.Printf("ptr2的地址为%p,值为%v\n", &ptr2, ptr2)

	// * 解引用，获取指针指向的地址了里的真实值
	fmt.Printf("ptr指向的值为: %v\n", *ptr)   // 这个输出ptr指向的内存里的10
	fmt.Printf("ptr2指向的值为: %v\n", *ptr2) // 这个输出ptr2指向的ptr里的10的地址

	// 连续两次**直接获取最深层的数值
	fmt.Printf("通过ptr2获取最终的值为 %v\n", **ptr2) // 输出 10

	//
	// test
	var num int = 9
	fmt.Printf("num的地址是：%v\n", &num)
	// 需要通过指针修改他的值
	var ptr3 *int = &num
	fmt.Printf("未修改的值为：%v\n", *ptr3)
	*ptr3 = 99
	fmt.Printf("修改后的值为：%v\n", *ptr3)

	//
	// true or false
	//var a int = 64
	//var ptr4 *float64 = &a  这样是不允许的，指向什么类型的指针必须接收相应类型
	//fmt.Println(ptr4)

	var a = 300
	var b = 400
	var ptr4 *int = &a
	*ptr4 = 100
	ptr4 = &b // 这种是允许的，只要类型对即可
	*ptr4 = 200
	fmt.Printf("a=%d,b=%d,ptr4=%d\n", a, b, *ptr4)

}
