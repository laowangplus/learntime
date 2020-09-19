package main

import "fmt"

//https://exercise.acmcoder.com/online/online_judge_ques?ques_id=3364&konwledgeId=40
//度度熊收到了一个只有小写字母的字符串S，他对S的子串产生了兴趣，S的子串为S中任意连续的一段。他发现，一些子串只由一种字母构成，他想知道在S中一共有多少种这样的子串。
//
//例如在串”aaabbaa”中，度度熊想找的子串有”a”,”aa”,”aaa”,”b”,”bb”五种。
//
//
//
//（本题只考虑子串的种数，相同的子串在多个位置出现只算一次）

func main(){
	var s = "aaabba"
	//fmt.Scanf("%s", &s)

	lenth := len(s)
	count := make(map[string]int)

	for i := 0; i < lenth; i++ {
		for j := i; j < lenth; j++ {
			if s[i] == s[j] {
				count[s[i:j+1]] = 1
				//fmt.Println(s[i:j+1])
			}else{
					break
			}
		}
	}

	fmt.Println(len(count))
}
