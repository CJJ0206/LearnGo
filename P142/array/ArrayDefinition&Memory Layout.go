package main

import "fmt"

// 8.7

/*
数组定义和内存布局
	var 数组名 [数组大小]数据类型
	var a [5]int

数组的内存是连续的，直接用地址来取值是最快的


*/

func TestArray() {
	var arr [3]int // 8个字节
	arr[2] = 20
	// 当我们订一晚数组时，其实里面的元素就已经有默认值 0 了
	fmt.Println(arr)
	fmt.Printf("%p \n", &arr) // %p 地址格式化
	fmt.Println(&arr[0])      // 其实可以看到数组的地址就是首元素的地址
	fmt.Println(&arr[1])      // 第一个位置加对应数据类型的字节数（十六进制的）
	fmt.Println(arr)
}

func TestInput() {

	var arr [5]float64
	var score float64

	for i, _ := range arr {
		// 这个效果是等价于 for i := range arr 的
		// 只传一个值默认是索引，对于range 20这种纯数值的现在只会返回值了
		fmt.Println("倾输入成绩：")
		_, err := fmt.Scanln(&score)
		if err != nil {
			fmt.Println(err)
		}
		arr[i] = score // 上面取索引刚好拿来用
	}
	fmt.Println(arr)
}

// FIXME 四种初始化数组的方式
func name() {

	// 可以看到不管怎么定义，数据类型都是必不可少的
	var numsArray1 [3]int = [3]int{1, 2, 3}               // 申明加赋值
	var numsArray2 = [3]int{1, 2, 3}                      // 这个是等价于上面那个的
	var numsArray3 = [...]int{1, 2, 3, 4, 5, 6}           // ... 的话就可以自己任意填几个元素
	var names = [3]string{1: "tom", 2: "jerry", 0: "cjj"} // 指定对应位置上的值

	fmt.Println(numsArray1, numsArray2, numsArray3)
	fmt.Println(names)

}
