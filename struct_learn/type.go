package main

import "fmt"

type Status interface {
	Value() uint8
}

type status bool

func (s status) Value() uint8 {
	if s {
		return 1
	} else {
		return 0
	}
}

func zh(s Status) {
	fmt.Println(s.Value())
}

func main() {
	var s status
	v1 := false
	s = status(v1)
	zh(s)
	print(s.Value())
}
