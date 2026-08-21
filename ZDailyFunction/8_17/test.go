package main

import (
	"cmp"
	"fmt"
	"slices"
	"strings"
)

type stuff struct {
	name   string
	age    int
	salary float64
}

type Log struct {
	Timestamp int64
	Level     string
	Message   string
}

func testSortFunc() {
	s := []stuff{
		{"cjj", 23, 400},
		{"sby", 24, 4600},
		{"qg", 23, 5000},
		{"gjl", 45, 24000},
		{"zx", 24, 4700},
	}

	slices.SortFunc(s, func(a, b stuff) int {
		// if n := cmp.Compare(a.age, b.age); n != 0 { // 这里比价的是如果a>b就会把a移到前面 这是升序
		if n := cmp.Compare(b.age, a.age); n != 0 { // 只需要把参数顺序倒过来就成了 倒序了
			return n
		} else if n == 0 {
			return cmp.Compare(a.salary, b.salary)
		}
		return cmp.Compare(a.name, b.name)
	})
	fmt.Println(s)
}

func testCompactFunc() {
	log := []Log{
		{Timestamp: 1000, Level: "INFO", Message: "User login"},
		{Timestamp: 1001, Level: "INFO", Message: "User login"}, // 与上一条重复（折叠）
		{Timestamp: 1002, Level: "WARN", Message: "Disk full"},
		{Timestamp: 1003, Level: "INFO", Message: "User login"},
	}
	logs := slices.CompactFunc(log, func(a, b Log) bool {
		return a.Level == b.Level && a.Message == b.Message
	})
	for _, l := range logs {
		fmt.Printf("  [%s] %s (Time: %d)\n", l.Level, l.Message, l.Timestamp)
	}
	fmt.Println(len(logs))

	word := []string{"go", "Go", "GO", "rust", "Rust"}
	words := slices.CompactFunc(word, func(a, b string) bool {
		return strings.EqualFold(a, b)
	})
	fmt.Println(words) // fixme 直接返回bool就行

}
