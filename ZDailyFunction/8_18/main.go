package main

import (
	"cmp"
	"fmt"
	"os"
	"slices"
	"strings"
)

type Task struct {
	ID        int
	Title     string
	Completed bool
}

func main() {
	/*  cmp.Or ： func Or[T comparable](vals ...T) T
	形参：
		vals ...T : 可变参数列表。类型T满足可对比类型
	返回值类型：
		T : 按从左到右顺序，返回参数列表中第一个不等于其类型零值的项；如果全部参数都是零值（或未传入），返回类型T的零值
	核心功能：未变量提供多级备用值/默认值回退链
	设计价值：
	*/
	// 场景一：读取环境变量，支持多级默认值回退
	userInputPort := ""
	envPort := os.Getenv("APP_PORT") // 从当前操作系统环境变量中读取名为 APP_PORT 的值
	defaultPort := "8080"

	finalPort := cmp.Or(userInputPort, envPort, defaultPort) // fixme 按照哪个生效就用哪个，始终有后备值可用(按照靠前的输出)
	fmt.Println("生效端口:", finalPort)                          // 输出: 生效端口: 8080

	// 场景二：处理数字类型的默认超时时间（秒）
	var userTimeout int = 0 // 用户未指定（零值 0）
	var configTimeout int = 0
	finalTimeout := cmp.Or(userTimeout, configTimeout, 30) // 回退到默认 30 秒
	fmt.Println("超时时间:", finalTimeout)                     // 输出: 超时时间: 30

	// 场景三：若全部都是零值，则安全返回零值
	emptyStr := cmp.Or("", "", "")
	fmt.Printf("全零值返回: '%s'\n", emptyStr) // 输出: 全零值返回: ''

	// ------------------------------------------------------------------------------------------------
	/* slices.delete : func Delete[S ~[]E, E any](s S, i, j int) S  fixme 这个很简明
	形参说明；
		s S : 待修改的目标切片，元素类型E可以为任意类型
		i int : 待删除区间的初始索引
		j int : 待删除区间的结束索引  （左闭右开）
	返回值类型：
		S : 删除指定区间的元素后的新切片视图（长度缩减为 len(s) - (j - i)）
	核心功能：
		原地从切片中移除索引范围在[i，j）内的子切片
	*/
	// 场景一：删除连续区间 [2, 4) 即索引 2 和 3 的元素
	letters := []string{"A", "B", "C", "D", "E", "F"}
	letters = slices.Delete(letters, 2, 4) // 删除了 "C", "D"
	fmt.Println("区间删除结果:", letters)        // 输出: [A B E F]

	// 场景二：删除指定单个索引（例如删除索引 1 的元素）
	nums := []int{100, 200, 300, 400}
	idx := 1
	nums = slices.Delete(nums, idx, idx+1) // 删除了 200
	fmt.Println("单项删除结果:", nums)           // 输出: [100 300 400]

	// -------------------------------------------------------------------------------------------------
	/*  slices.DeleteFunc ： func DeleteFunc[S ~[]E, E any](s S, del func(E) bool) S
	上面版本的函数版本，支持自定义写函数来delete
	*/

	// 场景一：批量清除已完成（Completed == true）的任务
	tasks := []Task{
		{ID: 1, Title: "编写接口文档", Completed: true},
		{ID: 2, Title: "重构数据库连接池", Completed: false},
		{ID: 3, Title: "修复登录漏洞", Completed: true},
		{ID: 4, Title: "编写单元测试", Completed: false},
	}

	// 传入谓词：t.Completed 为 true 时删除
	tasks = slices.DeleteFunc(tasks, func(t Task) bool {
		return t.Completed
	})

	fmt.Printf("未完成任务列表（共 %d 条）:\n", len(tasks))
	for _, t := range tasks {
		fmt.Printf("  - ID: %d, Title: %s\n", t.ID, t.Title)
	}

	// 场景二：删除空字符串或仅包含空格的项
	tags := []string{"Go", "  ", "CloudNative", "", "Backend"}
	tags = slices.DeleteFunc(tags, func(tag string) bool {
		return strings.TrimSpace(tag) == "" // 为空则删除
	})
	fmt.Println("清理后的有效标签:", tags) // 输出: [Go CloudNative Backend]

	testOr()
	testSliceDelete()
	testDeleteFunc()
}
