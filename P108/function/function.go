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
