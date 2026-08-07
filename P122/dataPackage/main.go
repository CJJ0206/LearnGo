package main

import (
	"fmt"
	"time"
)

// 时间和日期相关函数都是在 time 包里的 , time.Time 类型，用于表示时间

/*
time.Now 函数获取当前时间 然后这个Time对象有各种子方法获取时间

一些时间常量的意义
const(
	Nanosecond Duration = 1 纳秒
	Microsecond 		= 1000 * Nanosecond = 1 微秒
	Millisecond 		= 1000 * Microsecond = 1 毫秒
	Second  			= 1000 * Millisecond = 1 秒
	Minute 				= 60 * Second = 1 分钟
	Hour				= 60 * Minute = 1 小时
) // 这几个常量是不用定义直接通过 time. 就可调用，用来当作时间单位

time.Sleep() 休眠函数

time.Unix（Nanosecond Microsecond Millisecond） 系列时间戳 （作用是获取随机数字）返回的是1970.1.1到现在的秒数/纳秒/....
// 这种不用随机种子，且随机性更高


*/

func main() {
	// 获取当前时间
	now := time.Now()
	fmt.Printf("time类型为：%T，值为：%v \n", now, now)

	// 通过 now 可以获取到 年月日、时分秒
	year := now.Year()
	month := now.Month()
	day := now.Day()

	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()

	// 格式化输出
	fmt.Printf("现在是%d-%d-%d %d:%d:%d \n", year, month, day, hour, minute, second)

	// Sprintf 把东西拼好，塞进一个变量里给程序用，不会在屏幕上显示任何东西，除非再去单独打印它
	dataStr := fmt.Sprintf("现在是%d-%d-%d %d:%d:%d", year, month, day, hour, minute, second)
	fmt.Println(dataStr)

	// 1月2日下午3点4分5秒2006年，这样记就可以
	fmt.Println(now.Format("2006-01-02 15:04:05"))
	fmt.Println(now.Format("2006/01/02 15")) // 截断到哪输出到哪

	// 需求：每隔1秒打印一个数字，打印到100
	// 需求2：每隔0.1秒打印
	i := 0
	for {
		i++
		fmt.Println(i)
		// time.Sleep(time.Second) // 休眠1秒
		time.Sleep(time.Millisecond * 100) // 只能这样写，别的方式不通过
		if i == 10 {
			break
		}
	}

	//Unix 和 UnixNano 的使用
	fmt.Println(now.Unix())
	fmt.Println(now.UnixNano())

	// 执行前获取时间戳 普通Unix的单位是秒
	start := time.Now().Unix()
	// test()
	end := time.Now().Unix()
	fmt.Printf("执行test消耗时间为：%v 秒 \n", end-start) // 十万次执行了4秒

}
