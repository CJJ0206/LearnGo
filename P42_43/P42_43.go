package main

import (
	"fmt"
	"unsafe"
)

// bool:golang的布尔值只能为True或者False，不像其他语言可以为0/非0
/*
基本数据类型的默认值：
整数：0
浮点：0
bool:false
string:"" (空字符串)

%v 按照原始值输出（通用的通配符）
*/

func main() {
	const isFalse = false
	fmt.Println(isFalse)
	fmt.Printf("b占大小为 %v\n", unsafe.Sizeof(isFalse)) // 占一个字节
	i := 65
	fmt.Printf("%c\n", i+1) // 65 是大写的a , 66 是大写的b

	var str = "hello"
	// str[0] = 'a'  // 是无法修改字符串中的字母的
	fmt.Println(str)

	str2 := "hi\ncjj\n" // 是允许包含转义字符的
	fmt.Println(str2)

	str3 := `  
	const isFalse = false
	fmt.Println(isFalse)
	fmt.Printf("b占大小为 %v\n", unsafe.Sizeof(isFalse)) // 占一个字节
	i := 65
	fmt.Printf("%c", i+1) // 65 是大写的a , 66 是大写的b
	var str = "hello"
	// str[0] = 'a'  // 是无法修改字符串中的字母的
	fmt.Println(str)
	` // 反引号，用来包围输出长串内容
	fmt.Println(str3)

	var str4 = "hi" + "cjj" + "whats" + // 字符串拼接的加号只能放在上面否则会报错
		"your" + "name"
	fmt.Println(str4)
}
