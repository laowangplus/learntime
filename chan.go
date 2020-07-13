package main

import (
	"fmt"
	"time"
)

func main() {
	a := make([]int, 3)
	for i, _ := range a {
		//s := i
		go func() {
			fmt.Println(&i)
		}()
	}
	time.Sleep(time.Second)
}
