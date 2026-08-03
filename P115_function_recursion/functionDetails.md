### 函数注意事项和细节讨论
1. 函数的形参列表可以是多个，返回值列表也可以是多个  
   ```go
    func SubSum(a int, b int) (int, int) {
        sum := a + b
        sub := a - b
        return sum, sub  // 形参和返回值列表均为多个例子
    }
   ```
2. 形参列表和返回值列表的数据类型可以是值类型和引用类型
   ```go
    func UpdatePlayerScore(p *Player, bonus int) (bool, *int) {
	      if p == nil {   //遇到指针，第一步最好先做判空处理，防止引发 panic
        }
     return false, nil
	    }
	// 1. 操作指针类型：直接修改了 p 指向的内存中的数据
	   p.Score += bonus
	// 2. 操作值类型：这里的 bonus 是个副本，修改它不会影响外面的变量
	   bonus = 0
	// 返回一个值类型 (true) 和一个指针类型 (&p.Score，即分数的地址)
	   return true, &p.Score
   ```

3. 首字母大写 public ,小写 private
   ```
    如果是外部的包导入后要用其函数，只能使用首字母大写的（公有）
   ```
   
4. 函数中的变量是局部的，函数外不生效
   ```go
   i := 10
   func test(i int) int {
      i++
      fmt.Println(i)
   }   
      
   test(i)
   fmt.Print(i)  // 输出结果是 11  10 ，所以函数内的变量是不影响外部真实变量的
   ```
5. 基本数据类型和数组默认都是进行的值传递，即进行值的拷贝，不影响原本地址里的变量
   ```go
   i := 10
   func test(p *int){
       *p++
   }
   // 这样是会修改外部的真实值的
   ```
6. 如果希望函数内的操作可以修改外部的值，那么传外部变量的地址即可·
   ```
   同上
   ```
7. go 的函数不支持重载

8.   <span style="color: orange;">go 中，函数是一种数据类型，可以赋值给一个变量，则该变量就是一个函数类型的变量了，通过这个变量可以调用函数</span>
   ```go
   func getsum(n1 int , n2 int) int {
        return n1 + n2
   }
   
   func main(){
        a := getsum  // 这里就相当于别的语言里的创建一个对象
        fmt.Println("a的类型为%T,getsum类型为%T",a,getsum)
        // 输出类型均为 func(int int)int 所以函数也是一种数据类型
    
        res := a(10,40)  // 然后 a 就是这个函数对象，直接调用函数
        fmt.Println(res)
   }
   ```
9. <span style="color: orange;">在 go 中函数也是一种数据类型，所以函数可以作为形参，并且调用！</span>
   ```go
   // 函数也是可以作为形参的，只要传值传对就行
   func myfun(funvar func(int int) int , num1 int ,num2 int){
        return funvar(num1,num2)  // 这里直接通过变量来调用函数了
   }
   
   // 这里我们直接传上面的 getsum 函数
   res2 := myfun(getsum , 50 , 60)
   // 作为形参 以及赋值的时候是没有括号的 是作为一个对象在传
   ```
10. <span style="color: orange;">支持对函数返回值命名</span>
   ```go
   func getsumandsub(n1 int , n2 int) (int int ){
        // 这种写法就要求我们搞清楚对应输出的顺序
    }
    // 通过直接在返回值类型列表指定返回名称
    func better(n1 int , n2 int)(sum int , sub int){
        sum = n1 + n2  // 这里也可以不用定义sum参数了
        sub = n1 - n2 
        return  // 返回也不用指定
    }
   ```
11. 使用 _ 忽略特定值
12.   <span style="color: green;">go 支持多个可变参数</span>
   ```go
   // 函数可以接收多个参数
   // 可变参数 args 必须放在形参列表的最后 
   func sum(n1 int, args ...int) int {
	total := n1 
	for i := 0; i < len(args); i++ {
		total += args[i] 
	    }
	return total
    }
    
    // 上面的for循环等价于这里，现在其实很少用类似c的古老for了
    //for _, v := range args {
           total += v
       }
   ```
13. 为了简化数据类型定义，go支持自定义数据类型
   ```go
   type myint int // 通过type这个关键字给int起别名
   // 但是 go 认为他们是不同数据类型，所以最后还要强转
    
   var num1 mytint
   num1 = 10
   fmt.Println(num1)


   // 例子
   type myfuntype func(int, int) int  
   // 此时，myfuntype 就代表输入输出为指定类型的函数类型

   func myfun(funvar func(int int) int , num1 int ,num2 int){
        return funvar(num1,num2)  // 这里直接通过变量来调用函数了
   }
   // 这里则可以直接用这个定义来代替，减少代码
   func myfun(funvar myfuntype , num1 int ,num2 int){
	   return funvar(num1,num2)  // 这里直接通过变量来调用函数了
   }
   ```

