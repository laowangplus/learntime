package main

//切片和 Map 赋值

func main() {
	m := make(map[int]*int)
	var x1 int
	m[0] = &x1 // BAD: cause of x1 escape

	s := make([]*int, 1)
	var x2 int
	s[0] = &x2 // BAD: cause of x2 escape

	//通常slice和map用户存储值的情况比较常见。避免用其存储指针可以减少不必要的变量逃逸。
}
