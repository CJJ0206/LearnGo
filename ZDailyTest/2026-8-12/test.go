package main

import "fmt"

type student struct {
	name string
	age  int
}

type Dog struct{ Name string }

func BarkFunc(d Dog) { fmt.Println(d.Name, "bark") }

func (d Dog) BarkMethod() { fmt.Println(d.Name, "bark") }

// 实现String方法，实现了fmt.Stringer接口
func (s *student) String() string {
	return fmt.Sprintf("name:[%s],age:[%d]", s.name, s.age)
}

func main() {
	s1 := student{"cjj", 23}
	s2 := &student{"cjj", 23}

	fmt.Println(&s1) // Println 调用时底层判断如果实现了String就会通过其格式输出
	fmt.Println(&s2)

	dPtr := &Dog{Name: "旺财"}
	// BarkFunc(dPtr)     // 这里会报错，因为函数严格要求类型匹配
	dPtr.BarkMethod() // fixme 这里则不会，因为方法具备语法糖特性：用指针变量调用值接收者方法时，编译器会自动做解引用 (*dPtr).BarkMethod()转换，编译顺利通过
}
