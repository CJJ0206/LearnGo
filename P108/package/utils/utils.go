package utils

import (
	"math"
)

// CalculateIoU 计算两个边界框的交并比 (IoU)
func CalculateIoU(box1, box2 [4]float64) float64 {
	xLeft := math.Max(box1[0], box2[0])
	yTop := math.Max(box1[1], box2[1])
	xRight := math.Min(box1[2], box2[2])
	yBottom := math.Min(box1[3], box2[3])

	// 如果没有交集（右边界小于左边界，或下边界小于上边界），交并比为 0
	if xRight < xLeft || yBottom < yTop {
		return 0.0
	}

	// 计算交集面积
	intersectionArea := (xRight - xLeft) * (yBottom - yTop)

	// 计算两个边界框各自的面积
	box1Area := (box1[2] - box1[0]) * (box1[3] - box1[1])
	box2Area := (box2[2] - box2[0]) * (box2[3] - box2[1])

	// 计算并集面积并求出 IoU
	return intersectionArea / (box1Area + box2Area - intersectionArea)
}
