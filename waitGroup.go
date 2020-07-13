package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()
			fmt.Println("go:", id)
		}(i)
	}

	fmt.Println("wait...")
	wg.Wait()
	fmt.Println("Done")
}
