package main

import "fmt"

/*
8.11
创建结构体变量的四种方式
	1.方式一：直接声明
		var person Person
	2.方式二：{}
		var person Person = Person{  }  // fixme 推荐使用这个，最简单，不易混淆
	3.方式三：&
		var person *Person = new(Person)  // todo 这样做，new出来的其实是一个指向person的指针
		var p2 = new(Person)  // 等价于这句
	4.方式四：{}
		var person *Person = &Person{}
		var person = &Person{}  // 同样等价省略数据类型

*/

type Person struct {
	name string
	age  int
	sex  string
}

func createStruct2() {
	p1 := Person{"cjj", 23, "man"}
	fmt.Println(p1)
}

func createStruct3() {
	var p2 = new(Person) // new 一个struct对象出来，那这个p2就是个指针了
	(*p2).name = "cjj"
	p2.name = "sby" // fixme 这两行在go里面是一样的，go为了简化，在底层就帮我们做了（*p）
	p2.age = 24
	p2.sex = "male"

	fmt.Println(*p2)
}

func createStruct4() {
	var person = &Person{} // 这里的person也是一个指针
	(*person).name = "cjj"
	person.age = 88      // fixme 同样底层简化了
	fmt.Println(*person) // 要带解引用符*，不然输出多一个&

	var person2 = &Person{"aaa", 66, "male"}
	fmt.Println(*person2)

}
