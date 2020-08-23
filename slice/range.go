package main

import "fmt"

func main() {
	data := []*struct{ num int }{{1}, {2}, {3}}
	for _, v := range data {
		v.num *= 10 // 直接使用指针更新
	}
	fmt.Println(data[0], data[1], data[2]) // &{10} &{20} &{30}
}
