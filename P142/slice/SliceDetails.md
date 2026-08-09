### 切片注意事项和细节说明
1. 切片初始化时 var slice = array【standIndex,endIndex】
    说明：从array数组的standIndex开始取，取到endIndex（左闭右开）

2. 切片初始化时仍然不能越界，范围在 【0~len】之间，但是可以动态增长
   - var slice = arr【0,end】可以简写成 slice = arr【:end】
   - var slice = arr【start,end】 可以简写成 slice = arr【start:】
   - var slice = arr【0，len】 可以简写为 slice = arr【:】

3. cap是一个内置函数，用于统计切片的容量，即最大可以存放多少元素
4. 切片定义完之后还不能使用，因为本身时空的，需要引用到一个数组或者make一个空间给它使用 或者使用append加元素进去也会创建空间
5. 切片可以继续切片
6. 使用append内置函数可以对切片进行动态增加
    ```
    切片操作底层分析
    - 切片append的本质就是对数组扩容
    - go底层会创建新的数组newArr安防扩容后大小
    - 将slice原本的元素拷贝到newArr
    - slice重新引用newArr
    - 这里所说的newArr是不可见的，底层维护的
    ```
7. Go 语言 copy() 的“最短长度”原理
8. copy 操作是“破坏性”的，它会直接覆盖（修改）目标切片里原来的数据

### My conclusion
- make([]type , len , cap)
    ```
    这个函数，如果不指定cap，那么cap默认等于len
    ```
- var slice = []int {1,2,3}
    ```
    如果这样进行初始化，那么这个切片的容量一定是等于我们指定的len
    ```
- 切片的拷贝使用的是copy内置函数完成的拷贝
    ```go
    copy(par1 ,par2):par1 和 par2 都是切片类型才可以进行拷贝
    // 把 后面的 赋值 给前面的
    // Go 在设计 copy 函数时，底层自带了长度的安全保护机制
    // 核心规则是：按元素个数拷贝，且只拷贝两个切片长度（len）中的最小值
    
    func main(){
        var a = []int{1,2,3,4,6}
        var slice = make(slice,10)  // 这里初始是10个0
        fmt.Println(slice)
  
        copy(slice,a)  // 这里做的是把a里面的内容拷贝到slice对应的位置
        fmt.Println(slice)
        // 所以输出是：[1,2,3,4,5,0,0,0,0,0]
  }
  
    ```
- ### Go 语言 copy() 的“最短长度”原理
    
    `copy(dst, src)` 在复制前，会比较目标切片和源切片的长度，**取较小的那个数字（`min(len(dst), len(src))`）决定复制的元素个数**。
    
    #### 情况 1：目标切片比源切片“短”
    * **场景**：源 `a` 是 `[1, 2, 3, 4, 5]`（长 5），目标 `slice` 是 `[0]`（长 1）
      * **最短**：目标切片（长度 1）
      * **结果**：只复制 1 个元素。把 `1` 填入目标，装满即停。`slice` 变成 `[1]`。
    
    #### 情况 2：源切片比目标切片“短”
    * **场景**：源 `a` 是 `[1, 2]`（长 2），目标 `slice` 是 `[0, 0, 0, 0, 0]`（长 5）
      * **最短**：源切片（长度 2）
      * **结果**：只复制 2 个元素。把 `1` 和 `2` 填入目标的前两个位置，剩余保持原样。`slice` 变成 `[1, 2, 0, 0, 0]`。


- ### copy 操作是“破坏性”的，它会直接覆盖（修改）目标切片里原来的数据。
    ```go
    func test1() {
        var a = []int{1, 2, 3, 4, 5}
        slice = make([]int, 1)
        slice[0] = 520
        fmt.Println(slice) 
        copy(slice, a)  // 这里值是会别覆盖的
        copy(a, slice)
        fmt.Println(slice)
        fmt.Println(a)
    }
  // 输出结果是 1 12345
    ```