package main

import "fmt"

type Monkey struct {
	name string
}
type LittleMonkey struct{ Monkey } // 匿名结构体继承老猴子
type Bird struct{ name string }
type Superman struct{ name string }

// 1. 定义接口（你原来就有的）
type fly interface {
	flying()
}

// 2. 定义各种类型并实现 flying() 方法
func (l *LittleMonkey) flying() { fmt.Println(l.name + " flying") }

func (b *Bird) flying() { fmt.Println(b.name + " 扇翅膀飞") }

func (s *Superman) flying() { fmt.Println(s.name + " 穿斗篷飞") }

// 3. 通用函数，参数是 fly 接口类型  fixme 这个如果在代码复杂的情况下是值得存在的
func letItFly(f fly) {
	f.flying()
}

// 4. 使用
func main() {
	m := &LittleMonkey{Monkey{name: "悟空"}}
	b := &Bird{name: "小鹰"}
	s := &Superman{name: "克拉克"}

	// 但是我们这里比较简单，直接对象运行就行
	m.flying()
	b.flying()
	s.flying()

	letItFly(m) // 输出：悟空 flying
	letItFly(b) // 输出：小鹰 扇翅膀飞
	letItFly(s) // 输出：克拉克 穿斗篷飞

	// fixme 可以理解为，接口的底层是不需要我们做的，但是在外部有文档可以直接调用
	// fixme 比如这里，当我们是调用者且知道有这个接口时，我们直接做如下的调用就行，啥也不用管，那边的代码都已经完善了
	// fixme 现在看不出优势，只是因为我们兼顾实现者的角色，所以麻烦
	monkey := &Monkeys{Animals{Name: "悟空", Weight: 30.8}}
	bird := &Birds{Animals{Name: "小鹰", Weight: 1.6}}
	// 直接传入即可，Go 自动识别它们都实现了 Animal 接口
	Reporter(monkey)
	Reporter(bird)
}
