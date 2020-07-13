package main

import (
	"fmt"
	"math"
	"sync"
)

func count() {
	x := 0
	for i := 0; i < math.MaxInt32; i++ {
		x += i
	}
	fmt.Println(x)
}

func test1(n int) {
	for i := 0; i < n; i++ {
		count()
	}
}

func test2(n int) {
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			count()
			wg.Done()
		}()
	}

	wg.Wait()

}

func main() {
	//n := runtime.GOMAXPROCS(0)
	//test1(n)
	//test2(n)
}
