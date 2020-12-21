package main

//间接调用

func foox(p *int, x int) {
	*p = x
}

func main() {
	var y1 int
	foox(&y1, 42) // GOOD: y1 does not escape

	var y2 int
	px := func(p *int, x int) {
		*p = x
	}
	px(&y2, 42) // BAD: Cause of y2 escape

	var y3 int
	p := foox
	p(&y3, 42) // BAD: Cause of y3 escape
	//传递函数作为参数，做闭包传递的方式，可能会导致变量逃逸到堆，减少这种方式们可以降低GC的压力。
}
