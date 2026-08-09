package main

// 8.8

// 区分切片和数值
// fixme 在定义时，唯一的语法区别就是中括号 [] 里面有没有指定长度
// fixme 在 Go 语言中，区分数组（Array）和切片（Slice）的“黄金法则”就是看中括号里有没有数值

import "fmt"

var slice []int

// CreateSlice1 方式一:让切片去引用一个已经创建好的数组 ===============================
func CreateSlice1() {
	var arr = [5]int{1, 2, 3, 4, 5}
	slice := arr[2:5] // 这里只要不是越过cap，他会自动往后取的
	// 不过这里刚好是左闭右开，正好取到5
	fmt.Println(slice)
	// fixme 这种形式的创建slice是共享参数的，arr会把slice暴露在外面
}

// CreateSlice2 方式二:通过make来创建切片 ==========================================
// 基本语法： var 切片名 []类型 = make([]类型,len,[cap])  // 容量可以不具体指定
// 总结就是不管是数组还是切片，永远都要注明数据类型
func CreateSlice2() {
	var slice = make([]int, 4) // 类型也可以直接在make的时候指定
	// go 会自己找一块合适的区域存放它
	fmt.Println(slice)       // 默认全0
	slice = append(slice, 1) // 往slice后面添加元素
	slice = append(slice, 2)
	slice[0] = 23
	fmt.Println(slice)
	fmt.Println(len(slice))
	// TODO 那么使用make创建的slice在内存中是怎么布局的呢
	// fixme 但是通过make创建的slice指向的是一个不暴露的内存地址
}

// CreateSlice3 方式三:定义一个切片，直接制定具体数组，类似make
func CreateSlice3() {
	slice := []int{1, 2, 3} // 切片也是可以直接 := 创建的
	arr2 := [3]int{1, 2, 3} // 数组和切片创建的区别只在于中括号里有没有数字
	var arr = [5]int{1, 2, 3}

	fmt.Printf("slice的类型是%T \n", slice)
	fmt.Printf("arr的类型是%T \n", arr)
	fmt.Printf("arr的类型是%T \n", arr2)
}

// =========================== 遍历 slice ======================================

func forSlice() {
	slice := []int{1, 2, 3}
	for i := range slice {
		fmt.Println(slice[i])
	}
}

// ========================== 初始化方法细节 =====================================

func TestInit() {
	var slice1 []int
	var arr = [5]int{1, 2, 3, 4, 5}
	slice2 := arr[2:5]
	slice3 := arr[:]  // 直接取所有元素
	slice4 := arr[3:] // 从初始位置直接取到最后
	slice5 := arr[:4] // 从最开始取到对应位置

	// 如果对一个空地址的slice进行赋值的话是会报错的
	// 所以需要先引用数组或者make空间、append元素也会创建空间
	fmt.Println(slice1) // fixme 直接创建是一个空的切片 指向 oxo 这个地址，这是固定的零地址或空地址
	fmt.Printf("slice1的地址是%p \n", slice1)
	fmt.Println(slice2)
	fmt.Printf("slice2的地址是%p \n", slice2)
	fmt.Println(slice3)
	fmt.Println(slice4)
	fmt.Println(slice5)

	// todo 切片是可以继续切片的
	sliceInSlice := slice3[1:4] // 从另一个切片中再取切片
	fmt.Println(sliceInSlice)   // 这样嵌套就又参数共享了
	sliceInSlice[1] = 100
	fmt.Println(sliceInSlice) // 会直接修改掉上才层的值
	fmt.Println(slice3)

}
