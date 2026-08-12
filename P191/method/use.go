package main

import "fmt"

func (p Person) calculate() {
	sum := 0
	for i := 1; i <= 1000; i++ {
		sum += i
	}
	fmt.Println("1~1000的和为：", sum)
}

func (p Person) calculate2(n int) {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	fmt.Printf("1~%d的和为：%d", n, sum)
	fmt.Println()
}

func (p Person) getSum(a, b int) {
	fmt.Println("a、b的和为：", a+b)
}

func test1() {
	p1 := Person{"cjj"}
	p1.speak()

	p1.calculate()
	p1.calculate2(20)
	p1.getSum(10, 20)
}
