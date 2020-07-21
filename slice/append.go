package main

import "fmt"

func main() {

	var s = make([]int, 0, 10)

	t := append(s, 1, 2, 3, 4) //t与s的地址相同，但是t长度为0，所以下面打印两值不相同
	fmt.Println(t)
	fmt.Println(s)

	fmt.Printf("addr:%p \t\tlen:%v content:%v\n", t, len(t), t)
	//addr:0xc00001e0a0               len:4 content:[1 2 3 4]
	fmt.Printf("addr:%p \t\tlen:%v content:%v\n", s, len(s), s)
	//addr:0xc00001e0a0               len:0 content:[]

	s = append(s, 1, 2, 3, 5)

	fmt.Printf("addr:%p \t\tlen:%v content:%v\n", t, len(t), t)
	//addr:0xc00001e0a0               len:4 content:[1 2 3 5]
	fmt.Printf("addr:%p \t\tlen:%v content:%v\n", s, len(s), s)
	//addr:0xc00001e0a0               len:0 content:[1 2 3 5]

}
