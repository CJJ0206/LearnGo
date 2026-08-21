package main

import (
	"cmp"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"path/filepath"
	"slices"
)

type User struct {
	Name string
	Age  int
}

func main() {
	/*
		slices.SortFunc
		所属包：；slice
		核心功能：根据自定义的比较函数对切片进行原地排序
		参数与返回值：
			s ~【】E：待排序的切片
			cmp func(a,b E) int : 比较函数。fixme 若 a < b 返回负数，若 a == b 返回 0，若 a > b 返回正数，底层会根据这个返回的数字做排序
		适用场景：结构体切片按特定字段排序、复杂对象的自定义逻辑排序。
	*/
	users := []User{ // 一个结构体切片
		{"Alice", 30},
		{"Bob", 22},
		{"Charlie", 25},
	}
	slices.SortFunc(users,
		func(a, b User) int { // 按年龄升序排序
			return cmp.Compare(a.Age, b.Age)
		})

	// fixme 传进去的可遍历元素底层会自动遍历对比
	slices.SortFunc(users,
		func(a, b User) int { // 先比较年龄
			if n := cmp.Compare(a.Age, b.Age); n != 0 {

				return n // 如果年龄不相等，直接返回正负结果
			}
			return cmp.Compare(a.Name, b.Name) // 年龄相同时，再比较姓名
		}) // fixme 这里都是第二个形参是一个完整的匿名函数

	fmt.Println("按年龄升序排序结果:", users)

	// -----------------------------------------------------------------
	/*
			filepath.WalkDir
			所属包：path/filepath
			核心功能：递归遍历指定目录树下的所有文件和子目录。相比旧版的filepath.Walk，WalkDir使用fs.DirEntry延迟获取文件元数据，显著减少了文件系统I/O开销，遍历效率更高
			参数与返回值：
				root string：起始遍历的根目录路径。
				fn WalkDirFunc：回调函数 func(path string, d fs.DirEntry, err error) error。
				返回 error：遍历过程中的错误信息。
			适用场景：高效扫描磁盘文件夹、批量处理指定扩展名的文件、统计目录占用空间。

		filepath.WalkDir： func WalkDir(root string, fn fs.WalkDirFunc) error
		fs.WalkDirFunc 到底是什么：
			官方 io/fs 包里的源码定义：type WalkDirFunc func(path string, d DirEntry, err error) error
		fn fs.WalkDirFunc：所以在使用时我们直接固定填这个形参就行了，底层已经帮我们固定为上面这个函数类型了
	*/
	rootDir := "." // 遍历当前目录
	err := filepath.WalkDir(rootDir,
		func(path string, d fs.DirEntry, err error) error { // 第二个形参，传的是walkDir碰到文件时的处理逻辑
			if err != nil {
				return err
			}
			if !d.IsDir() && filepath.Ext(path) == ".go" { // 只打印 Go 源码文件路径
				fmt.Println("找到 Go 文件:", path)
			}
			return nil
		}) //fixme 注意！这个匿名函数，传的是整个函数的定义

	if err != nil {
		fmt.Println("遍历出错:", err)
	}

}

/*
http.MaxBytesReader
所属包：net/http
核心功能：对 io.ReadCloser（如 HTTP 请求体 r.Body）包装一个最大读取字节数限制。若客户端发送的数据超出限制后续读取会直接返回错误，避免服务器因处理过大 Payload 而耗尽内存
参数与返回值：

	w ResponseWriter：HTTP 响应写入器（传入后当超限时自动设置链接关闭标记）
	r io.ReadCloser：待限制的读取流（通常为 r.Body）
	n int64：允许读取的最大字节数
	返回 io.ReadCloser：包含限制逻辑的包装流

适用场景：防护 API 接口免受超大 Body 攻击（DoS 防御）、限制文件上传大小
*/
func uploadHandler(w http.ResponseWriter, r *http.Request) {
	// 限制请求体最大为 1 MB (1024 * 1024 字节)
	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)
	_, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "请求体超出 1MB 限制或读取失败", http.StatusRequestEntityTooLarge)
		return
	}
	w.Write([]byte("读取成功"))
}
