package main

import (
	"fmt"
	"sync"
)

func main() {
	////异步信道
	//ch := make(chan int)
	//
	//pool := &sync.Pool{
	//	New: func() interface{} {
	//		return 0
	//	},
	//}
	//
	//pool.Put(2)
	//pool.Put(3)
	//pool.Put(4)
	//init := pool.Get()
	//fmt.Println(init)
	//
	//go func() {
	//	init := pool.Get() //线程安全
	//	fmt.Println(init)
	//	ch <- 1
	//}()
	//
	//<-ch
	//init = pool.Get()
	//fmt.Println(init)

	//time.Sleep(100000)

	objectPool(1)
	objectPool(2)
}

var pool = &sync.Pool{
	New: func() interface{} {
		return new(object)
	},
}

type object struct {
	x int
}

func objectPool(x int) {
	ob := &object{
		x: x,
	}

	pool.Put(ob)
	init := pool.Get()
	fmt.Printf("%p\n", init)

	init.(*object).x = x + 1

	pool.Put(init)
	init = pool.Get()
	fmt.Printf("%p\n", init)
}
