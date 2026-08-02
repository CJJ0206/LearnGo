package function

import "fmt"

// 外部的函数名必须要大写才是全局的

func Calculate(a float64, b float64, operate byte) float64 { // 这里最后指定下返回值数据类型
	var result float64
	switch operate {
	case '+':
		result = a + b
	case '-':
		result = a - b
	case '*':
		result = a * b
	case '/':
		result = a / b
	default:
		fmt.Println("操作符未收录")
	}
	return result
}

func Test(n1 int) {
	n1 += 1
	fmt.Printf("Test函数中的 n1 = %d ", n1)
}

// Sum 返回值类型列表
func Sum(n1 int, n2 int) int {
	return n1 + n2
} // 有返回值的函数，会把计算出的最终结果交给调用者

// SumSub 双返回值类型演示
func SumSub(a int, b int) (int, int) { // 可以看到这里写的是双返回值类型
	sum := a + b
	sub := a - b
	return sum, sub
} // go 返回多个值不用什么数组啥的，直接标注再返回就行
