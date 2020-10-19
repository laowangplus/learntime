package main

import "fmt"

type stack struct {
	vals []int
}

func (s *stack) push(value int)  {
	s.vals = append(s.vals, value)
}

func (s *stack) pop() int {
	l := len(s.vals)
	value := s.vals[l-1]
	s.vals = s.vals[:l-1]
	return value
}

func main()  {
	//var s = &stack{
	//	vals:[]int{},
	//}
	//s.push(1)
	//s.push(2)
	//s.pop()
	//
	//fmt.Println(s.vals)

	m := map[int]int{
		1:1,
		2:2,
	}
	a := func(arr map[int]int) {
		arr[1] = 3
	}

	a(m)
	fmt.Println(m)
}

