package main

// 8.21
// fixme 实现接口可以看作是对继承的一种补充

type athlete struct {
	Name string
}

type basketballPlayer struct {
	athlete
}
type footballPlayer struct {
	athlete
}

type student struct {
	Name string
}
type collegeStudent struct {
	student
}
type middleStudent struct {
	student
}

// 要求足球运动员和大学生掌握英语
// 理论上是可以直接对足球运动员和大学生各自绑定方法,但是这样会造成一个人一个写法的问题，不规范
// 所以用接口实现

type learnEnglish interface {
	learn()
}

// fixme 继承的价值主要是：提升代码复用性和可维护性
// fixme 接口的价值在于：设计，设计好各种规范（方法），让其它自定义类型去实现这些方法
