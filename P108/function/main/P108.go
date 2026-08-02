package main // 7.31
// main入口所在的文件夹的所有文件都只能用package main

import (
	"fmt"
	"learn/P108/function"
	"learn/P108/package/utils"
)

func main() {
	fmt.Println(function.Calculate(3, 3, '*'))

	// 计算 iou
	groundTruth := [4]float64{100.0, 100.0, 200.0, 200.0}
	prediction := [4]float64{120.0, 120.0, 220.0, 220.0}
	iou := utils.CalculateIoU(groundTruth, prediction)
	fmt.Println(iou)

	// 函数底层机制解析
	n1 := 10
	// 在调用函数时会为函数开辟单独的栈区空间存储中间变量
	// 一旦函数执行结束系统就会销毁这个空间
	function.Test(n1) // 这里Test函数使用完n1后，会结束进程，直接把函数里的n1的空间回收了
	fmt.Println()
	fmt.Println("main函数汇中的 n1 =", n1) // 同时呢main栈区里的n1未被销毁，所以这里就输出了原始的n1

	fmt.Println(function.Sum(1222, 22))

	res1, res2 := function.SumSub(13, 7) // 不想要某个结果就直接下划线忽略就行
	fmt.Printf("和为：%d,差为%d", res1, res2)

}
