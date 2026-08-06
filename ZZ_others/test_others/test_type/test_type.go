package main

// 8.5

import "fmt"

// 1. 核心操作：定义一个“专属 VIP 通行证”类型

type ExecStatus string

// 2. 定义该类型允许的枚举常量
const (
	StatusIdle    ExecStatus = "idle"
	StatusRunning ExecStatus = "running"
	StatusFailed  ExecStatus = "failed"
)

// 3. 优势展示 A：给自定义类型绑定方法（普通 string 做不到）
// 假设我们需要在前端 UI 上显示友好的中文状态

func (s ExecStatus) ToChinese() string {
	switch s {
	case StatusIdle:
		return "空闲中 ☕️"
	case StatusRunning:
		return "疯狂运行中 🏃‍♂️"
	case StatusFailed:
		return "执行失败 ❌"
	default:
		return "未知状态 ❓"
	}
}

// 4. 优势展示 B：极致的类型安全
// 这个函数的第二个参数，被死死限制为了 ExecStatus 类型

func UpdateStatus(scriptID string, status ExecStatus) {
	// 在函数内部，我们可以直接调用它绑定的方法
	fmt.Printf("系统通知：脚本 [%s] 的状态已更新为 -> %s (%s)\n", scriptID, status, status.ToChinese())
}

func main() {
	// ✅ 正确示范：传入定义好的常量
	UpdateStatus("Script_A", StatusRunning)
	UpdateStatus("Script_B", StatusFailed)

	// ❌ 错误示范：如果你把下面这行代码取消注释，程序根本无法编译！
	// UpdateStatus("Script_C", "running")
	// 编译器会直接无情报错：cannot use "running" (untyped string constant) as ExecStatus value...
}
