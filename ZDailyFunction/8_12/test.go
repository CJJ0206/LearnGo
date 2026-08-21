package main

import (
	"fmt"
	"slices"
	"strings"
)

func testContain() {
	// strings.Contains:比较两个string里是否包含指定字段
	arr := [3]string{"aaaa", "bbb", "cjj"}
	ok := strings.Contains(arr[0], "a")
	fmt.Println("包含 a :", ok)

	// slices.Contains:判断整个切片中是否有某个指定‘用户’
	slice := []string{"cjj", "hello", "sby"}
	ok2 := slices.Contains(slice, "sby") // 单查找字母是查不出来的，只能查找完整字段
	fmt.Println("包含 cjj :", ok2)
}

/*
📝 练习题 1：角色权限校验（练手 slices.Contains）
【业务场景】：在 Web 系统中，你需要校验当前用户的角色（role）是否有权限访问某敏感接口。
【题目要求】
已知系统允许访问的角色列表为：allowedRoles := []string{"admin", "editor", "auditor"}
请编写一个函数 HasAccess(userRole string) bool。
在函数内部使用 slices.Contains 判断 userRole 是否在允许列表中，如果在则返回 true，否则返回 false。


📝 练习题 2：电商库存与商品检索（练手 slices.ContainsFunc）
【业务场景】：电商后台需要快速判断商品列表中是否存在“缺货的高价商品”，以及是否存在指定名称的商品。
【题目要求】
请使用 slices.ContainsFunc 完成以下两个判断（可以直接写在 main 函数中）：
需求 A：判断商品列表中是否存在 “价格高于 5000 元 且 库存少于 5 件” 的预警商品。
需求 B：判断商品列表中是否存在 “商品名称包含 'macbook'（忽略大小写）” 的商品。
*/

var allowedRoles = []string{"admin", "editor", "auditor"}

// Product 已知商品结构体与初始数据如下
type Product struct {
	ID    int
	Name  string
	Price float64
	Stock int
}

func HasAccess(userRole string) bool {
	// 用slices.Contains判断userRole是否在允许列表中，如果在则返回true，否则返回false
	ok := slices.Contains(allowedRoles, userRole)
	return ok
}

func test1() {
	fmt.Println("cjj的权限是：", HasAccess("cjj"))
	fmt.Println("admin的权限是：", HasAccess("admin"))
}

func test2() {
	products := []Product{
		{ID: 101, Name: "iPhone 15 Pro", Price: 7999.0, Stock: 10},
		{ID: 102, Name: "MacBook Air", Price: 8999.0, Stock: 2}, // 高价且库存低
		{ID: 103, Name: "小米手环", Price: 249.0, Stock: 0},
	}
	ok := slices.ContainsFunc(products, func(item Product) bool {
		return item.Price > 5000 && item.Stock < 5
	})
	//ok2 := slices.ContainsFunc(products, func(item Product) bool {
	//	return slices.Contains(item, "macbook")
	//})
	// 需求 B：修正后的写法
	ok2 := slices.ContainsFunc(products, func(item Product) bool {
		return strings.Contains(strings.ToLower(item.Name), "macbook")
	})
	fmt.Println("有价格高于 5000 元 且 库存少于 5 件：", ok)
	fmt.Println("有macbook：", ok2)
}
