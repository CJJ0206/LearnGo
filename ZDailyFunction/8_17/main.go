package main

import (
	"cmp"
	"fmt"
	"math/rand"
	"slices"
	"strings"
)

type Employee struct {
	Department string
	Salary     int
}

type LogEntry struct {
	Timestamp int64
	Level     string
	Message   string
}

func main() {
	/*  cmp.Compare
	func Compare[T cmp.Ordered](x, y T) int
	形参说明：
		x T, y T：两个待比较的变量。类型约束 T 必须实现 cmp.Ordered（涵盖所以支持<>运算符的基础类型，所有符号、无符号类型、浮点、string）
	返回值类型：int
		若 x < y : 返回 -1
		若 x == y : 返回 0
		若 x > y : 返回 1
	核心功能：对两个有序类型的值执行标准的三路比较（Three-way Comparison）。

	在 Go 原生逻辑中，NaN == NaN 永远为 false。
	但 cmp.Compare 规范了浮点数的全序关系：它将 NaN 定义为严格小于任何非 NaN 数，且当 x 与 y 同为 NaN 时返回 0；-0.0 与 +0.0 被判定为相等并返回 0。

	注意事项与避坑指南
	多字段排序利器：在与 slices.SortFunc 配合时，可直接通过 if n := cmp.Compare(a.Field1, b.Field1); n != 0 { return n } 实现优雅的多级字段比较。
	不支持非基础有序类型：若比较结构体或自定义非基本类型，无法直接传给 cmp.Compare，需提取其内部的基础有序字段进行比较。
	*/
	fmt.Println("5 与 10 比较:", cmp.Compare(5, 10))       // 输出: -1
	fmt.Println("Go 与 Go 比较:", cmp.Compare("Go", "Go")) // 输出: 0
	fmt.Println("b 与 a 比较:", cmp.Compare("b", "a"))     // 输出: 1

	// 结合 slices.SortFunc 实现多级排序（部门升序，部门相同则薪资降序）
	employees := []Employee{
		{"Tech", 12000},
		{"HR", 8000},
		{"Tech", 15000},
	}

	// fixme Sort 函数是直接接收切片然后就可以排序
	// fixme 但是 SortFunc 接收的是一个func，go的函数名含func的基本都是这个意思
	/*
		无论底层用什么高级算法，它在对比两个元素 a 和 b 时，只看返回的 int 值的正负号：
		返回 < 0（比如 -1）：表示在排序结果中，a 应该排在 b 的前面（即 a < b）
		返回 0：表示 a 和 b 的排序权重相等。在非稳定排序中，它们相对位置随意；在稳定排序中，保留原有相对位置
		返回 > 0（比如 1）：表示在排序结果中，a 应该排在 b 的后面（即 a > b）
	*/
	// 底层会自动遍历 传进去的slice
	slices.SortFunc(employees,
		func(a, b Employee) int {
			if n := cmp.Compare(a.Department, b.Department); n != 0 { // 部门升序 （这里用传入顺序来控制的升序）
				return n
			}
			return cmp.Compare(b.Salary, a.Salary)
		}) // 部门相同时，薪资降序（交换 a 和 b 的位置）
	fmt.Printf("排序结果: %+v\n", employees)
	// 输出: [{Department:HR Salary:8000} {Department:Tech Salary:15000} {Department:Tech Salary:12000}]

	testSortFunc()

	/* slices.Compact / CompactFunc
	func CompactFunc[S ~[]E, E any](s S, eq func(E, E) bool) S
	形参说明：
		s S：待去重的目标切片。元素类型 E 为 any（支持不可直接比较的复合结构体或切片）。
		eq func(a, b E) bool：自定义等价判定函数。接收相邻的两个元素 a 和 b，若判定二者“等价/重复”则返回 true。
	返回类型：S：去重后的新切片视图。
	核心功能：根据自定义规则原地去除切片中相邻的重复元素。
	递进优势：
		解决了结构体切片中“按某特定字段去重”（如按用户 ID 去重，忽略其他字段变化）或“模糊去重”（如忽略大小写、浮点数在容差范围内去重）的需求。
	注意事项与避坑指南
		等价函数传递顺序：参数 eq(a, b) 中，a 代表已保留的基准元素，b 代表待比对的下一个元素。
		排序配合：与 slices.SortFunc 配合使用时，排序规则与等价判定的字段逻辑必须保持一致。
	*/
	var IDs [5]int
	slice1 := []int{}
	for i := range 5 {
		IDs[i] = rand.Intn(5)
		slice1 = append(slice1, IDs[i])
	}
	fmt.Println(slice1)
	fmt.Println(slices.Compact(slice1))
	fmt.Println("-----------------------------------------------------------------------")
	// [3 3 1 1 2]
	// [3 1 2]

	// 连续重复日志折叠（只看 Level 和 Message 是否相同，忽略 Timestamp）
	logs := []LogEntry{
		{Timestamp: 1000, Level: "INFO", Message: "User login"},
		{Timestamp: 1001, Level: "INFO", Message: "User login"}, // 与上一条重复
		{Timestamp: 1002, Level: "WARN", Message: "Disk full"},
		{Timestamp: 1003, Level: "INFO", Message: "User login"}, // 不连续，保留
	}

	logs = slices.CompactFunc(logs, func(a, b LogEntry) bool {
		return a.Level == b.Level && a.Message == b.Message
	})
	fmt.Printf("日志折叠结果（保留 %d 条）:\n", len(logs))
	for _, l := range logs {
		fmt.Printf("  [%s] %s (Time: %d)\n", l.Level, l.Message, l.Timestamp)
	}

	// 忽略字符串大小写的连续去重
	words := []string{"go", "Go", "GO", "rust", "Rust"}
	words = slices.CompactFunc(words, func(a, b string) bool {
		return strings.EqualFold(a, b) // 忽略大小写判断是否相同
		// fixme 所以总的来说，这个函数就是用来去除func判断为true的相邻的第一个重复元素
	})
	fmt.Println("忽略大小写去重结果:", words) // 输出: [go rust]

	testCompactFunc()

}
