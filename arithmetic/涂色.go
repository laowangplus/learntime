package main

import "fmt"

//输入包含两个字符串，都仅包含大写字母，每一种字母代表一种颜色。 第一个字符串S代表小A手上的颜料，
// 第二个字符串T代表画需要的颜料。
//
//1≤|S|,|T|≤1000

//stdout
//输出包含一个数，即最多能涂多少个色块。

func main()  {
	var s,t string
	fmt.Scanf("%s\n%s", &s, &t)

	out := make([]uint8, len(t))
	for j := 0; j < len(t); j++ {
		out[j] = t[j]
	}

	count := 0

	for i := 0; i < len(s); i++ {
		for k, v := range out{
			if s[i] == v {
				out[k] = 0
				count++
				break
			}
		}
	}

	fmt.Println(count)
}