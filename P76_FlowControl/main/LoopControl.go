package main

import "fmt"

// 嵌套分支不要过深，深了就很复杂了，建议控制在3层

func loopTest1(score float32, gender string) {
	if score < 8.0 {
		if gender == "man" {
			fmt.Println("进入男子组决赛。")
		} else if gender == "woman" {
			fmt.Println("进入女子组决赛。")
		}
	}
}
