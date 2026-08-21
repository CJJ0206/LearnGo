package main

import (
	"fmt"
	"sync"
)

// context.WithTimeOut

// sync.OnceFunc
func testOnceFunc() {
	testFunc := sync.OnceFunc(
		func() {
			fmt.Println("执行该函数")
		})
	fmt.Print("第一次执行结果：")
	testFunc()
	fmt.Println() // 锁定后只有第一次会执行

	fmt.Print("第二次执行结果：")
	testFunc()
	fmt.Println()

	fmt.Print("第三次执行结果：")
	testFunc()
}

// errors.Is
