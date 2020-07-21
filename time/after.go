package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go func() {
		time.Sleep(time.Second * 2)

		ch <- "result"
	}()

	select {
	case res := <-ch:
		fmt.Println(res)
	case <-time.After(time.Second * 2): //time.After()表示time.Duration长的时候后返回一条time.Time类型的通道消息。
		fmt.Println("timeout")
	}

	//<-time.After(time.Second * 2)
}
