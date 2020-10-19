package main

import "fmt"

type Engine struct {
	queue chan interface{}
}

// 推送请求到队列
func (e *Engine) Push(req interface{}) {
	e.queue <- req
}

// 开启服务协程
func (e *Engine) Serve() {
	for {
		req := <-e.queue
		// do some thing
		fmt.Println(req)
		_ = req
	}
}

func main() {
	var e = make(chan interface{})
	list := &Engine{
		queue: e,
	}
	list.Push(1)
	list.Serve()
	select {}
}
