package main

import (
	"fmt"
	"learn/P200/Encapsulation"
	"learn/P200/factoryMode/factory"
)

func main() {
	// 看如果是普通大写，是可以通过import调用的,但是小写就不行
	//var stu1 = factory.Student{
	//	Name:  "cjj",
	//	Score: 86.6,
	//}

	// fixme 通过一个函数像工厂一样对类先加工再给外部使用
	var n = "cjj"
	var i = 86.6
	stu1 := factory.NewStudent(n, i) // fixme 这个stu始终是一个指针，所以看是不是引用传递，要看那边函数的形参
	fmt.Println(*stu1)
	// stu1.name = "cjj"  同样当属性是小写是同样外部不可访问
	score := stu1.GetScore() // 这样就可以得到那边私有的属性了
	*score = 99.1
	fmt.Println(*score)

	p1 := Encapsulation.NewPerson("xxx")
	p1.SetName("cjj")
	p1.SetAge(23)
	fmt.Println(p1.GetName())
	fmt.Println(p1.GetAge())

}
