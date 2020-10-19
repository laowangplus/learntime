package main

import "fmt"

var arr []int
func main(){
	arr = []int{1,4,5,2,6,3}
	quickSort(0, 5)
	fmt.Println(arr)
}

func quickSort(left, right int){
	if left < right {
		//置换
		midd := arr[left]
		index := left
		start, stop := left, right
		for right != left {
			if arr[stop] < midd{
				swap(stop, index)
				start++
			}
			if arr[start] {
			}
		}
		quickSort(left, index-1)
		quickSort(index+1, right)
	}
}


func swap(i, j int){
	s := arr[i]
	arr[i] = arr[j]
	arr[j] = s
}