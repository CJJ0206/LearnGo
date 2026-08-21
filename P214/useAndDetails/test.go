package main

import (
	"fmt"
)

type Usb interface {
	Say()
}

func (stu *Student) Say() {
	fmt.Println("Say")
}

// Student 将student切片，按照score从大到小排序
type Student struct {
	name  string
	score float64
	age   int
}

// var stuSlice []Student // fixme 这是定义一个变量，不是类型，所以传参列表不能写 stu stuSlice
type stuSlice []Student // fixme 这样才是对的

func (stu stuSlice) Len() int {
	return len(stu)
}
func (stu stuSlice) Less(i, j int) bool {
	return stu[i].score < stu[j].score
}
func (stu stuSlice) Swap(i, j int) {
	stu[i], stu[j] = stu[j], stu[i]
}
