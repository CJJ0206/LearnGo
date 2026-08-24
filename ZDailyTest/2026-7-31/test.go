package main

import (
	"errors"
	"fmt"
	"strconv"
)

/*
strconv
test Itoa / Atoi and PharseInt // int to string / string to int

*/

func main() {
	i := 4
	s := "cjj"

	// Itoa 是一定可以成功的所以不会返回err
	// 但是Atoi有很大的风险，所以返回会多一个err(不是bool了的值，而是一个err类型的错误)
	i1 := strconv.Itoa(i)
	s1, ok := strconv.Atoi(s)
	if ok != nil {
		fmt.Println("转换失败")
	}
	s2, _ := strconv.Atoi("333")
	fmt.Printf("i1的数据类型是%T ,s1 的值为 %s \n", i1, i1)
	fmt.Printf("s1的数据类型是%T ,s1 的值为 %d \n", s1, s1)
	// 可以看到还是会输出一个结果，虽然转换失败，但还是会给他赋一个0值
	fmt.Printf("s2的数据类型是%T ,s1 的值为 %d \n", s2, s2)

	// strconv 包里的Parse..函数的的输入均为string,然后解析为对应类型
	// 可以看到下面的结果都是把string转换为对应类型
	// bitSize 这个参数的意思是目标数据类型的位大小（位宽），防止转换后溢出
	b, err := strconv.ParseBool("true")
	fmt.Printf("b的类型是%T \n", b)
	f, err := strconv.ParseFloat("3.1415", 64)
	fmt.Printf("b的类型是%T \n", f)
	i3, err := strconv.ParseInt("-42", 10, 64)
	fmt.Printf("b的类型是%T \n", i3)
	u, err := strconv.ParseUint("42", 10, 64)
	fmt.Printf("u的类型是%T \n", u)
	fmt.Println(b, f, i3, u, err)

	// fixme 利用 error 的多态性，精准拦截并处理这类 脏数据
	_, Err := strconv.ParseInt("hello", 10, 64)
	if Err != nil {
		var numErr *strconv.NumError
		if errors.As(err, &numErr) { // 检查是不是 strconv 包抛出的特定错误
			if errors.Is(numErr.Err, strconv.ErrSyntax) {
				fmt.Println("用户输入的根本不是数字！请重新输入。")
			} else if errors.Is(numErr.Err, strconv.ErrRange) {
				fmt.Println("用户输入的数字太大了，装不下！")
			}
		} else {
			fmt.Println("发生了其他未知错误:", err)
		}
	}

}
