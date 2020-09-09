package main

import (
	"os"
	"runtime/trace"
	"time"
)

func main() {
	f, err := os.Create("./trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	go func() {
		for true {
			time.Sleep(time.Second)
		}
	}()

	time.Sleep(3 * time.Second)

	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()
	// Your program here
}
