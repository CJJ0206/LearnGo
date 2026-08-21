### 接口注意事项
#### 1. 接口本身不能创建实例，但是可以指向一个实现了该接口的自定义类型的变量（实例）
#### 2. 接口中所有的方法都没有方法体。即均为没有实现的方法
#### 3. 在go中，一个自定义类型只有把某个接口的所有方法都实现，我们才说他实现了这个接口
#### 4. 在golang中，一个自定义类型只有实现了某个接口，才能将该自定义类型的实例赋给接口类型
#### 5. 只要是自定义的数据类型，就可以实现接口，不仅仅是结构体类型
   - ```go
     type a interface {
         Say()
         Leave()
     }
     type integer int
     func (i integer) Say() {
         fmt.Println("here is integer")
     }
     func (i integer) Leave() {
         fmt.Println("here is integer")
     }
     func main() {
         var i integer = 10
         var b a = i  // 可以看到能够把一个实现了接口的变量赋值
         b.Say()
     }
     ```
#### 6. 一个自定义类型可以实现多个接口
   - ```go
     type A interface {
        Say()
     }
     type B interface {
        Hello()
     }
     type monster struct{ }
     func (m *monster) Say(){
        fmt.Printl("say")
     }
     func (m *monster) Hello(){
        fmt.Printl("hello")
     }
     // 此时monster同时实现了两个接口
     var mon monster
     var a A = &mon  // 很灵活，通过各自接口名字直接就可以用
     var b B = &mon
     a.Say()
     b.Hello()
     ```
#### 7. go的接口中不能有任何变量
#### 8. 一个接口（比如A接口）可以继承多个别的接口（比如B\C接口），这时如果要实现A接口，也必须将B\C接口的方法全部实现
   - ```go
     type A interface {
        test01()
     }
     type B interface {
        test02()
     }
     type C interface {
        A  // 继承别的接口
        B
        test03()
     }
     type stu Student{ }
     // fixme 这里不需要知道test几是谁的方法，只有全部实现了才算实现了C接口，整体才能用
     func (s stu) test01(){
     }
     func (s stu) test02(){
     }
     func (s stu) test03(){
     }
     // 此时可以通过C的变量任意去调A\B的方法
     ```
#### 9. interface类型默认是一个指针（引用类型），如果没有对interface初始化就使用，会返回nil
#### 10. 空接口interface{ }没有任何方法，所以所有类型都实现空接口

#### error
```go
type Ainierface interfece {
	test1()
	test2()
}
type Binterface interface {
	test1()
	test3()
}
type Cinterface interface {
	Ainterface
	Binterface  
	// fixme 这里会直接报错，因为这两个接口有完全相同的方法名，会导致后续的混淆
	// fixme 即 Cinterface 里面现在有两个完全一样的 test1()函数，那肯定会报错
}

```