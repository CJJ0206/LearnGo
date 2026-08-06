package utils

// 8.4

import "fmt"

// 在函数中，我们通常需要创建资源（如：数据库连接、锁等），为了在函数执行结束后，及时释放资源，提出了defer(延时机制)

func Sum(n1 int, n2 int) int {
	// FIXME 当执行到defer时，系统会先不执行语句而是把语句压如一个独立的栈
	// FIXME 当函数执行完毕后，再从栈中，按照先入后出方式出栈执行
	defer fmt.Println("ok1 n1 = ", n1)
	defer fmt.Println("ok2 n2 = ", n2)

	res := n1 + n2 // 所以这个函数先执行的时这句
	fmt.Println("ok3 res =", res)
	return res
	// 这个函数的输出顺序是倒着的
}

// Sum2 FIXME 在defer将语句放入栈时，也会将相关的值拷贝同时入栈
func Sum2(n1 int, n2 int) int {
	// defer 的拷贝只是做一个快照，把当时的值拷贝了一下（值传递）
	defer fmt.Println("ok1 n1 = ", n1) // 输出 10
	defer fmt.Println("ok2 n2 = ", n2) // 输出 20

	n1++
	n2++

	// 但是传进来的值还是顺利进行的
	res := n1 + n2
	fmt.Println("ok3 res =", res) // 输出 32
	return res
}
