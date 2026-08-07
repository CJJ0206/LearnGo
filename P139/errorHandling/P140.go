package main

// 8.7

import (
	"errors" // errors 的 New 方法用来创建一个可以被接收的自定义的错误信息
	"fmt"
	"math/rand/v2" // 用来生成随机数 rand，IntN(100)
	"time"
)

// golang 处理一场的方式是：defer,panic,recover
// go 抛出一个panic异常，然后再defer中通过recover捕获这个异常，然后正常处理

func main() {
	test() // 加上错误处理的代码之后，可以看到不仅输出了错误，而且能够执行到最后
	fmt.Println("继续执行下面的代码")

	fmt.Println("开始测试test2")
	test2()

	guessNum()
	fish()
}

func test() {
	// 使用 defer recover 的组合来捕获和处理异常
	defer func() {
		// err := recover() // recover 这个内置函数可以捕获到异常
		if err := recover(); err != nil { // 错误不为空，则说明有异常
			// if 语句里是可以定义申明的
			fmt.Println("err = ", err)
		}
	}() // 使用defer放在函数最上面做预防，一定那执行不下去会最后执行defers

	num1 := 2
	num2 := 0
	res := num1 / num2 // 直接就会报错 panic
	fmt.Println(res)
}

// 函数去读取配置文件init.config的信息
// 如果文件名传入不正确，就返回一个自定义的错误
func readConfig(name string) (err error) {
	if name == "config.init" {
		// 读取
		return nil
	}
	return errors.New("读取文件错误") // 返回一个自建的错误信息
}

func test2() {
	err := readConfig("config.yaml")
	if err != nil {
		// panic(err) // 使用panic是会直接终止程序的，但是只是打印信息不会
		fmt.Println(err)
	}
	fmt.Println("test2继续执行")
}

// 这是一个猜数字的案例，看着不难，做完整还是需要时间的（复习）
func guessNum() {
	randNum := rand.IntN(100) + 1
	fmt.Println(randNum)
	chance := 10

	var guessNum int
	for {
		fmt.Println("请输入你要猜的数字：")
		_, err := fmt.Scanln(&guessNum)
		if err != nil {
			fmt.Println("输入错误")
			continue
		}

		if guessNum != randNum {
			chance--
			if chance <= 0 {
				fmt.Printf("次数用光了，没办法了！正确答案是: %d\n", randNum)
				break // 【新增】机会耗尽，退出循环
			}
			// 可以在这里顺便提示一下大了还是小了
			if guessNum > randNum {
				fmt.Printf("猜大了！你还有 %d 次机会\n", chance)
			} else {
				fmt.Printf("猜小了！你还有 %d 次机会\n", chance)
			}
		} else if guessNum == randNum {
			fmt.Printf("共用了%d次,", 11-chance)
			switch 11 - chance {
			case 1:
				fmt.Println("你真是个天才")
			case 2, 3:
				fmt.Println("你很聪明")
			case 4, 5, 6, 7, 8, 9:
				fmt.Println("一般")
			case 10:
				fmt.Println("终于")
			default:
				fmt.Println("没办法了")
			}
			break
		}
	}

}

// 根据三天打鱼两天晒网这句话，判断1990.1.1开始，如何判断之后的某一天改干啥
// 这个主要是时间相关的
func fish() {
	inputStr := "2026-08-07"                               // 假设的输入
	startDate, _ := time.Parse("2006-01-02", "1990-01-01") // 定义基准起始时间 (使用 Go 的魔法参考时间化石)
	targetDate, _ := time.Parse("2006-01-02", inputStr)    // 将输入的字符串解析为 time.Time 对象
	diff := targetDate.Sub(startDate)                      // 两个时间对象直接相减，得到一个 Duration (时间段)
	totalDays := int(diff.Hours() / 24)                    // 将时间段转换为小时，再除以 24，得到总天数 (强转为 int)
	fmt.Printf("从 1990-01-01 到输入日期，一共经过了 %d 天\n", totalDays)

	if totalDays%5 == 1 || totalDays%5 == 2 {
		fmt.Println("晒网")
	} else {
		fmt.Println("打鱼")
	}

}
