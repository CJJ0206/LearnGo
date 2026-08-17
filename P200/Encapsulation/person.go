package Encapsulation

import "fmt"

type person struct {
	Name   string
	salary int
	age    int
	// 内部这里大写的话外面只要有实例就能直接调用
}

// NewPerson 工厂函数:是外部可以创建私有实例
func NewPerson(Name string) *person {
	return &person{
		Name: Name,
	}
}

func (p *person) SetName(Name string) {
	if Name != "" {
		p.Name = Name
	} else {
		fmt.Println("输入有误")
	}
}
func (p *person) SetAge(age int) {
	if age > 0 && age < 65 {
		p.age = age
	} else {
		fmt.Println("输入错误")
	}
}

func (p *person) GetName() string {
	return p.Name
}
func (p *person) GetSalary() int {
	return p.salary
}
func (p *person) GetAge() int {
	return p.age
}
