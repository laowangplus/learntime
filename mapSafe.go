package main

import (
	"sync"

	"fmt"
)

func main() {

	var sm sync.Map

	go func() {
		for {
			sm.Store(1, 1)
		}
	}()

	go func() {
		for {
			fmt.Println(sm.Load(1))
		}
	}()

	select {}
}
