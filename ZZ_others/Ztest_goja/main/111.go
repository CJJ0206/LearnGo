package main

import (
	"fmt"

	"github.com/dop251/goja"
)

func main() {
	vm := goja.New() // 创建一个新的虚拟机实例

	// 下面这个用反引号包裹的、纯纯的 JavaScript 语法代码块，就是 JS 代码
	jsCode := `
		const a = 10;
		const b = 20;
		const multiply = (x, y) => x * y;
		
		// 最后执行函数的返回值，会被 Go 捕获
		multiply(a, b); 
	`

	// 把这段 JS 代码扔给 Goja 引擎去跑
	result, err := vm.RunString(jsCode) // 给这个虚拟机器引擎去跑
	if err != nil {
		panic(err)
	}

	num := result.Export().(int64)
	fmt.Println("JS 算出来的结果是:", num) // 输出 200
}
