package main

// Sum 函数返回 0-100 的整数之和
func Sum() int {
	const count = 100
	numbers := make([]int, count)
	for i := range numbers {
		numbers[i] = i + 1
	}

	var sum int
	for _, i := range numbers {
		sum += i
	}
	return sum
}

//func main()  {
//	fmt.Println(Sum())
//}
