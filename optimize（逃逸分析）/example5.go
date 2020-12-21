package main

//间接赋值

type S struct {
	M *int
}

func main() {
	type X struct {
		p *int
	}

	var i1 int
	x1 := &X{
		p: &i1, // GOOD: &X literal does not escape
	}
	_ = x1

	var i2 int
	x2 := &X{}
	x2.p = &i2 //  moved to heap: i2
	//go的逃逸分析不知道X2和i2的关系
}
