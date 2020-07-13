package main

import "fmt"

func deferDemo(ch chan int) {
	ch <- 12
}

func main() {
	ch := make(chan int)
	go deferDemo(ch)
	i := <-ch
	fmt.Println(i)
}
