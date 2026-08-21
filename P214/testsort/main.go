package main

import (
	"cmp"
	"fmt"
	"math/rand"
	"slices"
)

type SortInter interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}
type device struct {
	name        string
	id          int
	temperature float32
}
type deviceSlice []device // deviceSlice 类型，其实就是给[]device的别名，为了方便绑定方法
func (d deviceSlice) Len() int {
	return len(d)
}
func (d deviceSlice) Less(i, j int) bool {
	// 第一优先级：id
	if d[i].id != d[j].id {
		return d[i].id < d[j].id
	}
	// 第二优先级：temperature
	return d[i].temperature < d[j].temperature
}
func (d deviceSlice) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

func main() {
	var devs deviceSlice
	for range 10 { // fixme for循环遍历创建设备信息，然后append就行
		dev := device{
			name:        fmt.Sprintf("设备%d", rand.Intn(20)),
			id:          rand.Intn(3),
			temperature: rand.Float32()*60 + 20,
		}
		devs = append(devs, dev)
	}
	//sort.Sort(devs)
	//for _, dev := range devs {
	//	fmt.Printf("name:%v\t ID:%d\t temp:%f\t \n", dev.name, dev.id, dev.temperature)
	//} // 输出实现先按照ID排序然后按照温度升序

	fmt.Println(devs)
	slices.SortFunc(devs, func(a, b device) int {
		if a.temperature != b.temperature {
			return cmp.Compare(a.temperature, b.temperature) // 因为温度没有一样的所以直接全部执行的温度升序
		}
		return cmp.Compare(a.id, b.id) // 写反了和上面
	})
	for _, dev := range devs {
		fmt.Printf("name:%v\t ID:%d\t temp:%f\t \n", dev.name, dev.id, dev.temperature)
	}

}
