package main

import (
	"fmt"
	"math/rand"
	"sort"
)

/*
8.18
简单引入一些接口的使用场景
1.美国现在要制造轰炸机、武装直升机，专家只需把飞机需要的功能/规格定下来即可，然后让别人具体实现。

2.func Sort(data Interface)
这个要求的接口里面是长这样的，主要实现了这个接口，就可以用这个接口来实现排序
type Interface interface {
	Len() int
	......
	Less(i, j int) bool
	Swap(i, j int)
}
这是官方文档的sort函数，在绝大多数语言中，写一个通用的排序算法很让人头疼，因为算法本身需要知道“怎么比较两个元素”以及“怎么交换两个元素”
Go 的解法极其优雅：它把排序算法（怎么排）和数据结构（排什么）彻底分开了。
fixme 这个函数我们只需要实现好len、less、swap这三个方法，那么就实现了这个接口，那么底层会自己处理不同的数据类型，不用我们管了
Sort 函数内部的算法在运行时，根本不在乎你传进来的是 int 切片、字符串切片，还是一个几万行的复杂用户结构体列表。
它就像一个蒙着眼睛的机器，它只知道通过 Len() 问数据有多长，通过 Less() 问谁大谁小，然后通过 Swap() 发号施令交换位置。

3.接口有利于管理和控制开发的进度，经理设置好三个接口，分为5个自定义类型分发下去给手下实现，由于接口定义好了，所以相当于有个约束
*/

type a interface {
	Say()
	Leave()
}
type integer int

func (i integer) Say() {
	fmt.Println("here is integer")
}
func (i integer) Leave() {
	fmt.Println("here is integer")
}

// 实现对Hero结构体切片的排序：sort.Sort(data Interface)
var intSlice = []int{0, -1, 10, 7, 90}

type Hero struct {
	name string
	age  int
}
type HeroSlice []Hero

func (h HeroSlice) Len() int {
	return len(h)
}

// Less 方法决定了使用什么样的标准排序
// 这里写的是按照hero的年龄降序
func (h HeroSlice) Less(i, j int) bool {
	return h[i].age > h[j].age
}
func (h HeroSlice) Swap(i, j int) {
	h[i], h[j] = h[j], h[i] // fixme 这两和下面是等价的
	// fixme Go 处理多重赋值时，它的求值顺序保证了交换安全性：它会先完整计算出等号右边所有值，然后统一赋值给等号左边变量

	//temp := h[i]
	//h[i] = h[j]
	//h[j] = temp
}

func main() {
	var i integer = 10
	var b a = i
	b.Say()
	i.Say()

	var stu = Student{}
	var usb Usb = &stu // fixme 如果是指针类型的话，这里调用也要通过对应的地址调用，否则报错
	usb.Say()

	sort.Ints(intSlice) // 经过sort的处理之后就已经是一个有序的slice了
	fmt.Println(intSlice)

	// 对结构体切片进行排序
	var heroes HeroSlice
	for range 10 {
		hero := Hero{
			name: fmt.Sprintf("英雄%d", rand.Intn(100)),
			age:  rand.Intn(100),
		}
		heroes = append(heroes, hero) // 这种已存在的变量是不用:的否则会重新申明
	}
	for _, v := range heroes {
		fmt.Println(v)
	}
	// 此时直接放入就可以实现排序
	sort.Sort(heroes) // fixme 直接传给sort之后就实现了降序排列
	// fixme 通过这种方式呢，只需要对less函数里的参数进行改变，直接接可以实现对不同属性的排序
	fmt.Println("-------------------------")
	for _, v := range heroes {
		fmt.Println(v)
	}
	// fixme 这里呢是作为接口的一种应用形式讲解的这个函数，现在这个已经不如使用 slices.SortFunc

	students := stuSlice{
		{"cjj", 88.8, 23},
		{"aaa", 58.8, 24},
		{"ccc", 77.8, 22},
		{"bbb", 56.8, 23},
	}
	sort.Sort(students)
	for _, v := range students {
		fmt.Println(v)
	}

}
