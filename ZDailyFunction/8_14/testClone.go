package main

// fixme 8.14 注意（important）

import (
	"fmt"
	"slices"
)

type class struct {
	resumes []string
}

// NewClass 这里相当于是一个工厂函数
func NewClass(resumes []string) *class {
	return &class{slices.Clone(resumes)} // 入参也克隆，避免外部后续修改
	// todo 把这个结构体暴露出去，并且带着参数resumes出去，返回的是整个结构体实例
}

// Items class 的方法，用来返回一个实例的克隆体
func (c *class) Items() []string {
	return slices.Clone(c.resumes)
}

func (c *class) Items2() []string {
	return c.resumes
}

func test() {
	cfg := NewClass([]string{"a", "b"})
	got := cfg.Items() // 对原始值做一个clone，确保原始永远不被篡改
	got[0] = "hacked"
	fmt.Println(got)         // [hacked b]
	fmt.Println(cfg.Items()) // [a b]

	cfg2 := NewClass([]string{"a", "b"})
	got2 := cfg2.Items2()
	got2[0] = "dead"
	fmt.Println(got2)
	fmt.Println(cfg2.Items2()) // 可以看到再去获取这个实例的值时已经被修改了
}
