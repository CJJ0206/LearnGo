package main

import (
	"fmt"
	_ "unsafe" // 一个下划线，包放最后，表示、先不使用这个包
)

// 数据类型间的相互转换：golang不像别的语言可以自动进行转换他必须显式地切换

func main() {
	var i int32 = 100
	var n1 float32 = float32(i)
	var n3 = int8(i)
	fmt.Println(i, n1, n3)

	// 在转换是虽然有些精度不会报错但结果与我们期望的不一样

	// 题目一
	//var n1 int32= 12
	//var n2 int64
	//var n3 int8
	//n2 = n1 + 20  这里加法的结果是int32的而n2是int64的所以会报错
	//n3 = n1 + 20

	// 题目二
	//var n1 int32 = 12
	//var n3 int64
	//var n4 int8
	//n4 = int8(n1) + 127  这里会溢出
	//n3 = int8(n1) + 128  这里赋值类型不匹配

	var a1 int32 = 12
	a2 := int8(a1) + 127 // 输出是 -117 因为数值已经溢出了  即：127 + 12 = 139 结果溢出到符号位上了 10001011
	fmt.Println(a2)
}
