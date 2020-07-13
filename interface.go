package main

import (
	"fmt"
)

type Phone interface {
	call()
}

type iPhone struct {
	name string
}

func (phone iPhone) call() {
	fmt.Println("iPhone")
}

func test(p Phone) {
	p.call()
}

func main() {
	phone := iPhone{}
	test(phone)
}
