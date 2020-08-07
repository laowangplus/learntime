package main

import "bytes"

// 错误使用 slice 的拼接示例
func main() {
	path := []byte("AAAA/BBBB")
	sepIndex := bytes.IndexByte(path, '/') // 4
	println(sepIndex)
	p1 := path[:sepIndex:sepIndex] // 此时 cap(p1) 指定为4， 而不是先前的
	p2 := path[sepIndex+1:]

	println("p1: ", string(p1)) // AAAA
	println("p2: ", string(p2)) // BBBB
	p1 = append(p1, "XXXXXXX"...)

	println("p1: ", string(p1)) // AAAAX
	println("p2: ", string(p2)) // XBBBB
}
