package main

import (
	"cmp"
	"fmt"
	"io/fs"
	"path/filepath" // path包里的
	"slices"
)

// slices.SortFunc
func testSort() {
	users := []User{ // 我这里能直接调用到main里的结构体
		{"Alice", 30},
		{"Bob", 22},
		{"Charlie", 25},
	}
	slices.SortFunc(users,
		func(a, b User) int {
			return cmp.Compare(a.Age, b.Age) // 看正负底层排序
		})
	fmt.Println(users) // todo 原地排序，直接输出就行
}

// filepath.WalkDir： func WalkDir(root string, fn fs.WalkDirFunc) error
func testWalkDir() {
	dir := "." // 这个就是当前根目录
	err := filepath.WalkDir(dir,
		func(path string, d fs.DirEntry, err error) error { // 第二个形参，传的是walkDir碰到文件时的处理逻辑
			if err != nil {
				return err
			}
			if !d.IsDir() && filepath.Ext(path) == ".go" { // 只打印 Go 源码文件路径（看不懂）
				fmt.Println("找到 Go 文件:", path)
			}
			return nil
		})
	if err != nil {
		fmt.Println("遍历出错:", err)
	}
}

// http.MaxBytesReader
