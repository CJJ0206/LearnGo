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
}
