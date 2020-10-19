package main

import "fmt"

//小美喜欢字母E，讨厌字母F。在小美生日时，小团送了小美一个仅包含字母E和F的字符串，
// 小美想从中选出一个包含字母E数量与字母F数量之差最大的子串。
//
//*子串：从字符串前面连续删去若干个字符，从后面连续删去若干个字符剩下的字符串
// （也可以一个都不删），例如abcab是fabcab的子串，而不是abcad的子串。
// 我们将空串看作所有字符串的子串。

func main()  {
	var n int
	fmt.Scanf("%d", &n)
	var s string
	fmt.Scanf("%s", &s)
	var location []int
	for i := 0; i < n; i++ {
		if s[i] == 'F' {
			location = append(location, i)
		}
	}


}