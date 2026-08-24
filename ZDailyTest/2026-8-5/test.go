package main

import (
	"fmt"
)

// -------- 补充前面缺失的 OrderStatus 定义 --------

type OrderStatus string

const (
	StatusPending OrderStatus = "PENDING"
	StatusPaid    OrderStatus = "PAID"
)

func UpdateStatus(status OrderStatus) {
	fmt.Printf("[UpdateStatus]  当前订单状态: %s\n", status)
}

// -------------------------------------------------

// CreateUser 调用者很容易写错顺序：CreateUser("active", "admin", "alice")
func CreateUser(name string, role string, status string) {
	fmt.Printf("[CreateUser]    用户名: %s, 角色: %s, 状态: %s\n", name, role, status)
}

type Role string
type UserStatus string

// CreateUser2 编译器会强制你传入正确的类型，就算你想传错顺序都做不到
func CreateUser2(name string, role Role, status UserStatus) {
	fmt.Printf("[CreateUser2]   用户名: %s, 角色: %s, 状态: %s\n", name, role, status)
}

type ExecStatus string

const (
	StatusIdle    ExecStatus = "idle"
	StatusRunning ExecStatus = "running"
)

func UpdateStatus3(id string, status ExecStatus) {
	fmt.Printf("[UpdateStatus3] 任务ID: %s, 执行状态: %s\n", id, status)
}

// 实现 DayType 类型：底层 string，定义工作日/周末两个枚举常量，绑定一个方法返回该天是否要上班。
type Day string

const (
	workday Day = "workday"
	weekday Day = "weekday"
)

func (d Day) work() bool {
	return d == workday
}

func main() {
	UpdateStatus(StatusPaid) // ✅ 正确

	// fixme 不报错：因为这里作为传参是一个无类型常量，被隐式转换为了 OrderStatus
	UpdateStatus("Apple")

	// ❌ 编译报错：假设 userName 是 string 类型，即使它的值刚好是 "PAID" 也不行
	// userName := "PAID"  // fixme 这样写一定是一个 string，所以会报错
	// UpdateStatus(userName)  // 报错

	// 这行同样不报错，因为 "running" 是无类型常量，隐式转换为了 ExecStatus
	UpdateStatus3("A", "running")

	// 对比一下，如果这样调用 CreateUser2，编译器就会强制保护你：
	// CreateUser2("Alice", Role("admin"), UserStatus("active"))

	d1 := Day("workday")
	b := d1.work() // todo 可以直接通过变量调用
	fmt.Println(b)
}
