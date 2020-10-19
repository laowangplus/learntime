package main

import "fmt"

//有一个X*Y的网格，小团要在此网格上从左上角到右下角，只能走格点且只能向右或向下走。请设计一个算法，计算小团有多少种走法。给定两个正整数int x,int y，请返回小团的走法数目。


func main(){
	var x, y int
	fmt.Scanf("%d %d", &x, &y)

	if x == 1 && y == 1 {
		fmt.Println(2)
		return
	}

	//二维数组表示到每一格的走法, 下标为0的边缘格子走法始终为1
	dp := make([][]int, x+1)
	for i := 0; i <= x; i++ {
		dp[i] = make([]int, y+1)
		dp[i][0] = 1
	}

	for k := range dp[0] {
		dp[0][k] = 1
	}

	for i := 1; i <= x; i++ {
		for j := 1; j <= y; j++ {
			dp[i][j] = dp[i][j-1] + dp[i-1][j]
		}
	}

	fmt.Println(dp[x][y])

}