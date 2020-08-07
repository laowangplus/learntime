package main

import (
	"bytes"
	"fmt"
	"strings"
)

//对于string和bytes类型，且大小写不严谨的，可以使用 bytes.EqualFold() 和 strings.EqualFold() 比较

func main() {
	s1 := "ABC"
	s2 := "abc"
	fmt.Println(strings.EqualFold(s1, s2))

	s3 := []byte("ABC")
	s4 := []byte("abc")
	fmt.Println(bytes.EqualFold(s3, s4))

	var s5 []byte
	var s6 []byte = nil
	fmt.Println(bytes.EqualFold(s5, s6))
}
