package main

import (
	"fmt"
	"time"
)

func sliceTest(s []int) []int {
	s = append(s, 1)
	fmt.Printf("%p\n", &s)
	return s
}

func main() {
	s := make([]int, 0, 5)
	for {
		s = append(s, 1)
		fmt.Printf("%p, %v, %v\n", &s, len(s), cap(s))
		time.Sleep(100000000)
	}
}
