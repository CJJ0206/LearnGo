package main

import "fmt"

func main() {
	s1 := []string{"cjj", "aaa", "bbb"}
	fmt.Printf("%p\n", s1)  // 这是s1底层指向的地址
	fmt.Printf("%p\n", &s1) // 这是s1自己的地址

	s1 = append(s1, "ccc", "ddd")
	fmt.Printf("%p\n", s1)
	fmt.Printf("%p\n", &s1) // fixme 扩容后自身的地址是没变的，但是指向的地址改变了

	s2 := []string{"sb", "2b"}
	copy(s1, s2) // （src dst）  fixme 源序列长的话，则只会覆盖前面几个元素
	fmt.Println(s1)

	str := "hello"
	// str[0] = 'a'  // 不允许直接修改string里的值
	// todo 但是可以通过把string转成slice去修改
	sli := []byte(str)
	sli[0] = 'a' // 可以通过切片直接去修改原string
	fmt.Printf("%c \n", sli)

	fmt.Println("----------------------------------")
	result := fbn(6)
	fmt.Println(result)
}

func fbn(n int) []int64 {
	slice := make([]int64, n)
	slice[0] = 1
	slice[1] = 1
	for i := 2; i < n; i++ {
		slice[i] = slice[i-1] + slice[i-2]
	}
	return slice
}
