package main

import (
	"cmp"
	"fmt"
	"math/rand"
	"os"
	"slices"
)

func testOr() {
	userInputPort := "8081"
	envPort := os.Getenv("APP_PORT")
	defaultPort := "8080"

	finalPort := cmp.Or(userInputPort, defaultPort, envPort)
	fmt.Println(finalPort)
}

func testSliceDelete() {
	goods := []string{"pear", "apple", "orange", "iphone"}
	slices.Delete(goods, 2, 3) // 左闭右开，所以这只删了orange
	fmt.Println(goods)
	// todo 这个还是很简单的
}

type Goods struct {
	id    int
	price float64
}

var goodSlice []Goods

func testDeleteFunc() {
	//goods := []string{"pear", "apple", "orange", "iphone", "apple"}
	//slices.DeleteFunc(goods, func(g string) bool {
	//	if g == "apple" {
	//		return true
	//	}
	//	return false
	//})
	//fmt.Println(goods)

	for range 5 {
		goodSli := Goods{
			id:    rand.Intn(5),
			price: rand.Float64()*20 + 10,
		}
		goodSlice = append(goodSlice, goodSli)
	}
	for _, good := range goodSlice {
		fmt.Println(good)
	}
	goodSlice = slices.DeleteFunc(goodSlice, func(g Goods) bool { // fixme 这里一定要把原值赋回去，不然输出会保留删除后的{0，0}
		return g.id == 0 // fixme 这里是只要return了true就删除
	})
	fmt.Println("new slices")
	for _, good := range goodSlice {
		fmt.Println(good)
	}
}
