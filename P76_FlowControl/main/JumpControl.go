package main // 7.31

import (
	"fmt"
	"math/rand"
)

// 随机生成数字，当生成99时，看看打印了几次

func num() {
	var i int
	for {
		n := rand.Intn(100) + 1 // 这个函数现在已经不用手动设置随机种子了
		i++
		if n == 99 {
			break // 表示跳出
		}
	}
	fmt.Printf("用了%d次", i)
}

// break 细节说明------------------------------------------------------
/*
正常情况下只跳出本层循环
标签的使用: 循环可以用标签来命名（开头加标签：），break时指定特定的循环名，会直接强制找到它打断整体逻辑
*/
func labBre() {
OuterLoop: //没有用到是会报红的
	for i := 1; i <= 100; i++ {
	InnerLoop:
		for j := 1; j <= 100; j++ {
			if j == 23 {
				break InnerLoop
			}
		}
		if i == 66 {
			break OuterLoop
		}
	}
}

func login() {
	var name string
	var code string
	var realName string = "张三"
	var realCode string = "666666"

	for i := 1; i <= 3; i++ {
		fmt.Println("欢迎登录，你有三次机会请输入用户名和密码")
		_, err := fmt.Scanln(&name, &code)

		fmt.Printf("你还有%d次机会\n", 3-i)
		if err != nil {
			fmt.Println("输入错误，请重新输入:")
		} else if name != realName || code != realCode {
			fmt.Println("输入错误，请重新输入:")
		} else if name == realName || code == realCode {
			fmt.Println("您已登录成功！")
			break
		} else {
			fmt.Println("3次机会已用完，账号已锁定！")
		}

	}
}

// continue ----------------------------------------------------------
/*
continue语句用于结束本次循环，继续执行下一次循环
当continue出现再循环中时，也是可以通过标签来指明要跳出到哪一层循环的
*/
func testCont() {
	for i := 1; i <= 10; i++ {
		if i == 2 {
			continue // 会发现输出没有 2，是因为continue在到2的时候直接从自己跳出回for进行下一轮了
		}
		fmt.Print(i)
	}
}

// test ---------------------------------------------------------------

func JumpTest1() {
	for i := range 13 {
		if i == 10 {
			continue
		}
		fmt.Printf("i=%d", i)
	}
} // 结果是不输出10，其余输出

func JumpTest2() {
	for i := range 2 {
		for j := 1; j < 4; j++ {
			if j == 2 {
				continue
			}
			fmt.Printf("i=%d,j=%d", i, j)
		}
	}
} // 结果：（0,1）(0,3) (1,1) (1,3)

func JumpTest3() {
here:
	for i := range 2 {
		for j := 1; j < 4; j++ {
			if j == 2 {
				continue here
			}
			fmt.Println("i=", i, "j=", j)
		}
	}
} // 结果：（0，1） (1,1)

func oddNum() {
	// 打印 1 ~ 100 的奇数，用for + continue
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Print(i)
	}
}

func judgeNum() {
	var num int
	var countNe int
	var countPo int
	for i := 1; i <= 100; i++ {
		fmt.Print("请输入数据：")
		_, err := fmt.Scanln(&num)
		if err != nil {
			fmt.Print("输入有误。请重新输入。")
			continue
		}
		if num > 0 {
			countPo++
			continue
		} else if num < 0 {
			countNe++
			continue
		} else {
			fmt.Println("接收到 0，结束输入。")
			break
		}
	}
	fmt.Printf("\n--- 最终统计结果 ---\n")
	fmt.Printf("正数个数：%d\n", countPo)
	fmt.Printf("负数个数：%d\n", countNe)
}

// goto -----------------------------------------------------------------
/*
go语言的goto是可以无条件的转移到程序指定的行去的
goto通常与条件语句配合使用，实现条件转移、跳出循环体等功能
一般不主动使用goto，会造成程序混乱

所以我不学了
*/
