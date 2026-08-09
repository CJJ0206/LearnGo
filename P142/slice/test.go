package main

import "fmt"

func test1() {
	fmt.Println("Here is test")
	var a = []int{1, 2, 3, 4, 5}
	slice = make([]int, 1)
	slice[0] = 520
	fmt.Println(slice)
	copy(slice, a) // 这里是会直接把原始值覆盖掉，然后520这个值就从内存消失了
	copy(a, slice) // 大小都不会报错
	fmt.Println(slice)
	fmt.Println(a)
}

/*
fixme 说明
这里我们用slice去截取数组，这样之后slice其实指向的就是这个数组的开头
然后我们再吧这个slice赋值给另一个slice2，此时这个三个内容均指向同一个地址，也就是arr的地址
所以通过修改任意一个里面的值，整体都是跟着变化的
引用类型，所以在传递时遵循引用传递机制
*/
func test2() {
	fmt.Println("Here is test2")
	var slice []int
	var arr = [...]int{1, 2, 3, 4, 5}
	slice = arr[:]
	var slice2 = slice
	slice2[0] = 520 // 这三个的地址是完全一样的，所以修改任意一个整体都会跟着变化
	slice[1] = 250
	arr[2] = 125

	fmt.Println(slice)
	fmt.Println(slice2)
	fmt.Println(arr)
}

func test3(slice []int) {
	slice[0] = 250
}

/*
编写一个函数fbn(n int)
实现接收一个int，能够将斐波那契数列放到切片中

fixme 这里注意直接使用切片记录斐波那契数列（复习）
*/
func fbn(n int) []int64 {
	slice := make([]int64, n)
	slice[0] = 1
	slice[1] = 1
	for i := 2; i < n; i++ {
		slice[i] = slice[i-1] + slice[i-2]
	}
	return slice
}
