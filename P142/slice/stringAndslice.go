package main

import "fmt"

// 8.9

/*
string 底层是一个byte数组，因此string是可以进行切片处理的

*/

func sliceString() {
	str := "hello@cjj" // 现在我们可以用切片把指定内容取出来
	slice := str[5:]   // 这个也是左开右闭的
	fmt.Println(slice)
	// 但是对然可以这样输出指定内容，但是string本身是不可变的，若强行修改是会报错的
	// 若果需要修改字符串，如下
	// 我们要修改string中的char怎么做呢
	arr := []byte(str)
	arr[0] = 'a'
	// arr[1] = '你'  // 中文的类型是不符合一个byte的数据类型的，所以会直接报错
	str = string(arr)
	fmt.Println(str) // 这种方法虽然可以，但是如果要改中文还是有问题的

	// 通过rune就可以处理
	// fixme 因为rune是按字符处理的，所以兼容汉字，他把不论是英文还是中文还是数字都当作rune
	arr2 := []rune(str)
	arr2[0] = '我'
	fmt.Println(string(arr2))

}
