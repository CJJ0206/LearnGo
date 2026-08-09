package main

// 8.8 8.9
// func append(slice []Type, elems ...Type) []Type
// 可以看到官方的函数接收 slice 和 元素 返回 slice  (现在默认第一个参数大多是 老slice)

import "fmt"

func TestAppend() {
	fmt.Println("Here is the append function test.")
	var slice = []int{1, 2, 3}
	fmt.Printf("原始slice的地址是:%p,\n\t值为：%v \n", &slice, slice)
	slice = append(slice, 23, 24, 25) // 可以直接添加多个的，返回值赋值回去才真正改变了自身
	fmt.Printf("原始slice的地址是:%p,\n\t值为：%v \n", &slice, slice)
	slice = append(slice, slice...)
	// fixme ... 在 go 里面的意思是打散拆包，像python的*解包一样，一个个拆出来返回
	fmt.Println(slice)
}

/*
todo 分析append的底层原理(append(slice,1,2,3))

fixme Condition 1 (append超过容量3时，才会进行新建切片)
|0xc042.. | 3 | 3 |
(| 对应)
| 100 | 200 | 300 |   ————> | 0 | 0 | 0 | 0 | 0 | 0 |  // 创建扩容后的空切片（因为数组底层是不支持扩容的）
当执行append时，他首先会去创建一个数组，数组的大小就是再原基础上扩容

要先把原slice的值拷贝过去 ————> | 100 | 200 | 300 | 0 | 0 | 0 |
这个新的数组是go创建出来的，在填入原切片的之之后，go会根据新切片首地址去做对应的偏移，从而把新值填到后面

扩容后的切片如果还是赋到原先的slice上，那这个slice的指针其实时会变的，因为指针指向了这个新的slice

fixme Condition 2 （append仍在容量范围内时）
|0xc042.. | 3 | 6 |
此时底层数组长原本就长这样：| 100 | 200 | 300 | 0 | 0 | 0 |
当执行 append(slice, 1, 2, 3) 时，底层并不会产生拷贝旧数据的动作，而是直接写入新数据
*/

func TestAppend2() {
	fmt.Println("测试append扩容底层")
	var slice2 = make([]int, 4) // 我们指定slice的长度和容量均为4 （make不指定cap时就是等于len的）
	fmt.Printf("slice2的长度是：%d,slice2的cap是：%d \n", len(slice2), cap(slice2))
	fmt.Printf("slice2原来的地址是%p \n", slice2)
	slice2[0] = 9
	slice2[1] = 8
	slice2[2] = 7
	slice2 = append(slice2, 1, 2, 3)
	fmt.Printf("slice2的长度是：%d,slice2的cap是：%d \n", len(slice2), cap(slice2))
	// slice2 被覆盖到新地址后原来的slice就会被回收掉
	fmt.Printf("slice2后来的地址是%p \n", slice2) // fixme 这两处的打印是不加 & 的
	// fixme &slice2 是在查这个切片的外壳放在哪了
	// fixme 直接使用 slice2 配合 %p ; Go对这种写法有特殊处理,它会打印“外壳”的地址
	// fixme 而是去取外壳里的 Data 属性，打印出真正存放数据的底层数组的首地址

	var slice3 = make([]int, 4) // 我们指定slice的长度和容量均为4 （make不指定cap时就是等于len的）
	fmt.Printf("slice3的长度是：%d,slice3的cap是：%d \n", len(slice3), cap(slice3))
	fmt.Printf("slice3原来的地址是%p \n", slice3)
	slice2[0] = 9
	slice2[1] = 8
	slice2[2] = 7
	slice4 := append(slice2, 1, 2, 3) // 冒号 是需要创建时才用的
	// 这样就不会影响原本切片的地址
	fmt.Printf("slice4的长度是：%d,slice4的cap是：%d \n", len(slice4), cap(slice4))
	fmt.Printf("slice4后来的地址是%p \n", slice4) // fixme 这两处的打印是不加 & 的

}

// TestCopy ==================== copy 讲解 ================================
func TestCopy() {
	// 切片使用copy内置函数完成拷贝
	var slice1 = []int{1, 2, 3}
	var slice2 = make([]int, 6) // 六个0
	// 他两的数据空间是相互独立的

	copy(slice2, slice1)
	fmt.Println(slice2) // 输出可以看到后面的3个0是不会别覆盖或者删除的
	// 这里地址肯定是不变的，只是把值拷贝了过来
}
