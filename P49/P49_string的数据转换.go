package main

/*
%d:整形输出
%v:值的默认格式表示
%t:true/false  bool的默认值是false
%e:科学计数法
%q:会自动把输出加上双引号，且如果有转义字符会一起输出
%s:就是普通的string的通配符
%T:输出数据类型
%f:浮点类型
*/

import (
	"fmt"
	"strconv"
)

func main() {
	var num1 int = 99
	var num2 float32 = 23.456
	// var b bool = true
	var str string = "hello"

	// 使用第一种方法来转换（fmt.Sprintf 方法）把各种类型转成 string
	str = fmt.Sprintf("%d\n", num1) // 接收
	fmt.Println(str)

	str2 := fmt.Sprintf("%f", num2) // 这个str2由于未提前申明所以需要用一个 : 来申明一下
	fmt.Printf(str2)

	// 第二种转换方式（strconv）把各种数据类型转为string
	var num3 int = 99
	//var num4 float32 = 23.456
	const b2 bool = true

	// strconv.FormatInt 强制要求传入的第一个参数类型必须是 int64
	str = strconv.FormatInt(int64(num3), 10) //10位小数
	fmt.Printf("转换后的数据类型是%q\n", str)

	str = strconv.FormatBool(b2)
	fmt.Printf("转换后的数据类型是%q\n", str)

	// 第三种防方式 strconv包中的 Itoa 函数
	var num5 = 4567
	str = strconv.Itoa(num5)
	fmt.Printf("转换后的数据类型是%s\n", str) // 直接就转成了字符串
	// 这些操作都是根据填入数据的值去创造一个对应的str而原数据本身是没有改变的

	var num6 int64 = 3456
	str = strconv.Itoa(int(num6)) // Go认为int是一个独立的数据类型，有些函数的输入强制是这个类型，就需要转换
	fmt.Printf("转转的数据类型是%q\n", str)

	/*
		使用strconv包的函数有(这些函数会返回两个值)
		ParseBool
		ParseFloat
		ParseInt
		ParseUint
	*/

	// string 转基本数据类型
	var str5 string = "true"
	var b bool
	b, _ = strconv.ParseBool(str5) //
	fmt.Printf("数据类型是：%T,数据为%t\n", b, b)

	var str6 = "123456"
	var i int64
	i, _ = strconv.ParseInt(str6, 10, 32) // 转成10进制，64位的，这个32只是用来限制转换后不要溢出32位
	// 这个函数永远返回64位
	fmt.Printf("数据类型是%T，值为%d\n", i, i)

	//
	//
	// 1. 解析并限制在 32 位以内（返回的 val64 依然是 int64）
	val64, err := strconv.ParseInt(str6, 10, 32)
	if err != nil { // Go中nil代表空 / 没有值
		fmt.Println("转换失败或超出32位范围:", err)
	} else {
		// 2. 因为前面已经通过了 32 位的安全校验，现在你可以放心地强转，绝对不会丢失精度！
		var finalVal int32 = int32(val64) // 需要使用强转来转换
		fmt.Printf("最终数据类型是%T，值为%d\n", finalVal, finalVal)
	}

	// 转float
	var str7 string = "3.1415"
	var f float64                       // 这个转Float和上面的转int接收的参数都是默认64位
	f, _ = strconv.ParseFloat(str7, 64) // 需要其他位数的值就需要强转
	fmt.Printf("数据类型是 %T，数据位 %f\n", f, f)

	//
	// 注意
	var str8 string = "hello" // 由于无法将hello转为int，所以会赋值一个默认值
	var n3 int64
	n3, _ = strconv.ParseInt(str8, 10, 64)
	fmt.Printf("数据类型是 %T，数据为 %d \n", n3, n3)

}
