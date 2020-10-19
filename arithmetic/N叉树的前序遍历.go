package main


 type Node struct {
     Val int
     Children []*Node
 }


func preorder(root *Node) []int {
	stack := make([]*Node, 0) //初始化栈
	result := make([]int, 0) //结果

	stack = append(stack, root)

	for 0 < len(stack) {
		for root != nil{
			result = append(result, root.Val) //遍历结果

			//叶子节点
			if len(root.Children) == 0{
				break
			}

			for i := len(root.Children)-1; i > 0; i--{
				stack = append(stack, root.Children[i]) //入栈
			}

			root = root.Children[0]
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
	}

	return result

}