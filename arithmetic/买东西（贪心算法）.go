package main

import "fmt"

//服装店新进了a条领带，b条裤子，c个帽子，d件衬衫，现在要把这些搭配起来售卖。有三种搭配方式，一条领带和一件衬衫，一条裤子和一件衬衫，一个帽子和一件衬衫。卖出一套领带加衬衫可以得到e元，卖出一套裤子加衬衫可以得到f元，卖出一套帽子加衬衫可以得到g元。现在你需要输出最大的获利方式。


func main(){
	var a,b,c,d,e,f,g int
	fmt.Scanf("%d %d %d %d %d %d %d", &a ,&b ,&c ,&d ,&e ,&f ,&g)

	money := 0


		for {
			if d <= 0 {
				break
			}
			if e >= f && e >= g {
				if a > 0 && d > 0 {
					money += g
					a--
					d--
				}else{
						e = 0
					continue
				}
			}

			if f >= e && f >= g {
				if b > 0 && d > 0 {
					money += f
					b--
					d--
				}else{
						f = 0
					continue
				}
			}

			if g >= e && g >= f {
				if c > 0 && d > 0 {
					money += g
					c--
					d--
				}else{
					g = 0
					continue
				}
			}
		}

	fmt.Println(money)

}
