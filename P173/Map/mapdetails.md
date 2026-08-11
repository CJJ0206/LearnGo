### map的使用细节
1. map是引用类型，遵守引用类型传递规则，在一个函数接收map并修改后，map的原值是会改变的
2. map的容量到达后，再向map增加元素，会自动扩容(**_但是切片只能用append_**)，并不会发生panic,也就是说**map能动态的增长键值对**
3. **_map的value也经常使用struct类型_**，更适合管理复杂的数据，比如value是一个studet结构体
    ```go
    type Stu struct{
        Name string
        Age int
        Grade float64
    }
    func main(){
        students := make(map[string]Stu,10)
        stu1 := Stu{"tom",10,99.4}
        stu2 := Stu{"jerry",18,95.4}
        stu3 := Stu{"cjj",23,99}
        students["no1"] = stu1
        students["no2"] = stu2
        students["no3"] = stu3
        fmt.Println(students)
    
        for k,v := range students{
            fmt.Println("编号",k)
            fmt.Println("名字",v.Name)
            fmt.Println("年龄",v.Age)
            fmt.Println("成绩",v.Grade)
		}
   }
    ```
4.