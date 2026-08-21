package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Here is testCut.")
	testCut()

	fmt.Println("Here is testCutPrefix.")
	testCutPrefix()

	/*  strings.Cut
	 func Cut(s, sep string) (before, after string, found bool)
	所属包：strings
	形参说明：
		s string :待切割的目标字符串
		sep string :用于查找的分隔字符串
	返回值类型：
		before string:分隔符 sep 首次出现位置之前的子字符串
		after string :分隔符 sep 首次出现位置之后的子字符串
		found bool : 指示是否找到分隔符 sep 。若为true则是找到了；若为false则before为原字符串s,after字符串为“”。
	设计背景与优势：
		strings.Cut 将查索引、安全截取、状态判断三合一，代码极为优雅，避免了此函数之前的各种麻烦。

	匹配规则：
		仅按首次出现的 sep 进行切割。如果字符串中有多个相同的 sep，后续的 sep 会全部留在 after 中。如果需要从后往前切，可结合 strings.LastIndex。
	多重分割：
		如果需要按分隔符切割成多段（如处理 CSV 或按逗号切割多项），应使用 strings.Split。
	*/

	// 场景一：未匹配到分隔符
	str := "hello world"
	before, after, found := strings.Cut(str, "#")
	fmt.Println("Found:", found)   // 输出: Found: false
	fmt.Println("Before:", before) // 输出: Before: hello world
	fmt.Println("After:", after)   // 输出: After: (空字符串)

	// 场景二：解析键值对（Key-Value）
	header := "Content-Type: application/json"
	key, val, found := strings.Cut(header, ": ")
	if found {
		fmt.Printf("Key: '%s', Value: '%s'\n", key, val) // 输出: Key: 'Content-Type', Value: 'application/json'
	}

	// 场景三：处理包含多个分隔符的情况（只按第一次出现切割）
	dsn := "user:pass@tcp(127.0.0.1:3306)/dbname"
	user, rest, ok := strings.Cut(dsn, ":")
	if ok {
		fmt.Printf("User: '%s', Rest: '%s'\n", user, rest) // 输出: User: 'user', Rest: 'pass@tcp(127.0.0.1:3306)/dbname'
	}

	// ------------------------------------------------
	/*  strings.CutPrefix  / strings.CutSuffix  （切前缀/切后缀）
	func CutPrefix(s, prefix string) (after string, found bool)
	所属包：strings
	形参：
		s  string : 待检查和徐建的目标源字符串
		prefix string : 需要匹配并去除的前缀字符串
	返回值类型：
		after string : 去除前缀后的子字符串。若未匹配到前缀，则返回原字符串 s
		found bool : 指示 s 是否以 prefix 开头。若匹配则返回 true,否则返回false。
	核心功能：
		检查字符串 s 是否以 prefix 开头，如果是，则剥离该前缀并返回剩余的部分

	注意事项：
	1.与 strings.TrimPrefix 的区别：
		TrimPrefix(s, prefix)：无论是否匹配只返回修剪后的字符串，不告诉你“到底有没有包含该前缀”。
		CutPrefix：多返回一个 found bool，适合需要依据“是否存在前缀”进行条件分支处理的场景（如解析带有 Bearer 的 Token 报头）。
	2.后缀裁剪需求：同理，如果需要处理后缀（Suffix），可以使用姐妹函数 strings.CutSuffix(s, suffix string) (before string, found bool)。
	*/

	// 场景一：解析 Authorization 请求头里的 Bearer Token
	authHeader := "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"
	token, ok := strings.CutPrefix(authHeader, "Bearer ")
	if ok {
		fmt.Println("成功提取 Token:", token) // 输出: 成功提取 Token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9
	} else {
		fmt.Println("格式错误：未包含 Bearer 前缀")
	}

	// 场景二：前缀不匹配时的处理
	url := "https://golang.org"
	httpURL, found := strings.CutPrefix(url, "http://")
	fmt.Println("是否包含 http:// 前缀:", found) // 输出: false
	fmt.Println("结果字符串:", httpURL)         // 输出: https://golang.org (保持原样)

}
