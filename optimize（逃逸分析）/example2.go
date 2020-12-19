package main

type Point struct{ X, Y int }

const Width = 640
const Height = 480

type pool struct {
	P interface{}
}

func Center(p Point) *Point {
	return &p
}

func Center2(p interface{}) interface{} {
	_ = pool{
		P: p,
	}
	return &p
}

func NewPoint() {
	p := Point{}
	Center(p)
	Center2(p)
	//fmt.Println(p)
	//go func() {
	//	Center(&p)
	//}()
}

func main() {
	NewPoint()
}
