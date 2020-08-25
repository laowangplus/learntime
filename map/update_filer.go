package main

import "fmt"

type data struct {
	name string
}

//func main() {
//	m := map[string]data{
//		"x": {"Tom"},
//	}
//	r := m["x"]
//	r.name = "Jerry"
//	m["x"] = r
//	fmt.Println(m) // map[x:{Jerry}]
//}

func main() {
	m := map[string]*data{
		"x": {"Tom"},
	}
	m["z"].name = "what???"
	fmt.Println(m["x"])
}
