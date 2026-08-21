package main

import (
	"fmt"
	"slices"
)

func main() {
	/* slices.Clone ：func Clone(s S) S
	所属包：slices
	形参：s S : 待克隆的目标切片，可以是任意数据类型的切片
	返回值类型：S：新分配的切片副本，若传入的 s 为 nil,则返回 nil
	核心作用：对切片进行浅拷贝，分配一块全新的底层数组
	底层：
		判空，然后分配容量等于 len(s) 的全新底层数组，并通过内置的 copy() 将原切片元素复制过去。新切片的 len 和 cap 均等于原切片的 len。

	浅拷贝限制：
		如果切片元素是指针类型（如 []*User）或包含引用字段（如包含内部切片/Map 的结构体），
		slices.Clone 仅复制指针地址，指针指向的实际对象仍然是共享的。如需完全独立，需自行进行深拷贝（Deep Copy）。
	内存隔离与封装：
		当结构体对外暴露切片 Getter 方法时，推荐返回 slices.Clone(s.items)，防止外部调用方修改内部切片数据。
	*/
	original := []int{10, 20, 30}
	cloned := slices.Clone(original) // 克隆一个独立的切片副本
	cloned[0] = 999                  // 修改克隆后的切片

	// fixme 地址是不同的会新生成一块地址
	fmt.Printf("原切片地址为%p", &original) // 输出: 原切片: [10 20 30] (不受影响)
	fmt.Printf("克隆切片地址为%p", &cloned)  // 输出: 克隆切片: [999 20 30]

	// fixme 保护封装性
	cfg := NewConfig([]string{"a", "b"})

	got := cfg.Items()
	got[0] = "hacked"

	fmt.Println(got)         // [hacked b]
	fmt.Println(cfg.Items()) // [a b]

	test()

	// --------------------------------------------------------------
	/* slices.Clip ： func Clip(s S) S
	所属包：slices
	形参：s S：待裁剪容量的目标切片
	返回值类型：S：容量已被限制的新切片，若 s 为 nil 则返回 nil
	核心功能：经切片的容量 cap 裁剪值与其 len 一致，即强制 len = cap

	为什么需要 slices.Clip？
	当对切片进行子切片操作（如 sub := original[:2]）时，sub 的容量通常会延伸到 original 的末尾。
	此时如果对 sub 调用 append()，由于容量未满，它会直接覆盖写 original 中原有的后续元素！
	使用 slices.Clip(sub) 之后，sub 的容量被锁定为当前长度，任何后续的 append() 操作都会强制触发底层数组重新扩容分配，从而彻底避免意外覆盖原切片数据。
	*/
	raw := []string{"A", "B", "C", "D"}
	// 1. 未使用 Clip 的危险情况：
	sub1 := raw[:2] // len=2, cap=4 (仍能看到底层的 C 和 D 位置)
	_ = append(sub1, "X")
	fmt.Println("未使用 Clip 时 append 导致原切片被覆盖:", raw)
	// 输出: [A B X D] （注意：原先的 "C" 被覆盖为 "X"！）

	// 恢复数据
	raw = []string{"A", "B", "C", "D"}
	// 2. 使用 Clip 的安全情况：
	sub2 := slices.Clip(raw[:2]) // len=2, cap=2 (容量锁定)
	sub2 = append(sub2, "X")     // 由于容量已满，append 会自动开辟新数组

	fmt.Println("使用 Clip 后原切片保持安全:", raw) // 输出: [A B C D] (未被覆盖)
	fmt.Println("sub2 得到全新独立数据:", sub2)   // 输出: [A B X]
}
