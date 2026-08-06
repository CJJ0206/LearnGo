package utils

// 8.3

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

// ===============================init 函数======================================

var Age int
var Name string

// Age 和 Name 全局变量，我们要在main.go中使用
// 但是我们要先进行初始化
// := 只能写在函数内部，写在外面会报错

func init() {
	fmt.Println("utils init function()...") // 这里是最先执行的
	Age = 25
	Name = "Tom"
} // 这种简单的初始化往往直接 var 定义就行

// 操作系统的例子
// 但是遇到每次运行都要选择的init时，往往是通过复杂条件判断执行的，这只是个演示
var (
	OSName        string // 自定义的操作系统名称
	PathSeparator string // 模拟针对不同系统的路径分隔符
)

func init() {
	switch runtime.GOOS { // runtime.GOOS 会返回当前系统的底层标识
	case "darwin":
		// 注意：苹果 macOS 系统的底层内核叫 Darwin，所以在 Go 里它的名字是 "darwin"
		OSName = "macOS"
		PathSeparator = "/"
		fmt.Println("🍎 系统初始化：检测到 macOS 环境，加载 Mac 专属配置...")

	case "linux":
		OSName = "Linux"
		PathSeparator = "/"
		fmt.Println("🐧 系统初始化：检测到 Linux 环境，加载 Linux 专属配置...")

	case "windows":
		OSName = "Windows"
		PathSeparator = "\\"
		fmt.Println("🪟 系统初始化：检测到 Windows 环境，加载 Win 专属配置...")

	default:
		fmt.Println("❓ 系统初始化：检测到未知或少见系统...")
	}
}

// ============================== 匿名函数 ========================================
// 如果我们希望函数只被使用一次，可以使用匿名函数，匿名函数就是灭有名字的系统
// 在定义匿名函数时直接调用，这种方式匿名函数只能使用一次

// =======================commonly used string function===========================

func TestStr() {
	str2 := "hello北京"
	fmt.Println(len(str2))       // 一个汉字占三个字节
	for _, value := range str2 { // range 遍历字符串时返回的value时rune
		fmt.Printf("%s", string(value))
	}
} // value 的底层类型是rune , 同时还加了强转 string 所以是绝对安全的

func TestStr2() {
	str := "hello南京"
	r := []rune(str)
	for _, value := range r {
		fmt.Printf("%c", value)
	}
	fmt.Println()
}

func TestAtoi() {
	n, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("转换错误") // 处理 非数字的 字符串会把报错
	}
	fmt.Println(n)
}

func ConvByte() {
	var byteType = []byte("Hello World") // 强转用的是圆括号
	fmt.Println(byteType)
}

func ConvStr() {
	str := string([]byte{97, 98, 99}) // 这里面在初始化赋值所以是花括号
	fmt.Println(str)                  // a b c
}

func Contain() {
	val := strings.Contains("Hello World", "Wor")
	fmt.Printf("输出类型是 %T, 输出值为 %v \n", val, val)
}

func Contain2() {
	num := strings.Count("Hello World", "o")
	fmt.Println(num)
}

func Equal() {
	bo := strings.EqualFold("Hello", "hello") // 这里就是相等的
	fmt.Println(bo)
	// fmt.Println("hello" == "Hello") // 这个会直接提醒不相等
}

func JudgeExist() {
	re := strings.Index("Hello World", "l") // 打印 l 第一次出现的位置
	fmt.Println("l第一次出现的位置是：", re)
}

func LsatSee() {
	pos := strings.LastIndex("Hello World", "l")
	fmt.Println("l 最后出现的位置是：", pos)
}

func Replace() {
	str2 := "go go hello"
	str := strings.Replace(str2, "go", "go 语言", 1) // 直接传递变量是会改变原值的
	fmt.Println(str)
}

func Split() {
	strArr := strings.Split("hello , cjj , what", ",") // 按照逗号拆成了三个字符串的数组
	fmt.Println(strArr)
	for index, value := range strArr {
		fmt.Println(index, value)
	}
}

func UpLow() {
	str := strings.ToLower("GOOGBY")
	fmt.Println(str)
}

func Clear() {
	str := " hello word "
	fmt.Println(str)
	str2 := strings.TrimSpace(str)
	fmt.Println(str2)

	str3 := "! No No No !"
	str4 := strings.Trim(str3, "! ") // 这里是可以合并空格在里面的
	fmt.Println(str4)

	str5 := "?。 what s wrong"
	fmt.Println(strings.TrimLeft(str5, "?。"))

	str6 := "yo yo yo >"
	fmt.Println(strings.TrimRight(str6, ">"))
}
