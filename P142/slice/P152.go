package main

import "fmt"

// 8.7
// 需要一个数组哟关于保存学生的成绩，但是学生的个数是不确定的

/*
切片基本介绍
1.slice
2.切片是数组的一引用，因此切片是引用类型，在进行传递时，遵守引用传递的机制
3.切片的使用和数组类似，遍历、访问、len 都一样
4.切片的长度是可以变化的，他是一个可以变化的动态数组
5.切片定义的基本语法
	var 变量名 []类型
	如：var a []int


切片在内存中的布局

在 64 位架构下，一个切片在内存中固定占用 24 个字节，它由三个 8 字节的字段组成

go的底层slice其实就是一个struct
type slice struct {
    array unsafe.Pointer // 指向底层数组的指针 (8 bytes)
    len   int            // 切片的当前长度 (8 bytes)
    cap   int            // 切片的容量 (8 bytes)
}

slice |0xc0... | len | cap |  // 切片的构成是切的首地址 | 长度 | 容量

所以如果要读取切片后续的数据，用的方法是在地址上加上对应数据类型的字节数（引用类型）
这样做的是在地址层面的操作，所以会带来极致的性能，但共享数据有可能会导致内存泄漏问题

*/

func main() {
	var intArr [5]int = [5]int{1, 2, 3, 4, 5}
	// intArr[1:3] 表示slice引用intArr这个数组的起始位置为1，左闭右开
	slice := intArr[1:3]
	fmt.Println(slice)
	fmt.Println(cap(slice)) // 输出的是切片的容量，这个一般是初始大小的两倍， 防止频繁扩大

	fmt.Println(&intArr[1])
	fmt.Println(&slice[0]) // 他们之间的地址是相同的

	// 那么我们通过slice可否去修改原数组的值呢
	slice[0] = 23
	fmt.Println(intArr) // 因为slice本质上是在arr的地址上做操作，所以他们是数据共享的

	CreateSlice1()
	CreateSlice2()
	CreateSlice3()

	forSlice()

	TestInit()

	TestAppend()
	TestAppend2()
	TestCopy()

	test1()
	test2()

	var slice3 = []int{1, 2, 3, 4}
	fmt.Println(slice3)
	test3(slice3) // 函数接受的是切片，那切片的本质是地址，那么函数就能够修改切片里面的值
	fmt.Println(slice3)

	sliceString()

	fmt.Println(fbn(20))
}
