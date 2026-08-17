package main

import "fmt"

/*
1.结构体可以使用嵌套匿名结构体所有的字段和方法;即：首字母大写或者小写的方法、字段 都可以使用
2.匿名结构体字段访问可以简化(不过易混淆)
	func main(){
		var b B
		b.A.name = "cjj"  -->  等价于 b.name = "cjj"
	}
3.当结构体和匿名结构体有相同的字段或者方法时，编译器采用就近原则访问，如果希望指定访问匿名结构体的字段和方法，可通过匿名结构体名来区分
4.结构体嵌入两个或者多个匿名结构体，如两个匿名结构体有相同的字段和方法（同时结构体本身没有同名的字段和方法），再访问时，就必须明确指定匿名结构体的名字
*/

type A struct {
	Name string
	age  int
}

func (a *A) SayOk() {
	fmt.Println("A is ok")
}
func (b *B) SayOk() {
	fmt.Println("B is ok")
}
func (a *A) hello() {
	fmt.Println("hello", a.Name)
}

type B struct {
	A
	Name string
}

func main() {
	var b B
	// 可以看到不论是大写还是小写，全都可以使用（换包导入不知道还行不行）
	//b.SayOk()
	//b.hello()

	b.Name = "jack"
	b.age = 20

	//b.A.Name = "tom"
	//b.Name = "jack" // fixme 因为 b 里没有Name这个属性，所以编译器会自己到匿名结构体里去找

	b.SayOk()
	// b.A.Name = "jerry"  // 这样对A赋完之后b才能获取到值
	b.hello() // todo 当我们用b去调A的这个方法时，调用a.name是什么结果呢
	//func (a *A) hello() {
	//	fmt.Println("hello",a.Name)
	//}
	// todo 答案是空串

}
