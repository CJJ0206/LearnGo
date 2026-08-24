package main

import (
	"fmt"
	"time"
)

func main() {
	p := new(int)   // new 会创建一个指定类型大小的空间，返回该内存的地址（指针）,兵并且会赋零值
	fmt.Println(p)  // fixme 这是p指向的地址
	fmt.Println(&p) // fixme 这个是p自己的地址
	fmt.Println(*p)

	num1 := 100
	fmt.Printf("num1类型=%T, 值=%v, 地址=%v\n", num1, num1, &num1)
	num2 := new(int)
	fmt.Printf("num2类型=%T, 值=%v, 地址=%v, 指向值=%v\n", num2, num2, &num2, *num2)

	// var m map[string]int
	// m["key"] = 1 // 这里不会显式报错，但是运行会panic
	m := make(map[string]int) // fixme 切片、map、channel 都需要使用make去初始化内存
	m["key"] = 1              // ✅
	fmt.Println(m)

	// fixme 所以new也不能用
	//m2 := new(map[string]int) // fixme 这里其实开辟了一个内存并默认零值（为nil）
	//(*m2)["first"] = 1        // fixme 但是呢底层是nil，真实的装数据哈希表根本就没有被创建出来，所以这里会panic
	//fmt.Println(*m2)

	// fixme make 不仅会分配内存，还同时会把底层结构构建
	s := make([]string, 4) // 直接通过make去创建slice其实也方便
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	s[3] = "d"
	fmt.Println(s)

	time0 := "2026-02-02"
	time1, _ := time.Parse("2006-01-02", time0) // fixme Parse函数是把对应的时间数据转换成time.Time对象
	fmt.Printf("%T", time1)

	start, _ := time.Parse("2006-01-02", "2026-02-02")
	end, _ := time.Parse("2006-01-02", "2026-03-12")
	diff := end.Sub(start)              // 得到 Duration (纳秒)
	totalDays := int(diff.Hours() / 24) // 转为天数
	fmt.Printf("totalDays=%d\n", totalDays)

	// 返回当前时间的time类型数据
	time4 := time.Now()
	year := time4.Year()
	month := time4.Month()
	day := time4.Day()
	fmt.Println(year, month, day)

}
