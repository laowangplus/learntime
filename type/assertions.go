package main

import "fmt"

type Err struct {
	S string
}

func (e *Err) Error() string {
	return e.S
}

func GetErr() error {
	return &Err{S: "abs"}
}

func GetError() error {
	return fmt.Errorf("nick")
}

func main() {
	e := GetError()
	i, ok := e.(*Err)

	if ok {
		fmt.Printf("ok, i = %+v", i)
	} else {
		fmt.Printf("fail")
	}
}

// 打印fail
