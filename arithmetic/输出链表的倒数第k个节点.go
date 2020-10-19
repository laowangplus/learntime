package main

import (
	"fmt"
)

type Node1 struct {
	val int
	node *Node1
}

func main(){
	list := &Node1{}

	head := list

	var n int
	fmt.Scanf("%d", &n)

	for i := 1; i <= n; i++{

		if i == n {
			head.val = i
			break
		}

		head.val = i
		head.node = &Node1{}
		head = head.node
	}

	head.node = list

	fmt.Println("-----")

	head = list
	xunhuan := 3

	for i := 1; true; i++ {
		if head.node == head {
			break
		}

		if (i+1)%xunhuan == 0 {
			fmt.Println(head.node.val)
			head.node = head.node.node
			i++
		}

		head = head.node
	}

	head = list

	//for {
	//	if head.node == nil {
	//		break
	//	}
	//	fmt.Println(head.val)
	//	head = head.node
	//}
}
