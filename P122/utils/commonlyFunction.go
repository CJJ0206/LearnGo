package utils

/*
字符串常用的系统函数（查手册有）

1. len(str)  （内置函数 builtin）
	func len(v Type) int
	The len built-in function returns the length of v, according to its type:
		Array: the number of elements in v.
		Pointer to array: the number of elements in *v (even if v is nil).
		Slice, or map: the number of elements in v; if v is nil, len(v) is zero.
		String: the number of bytes in v. （按字节）
		Channel: the number of elements queued (unread) in the channel buffer; if v is nil, len(v) is zero.

2. r := []rune(str)  // 将字符串强行转换成一个 rune 切片，处理中文乱码问题

3. strconv.Atoi()  // ASCII to Int
	func Atoi(s string) (int, error)  // Atoi is equivalent to ParseInt(s, 10, 0), converted to type int.
	// 他只能处理 数字 字符串， 对于非数字的字符串会报错，这个特性可以用来做一些标准的校验

4. strconv.Itoa()  // 整数转字符串
 	// 这是一定可以成功的

5. []byte  // 字符串转byte
	输出的是对应字母的ASCII码

6. []byte 转 字符串

7. strings.Contains("字符串","子字符串") (bool)
	查找子串是否在指定的字符串中

8. strings.Count("hello earth","e") (int)
	统计字符串中有几个指定的子串

9. strings.EqualFold("","") (bool)
	不区分大小写的字符串比较

10. strings.Index("NLT_abc","abc")
	返回子串在字符串第一次出现的index，如果没有就返回 -1

11. strings.LsatIndex(""，"")
	查找子字符串在字符串中最后出现的位置

12. strings.Replace(sting, old, new, n)
	string 可以是一个变量或者常量
	用指定的子串 new 替换 子串 old 的内容，n表示希望替换几处。-1 表示全部替换

13. strings.Split("","")
	按照指定的某个字符为分割标识符，将一个字符串拆分为字符串数组

14. strings.ToLower() / ToUpper
	字符串的大小写转换

15. strings.TrimSpace("string")
	将字符串两边空格去掉

16. strings.Trim("","")
	去掉左右两边我们指定的字符

17. strings.TrimLeft("","")
	去掉字符串左边的指定字符

17. strings.TrimRight("","")
	去掉字符串右边的指定字符




*/
