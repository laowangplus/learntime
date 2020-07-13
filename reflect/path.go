package main

import (
	"fmt"
	"reflect"
	"sync"
)

type TestObject struct {
}

func (o TestObject) Test() string {
	return "test1"
}

var (
	lock = sync.Mutex{}
)

func PatchInstanceMethod(target reflect.Type, methodName string, replacement interface{}) {
	m, ok := target.MethodByName(methodName)
	if !ok {
		panic(fmt.Sprintf("unknown method %s", methodName))
	}
	r := reflect.ValueOf(replacement)
	fmt.Println(m.Func, r)
}

func main() {
	PatchInstanceMethod(reflect.TypeOf(&TestObject{}), "Test", func() string {
		return "test2"
	})
	object := &TestObject{}
	fmt.Printf(object.Test())
}
