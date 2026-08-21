package main

import (
	"cmp"
	"fmt"
	"slices"
)

type Account struct {
	ID   int
	Name string
}

func main() {
	/*  slices.BinarySearch ： func BinarySearch[S ~[]E, E cmp.Ordered](s S, target E) (int, bool)
	形参:
		s S: 必须按照升序排好的目标切片。元素满足E（基础数值或者字符串）
		target E: 待查找的目标元素值
	返回值类型:
		int:
			若找到目标元素：返回该元素在切片中的索引位置
			若未找到：返回该元素为了保持切片有序应该被插值的索引位置
		bool:
			是否找到了目标元素（true存在，false不存在）
	核心功能：在有序切片执中执行经典的二分查找，时间复杂度O（log n）
	设计亮点（插入点索引）：
		在查找失败时（found == false），返回的 i 并非无效值（如 -1），而是该元素插入后仍能保持切片有序的精确下标。配合 slices.Insert 可以直接实现“保序插入”。

	注意事项与避坑指南
		必须保证切片已排序：二分查找的前提是切片已经处于升序排列。如果在未排序切片上调用，结果未定义（可能误报未找到）。
		重复元素处理：如果切片中存在多个相同的 target，该函数返回的索引可能是其中任意一个匹配项的下标（不保证是第一个或最后一个）。
	*/

	// 必须是有序切片
	numbers := []int{10, 20, 30, 40, 50}

	// 场景一：成功查找到元素
	idx, found := slices.BinarySearch(numbers, 30)
	fmt.Printf("查找 30 -> 索引: %d, 是否存在: %v\n", idx, found) // 索引: 2, 是否存在: true

	// 场景二：未找到时利用返回值实现“保序插入”
	target := 25
	insertIdx, found := slices.BinarySearch(numbers, target)
	if !found {
		fmt.Printf("未找到 %d，建议插入位置: %d\n", target, insertIdx) // 建议插入位置: 2
		// 使用 slices.Insert 在指定位置插入元素
		numbers = slices.Insert(numbers, insertIdx, target)
		fmt.Println("插入后的有序切片:", numbers) // 输出: [10 20 25 30 40 50]
	}

	/* slices.BinarySearchFunc : func BinarySearchFunc[S ~[]E, E, T any](s S, target T, cmp func(E, T) int) (int, bool)
	形参：
		s S:已经排好序的复合结构体切片，元素类型E
		target T:待查找的目标值。注意：类型 T 不需要和 E 相同（例如在 []User 切片中直接用 int 类型的用户 ID 进行查找）。
		cmp func(elem E, target T) int：自定义比较函数：
			若 elem < target，返回负数（< 0）；
			若 elem == target，返回 0；
			若 elem > target，返回正数（> 0）。
	返回值类型：
		int：匹配到的索引位置，或建议的插入位置。
		bool：是否命中。
	核心功能与递进优势：
		极大地简化了结构体切片的二分查找。由于 target 的类型 T 与元素类型 E 解耦，你不需要为了查找而专门构造一个完整的结构体假对象，直接传目标字段（如 ID 或姓名）即可检索。

	注意事项与避坑指南
		比较函数的入参顺序：第一个参数是切片中的元素 elem，第二个参数是你传入的查找目标 target，比较时切勿颠倒前后关系。
		比较逻辑与排序规则一致：比较函数 cmp 中所使用的排序准则，必须与切片之前的排序规则完全一致。
	*/

	// 已按 ID 升序排列的账户列表
	accounts := []Account{
		{ID: 1001, Name: "Alice"},
		{ID: 1005, Name: "Bob"},
		{ID: 1008, Name: "Charlie"},
	}

	targetID := 1005

	// 直接根据 int 类型的 targetID 进行二分查找，无需构造 Account 结构体
	idx, found = slices.BinarySearchFunc(accounts, targetID, func(a Account, target int) int {
		return cmp.Compare(a.ID, target)
	})

	if found {
		fmt.Printf("成功找到 ID=%d 的账户: %+v (索引: %d)\n", targetID, accounts[idx], idx)
		// 输出: 成功找到 ID=1005 的账户: {ID:1005 Name:Bob} (索引: 1)
	} else {
		fmt.Println("账户不存在")
	}

	/*  errors.Join ：func Join(errs ...error) error
	形参说明：
		errs ...error：可变参数列表，传入需要合并的一个或多个错误（可包含 nil）。
	返回值类型：
		error：
			若所有传入的项均为 nil（或未传入参数），返回 nil。
			若存在非 nil 错误，返回一个聚合错误对象。
	核心功能：将多个错误合并为一个单一的 error。其 Error() 方法会将各个错误的描述信息用换行符 \n 连接输出。
	底层机制与优势：
		该错误聚合实现了 Go 1.20 的 Unwrap() []error 接口（多重解包）。
		与 errors.Is 和 errors.As 完全兼容：后续对其调用 errors.Is(joinedErr, targetErr) 时，Go 会自动递归遍历里面合并的所有子错误进行匹配判断。
	注意事项与避坑指南
		自动过滤 nil：函数内部会自动忽略任何 nil 错误，无需在调用前手动写 if err != nil 过滤。
		多任务/Defer 场景首选：在批量执行异步任务、并发收集错误，或在 defer 中关闭多个资源并汇总报错时，errors.Join 是标准库的最佳实践，无需再引入第三方库（如 multierr）。
	*/

	testError()

	// --------------------------------------------------------------------------------
	testSliceBinary()
	testSliceBinaryFunc()

}
