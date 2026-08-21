package main

import (
	"fmt"
	"slices"
	"strings"
)

type User struct {
	ID   int
	Name string
	Tags []string // 注意：Tags 为切片，属于不可比较类型
}

func main() {
	/*slices.Contains
		func Contains[S ~[]E, E comparable](s S, v E) bool
		所属包：slices
		形参说明：
			s S：带查找的切片，其中元素类型为E（E 必须满足comparable约束，即支持 == 运算符）
			v E：需要匹配的目标元素
		返回值类型：
			bool（若切片中存在与 v 相等的元素则返回true,否则返回false）
		核心功能：
			判断切片中是否包含指定元素
		底层机制：
			采用 O(n) 线性扫描（对切片进行 for range 遍历），使用 Go 内置的 == 运算符逐个比较 v == s[i]。
			匹配到第一个相同元素时立即短路返回 true
	注意事项与避坑指南类型限制：
		切片元素类型 E 必须可比较（comparable）。如果切片元素包含切片（[]T）、Map（map[K]V）或函数（func）等不可比较类型，
		编译时会直接报错。这种情况下需使用下面介绍的 slices.ContainsFunc。性能与长切片：时间复杂度为 $O(n)$。
		如果切片已经排好序，建议改用 $O(\log n)$ 的 slices.BinarySearch；如果频繁查找且无序，建议改用 map[E]struct{}。
	*/
	languages := []string{"Go", "Python", "Rust", "Java"}

	// 查找切片中是否存在指定字符串
	hasGo := slices.Contains(languages, "Go")
	hasC := slices.Contains(languages, "C++")

	fmt.Println("包含 Go:", hasGo) // 输出: true
	fmt.Println("包含 C++:", hasC) // 输出: false

	testContain()

	/* -----------------------------------------------------------------------------
	slices.ContainsFunc : func ContainsFunc(s S, f func(E) bool) bool
	所属包：slices
	形参说明：
		s S：待查找的切片（切片数据类型任意）
		f func(E)：判断谓词函数，接收当前元素E，返回bool
	返回值类型：
		bool(只要有一个元素满足谓词函数f,就返回true,遍历完都不满足返回false)
	核心功能：
		使用自定义断言、谓词函数判断切片中是否存在符合条件的元素
	递进优势：
		解决了 slices.Contains 无法处理复杂结构体匹配、模糊/条件匹配或不可比较类型切片的问题。
	性能注意：
		谓词函数 f 会在遍历过程中被多次调用，应避免在 f 内部编写高开销操作（如发起网络请求或昂贵的内存分配）。
	并发安全：
		如果在 f 内部修改了切片元素或外部共享变量，需要自行保证并发安全；建议保持 f 为无副作用的纯函数。
	*/
	users := []User{
		{ID: 1, Name: "Alice", Tags: []string{"admin", "dev"}},
		{ID: 2, Name: "Bob", Tags: []string{"user"}},
	}

	// fixme ContainsFunc 的底层会自动对切片做for循环，然后给func元素，所以func一直接收item就行了
	// 1. 场景一：匹配结构体的特定字段（查找是否存在 ID 为 2 的用户）
	hasUser2 := slices.ContainsFunc(users, func(u User) bool { return u.ID == 2 })
	fmt.Println("存在 ID=2 的用户:", hasUser2) // 输出: true

	// 2. 场景二：忽略大小写的字符串模糊匹配
	// 忽略大小写判断
	hasPython := slices.ContainsFunc(languages, func(lang string) bool { return strings.EqualFold(lang, "PYTHON") })
	fmt.Println("忽略大小写匹配 Python:", hasPython) // 输出: true

	fmt.Println("here is test1")
	test1()

}
