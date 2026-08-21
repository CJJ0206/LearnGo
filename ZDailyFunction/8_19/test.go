package main

import (
	"cmp"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"slices"
)

func testSliceBinary() {
	sli := []float64{1, 2, 3, 4, 5}
	res, ok := slices.BinarySearch(sli, 4)
	fmt.Println(res, ok) // 3 true

	target := 3.5
	res2, ok2 := slices.BinarySearch(sli, target)
	fmt.Println(res2, ok2) // 3 false
	// todo 这里返回的3是他遍历后推荐的插入位置，然后直接通过insert就可以实现保序插入
	sli = slices.Insert(sli, res2, target) // fixme 这种对原切片做操作一定要覆盖给原值，否则输出不变
	fmt.Println(sli)
}

func testSliceBinaryFunc() {
	accounts := []Account{
		{ID: 1009, Name: "Alice"},
		{ID: 1005, Name: "Bob"},
		{ID: 1008, Name: "Charlie"},
		{ID: 1001, Name: "Jerry"},
	} // 由于二分接收的是一个升序，所以我们先处理成升序
	slices.SortFunc(accounts, func(a, b Account) int { // todo 这个是就地排序不需要覆盖回去
		return cmp.Compare(a.ID, b.ID)
	})
	for _, account := range accounts {
		fmt.Println(account)
	}
	// 然后我们做查找
	target := 1005
	id, found := slices.BinarySearchFunc(accounts, target, func(a Account, tar int) int {
		return cmp.Compare(a.ID, tar)
	})
	if found {
		fmt.Printf("成功找到 ID=%d 的账户: %+v (索引: %d)\n", target, accounts[id], id)
		// 输出: 成功找到 ID=1005 的账户: {ID:1005 Name:Bob} (索引: 1)
	} else {
		fmt.Println("账户不存在")
	}

	target1 := 1000
	res1, ok1 := slices.BinarySearchFunc(accounts, target1, func(a Account, tar int) int {
		return cmp.Compare(a.ID, tar)
	})
	fmt.Println(res1, ok1)
	accounts = slices.Insert(accounts, res1, Account{ // fixme 这里插入的元素类型一定要和account一致，直接放target1是报错的
		ID: target1, Name: "Tom",
	})
	for _, account := range accounts {
		fmt.Println(account)
	}
}

func testError() {
	err := closeResources()
	if err != nil {
		fmt.Printf("聚合报错输出:\n%v\n", err)
	}
	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("能够感知到包含“文件不存在”错误")
	}
	if errors.Is(err, os.ErrClosed) {
		fmt.Println("能够感知到包含“进程关闭”错误")
	}
}

func closeResources() error {
	var errs []error
	// 模拟执行多个清理操作
	if err := os.Remove("temp_file_1.txt"); err != nil {
		errs = append(errs, fmt.Errorf("文件1不存在，清理失败: %w", err))
	}
	if err := os.Remove("temp_file_2.txt"); err != nil {
		errs = append(errs, fmt.Errorf("文件2不存在，清理失败: %w", err))
	}
	// 聚合所有收集到的错误（若 errs 全为 nil 则返回 nil）
	return errors.Join(errs...) // fixme ... 是go的解包操作，把切片拆解成可变参数传入函数
}

// ... 解包使用
func testJieBao() {
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5, 6}
	// 将 s2 中的所有元素逐个追加到 s1
	s1 = append(s1, s2...) // s1 变为 [1, 2, 3, 4, 5, 6]

	// 准备动态参数列表
	args := []any{"Alice", 28, "Engineer"}
	// fmt.Printf(format string, a ...any)
	fmt.Printf("Name: %s, Age: %d, Role: %s\n", args...)

	paths := []string{"base", "modules", "auth", "token.go"}
	// filepath.Join(elem ...string) string
	fullPath := filepath.Join(paths...) // "base/modules/auth/token.go"
	fmt.Printf("FullPath: %s\n", fullPath)
}
