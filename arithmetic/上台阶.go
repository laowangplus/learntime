package main

import "fmt"

//有一楼梯共m级，刚开始时你在第一级，若每次只能跨上一级或二级，要走上第m级，共有多少走法？
//
//注：规定从一级到一级有0种走法。

func main(){
	var n int
	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		var m int
		fmt.Scanf("%d", &m)

		method := make([]int, m+1)
		method[1] = 1
		method[2] = 1
		for j := 3; j <= m; j++ {
			method[j] = method[j-1] + method[j-2]
		}
		fmt.Println(method[m])
	}
}
