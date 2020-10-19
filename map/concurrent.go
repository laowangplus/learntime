package main

import (
	"fmt"
	"sync"
	"time"
)

var concurrent map[int]*sync.Pool

type T struct {
	name string
}

func main() {
	concurrent = make(map[int]*sync.Pool)
	t := T{name: "laowang"}
	concurrent[1] = &sync.Pool{
		New: func() interface{} {
			return t
		},
	}
	go func() {
		fmt.Println(concurrent[1].Get())
	}()

	go func() {
		fmt.Println(concurrent[1].Get())
		t := T{name: "laowang2"}
		concurrent[1].Put(t)
	}()

	time.Sleep(1 * time.Second)
	fmt.Println(concurrent[1].Get())

}
