package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	list := make([]int, 0)
	for _, value := range nums {
		if len(list) < 4 {
			list = append(list, value)
		} else {
			flag_key := 0
			min := math.MaxInt32
			for k, v := range list {
				if min > v {
					min = v
					flag_key = k
				}
			}
			if min < value {
				list[flag_key] = value
			}
		}
	}
	min := math.MaxInt32
	for _, v := range list {
		if min > v {
			min = v
		}
	}
	fmt.Println(min, list)
}
