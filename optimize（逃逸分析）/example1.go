package main

// Sum 函数返回 0-100 的整数之和
func Sum() int {
	const count = 100
	numbers := make([]int, count) //没有逃逸到堆
	for i := range numbers {
		numbers[i] = i + 1
	}

	var sum int
	for _, i := range numbers {
		sum += i
	}
	return sum
}

/**
因为 numbers 切片仅在 Sum函数内部使用，编译器将在栈上存储这100个整数而不是堆。也没有必要对 numbers进行垃圾回收，因为它会在 Sum 返回时自动释放。
*/

//func main()  {
//	fmt.Println(Sum())
//}
