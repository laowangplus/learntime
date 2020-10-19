package main

import "fmt"

func main(){
	 arr := []string{"0100110", "1110110", "0100001"}
	//for i := 0 ;i < 3; i++{
	//	fmt.Scanf("%d", &arr[i])
	//}

	//big := []interface{}{}
	ludi_set := make([][]int, 0)

	for _, v := range arr {
		length := len(v)

		var ludi []int
		for i := 0 ;i < length; i++{
			if v[i] == '1' {
				ludi = append(ludi, i)
			}else{
				continue
			}
		}

		ludi_set = append(ludi_set, ludi)
	}



	fmt.Println(ludi_set)


}
