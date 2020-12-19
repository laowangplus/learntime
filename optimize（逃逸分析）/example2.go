package main

type Point struct{ X, Y int }
type Value struct{ int }

const Width = 640
const Height = 480

func Center(p *Point) {
	p.X = Width / 2
	p.Y = Height / 2
}

func Print(a, b *Value) int {
	a.int, b.int = 2, 3
	return a.int + b.int
}

func NewPoint() {
	p := new(Point)
	Center(p)
	a, b := Value{1}, Value{1}
	go func(a1, b1 Value) {
		_ = Print(&a1, &b1)
	}(a, b)
	//fmt.Println(a,b)
}

func main() {
	NewPoint()
}
