package main

/**
 *
 * @param tasks int整型一维数组 待加工的零件
 * @param n int整型 n的值
 * @return int整型
 */
func leastWorkTime( tasks []int ,  n int ) int {
	// write code here
	mp := make(map[int]int)

	lock := make(map[int]int)

	for _, v := range tasks {
		_, ok := mp[v]
		if !ok {
			mp[v] = 1
		}else{
			mp[v] += 1
		}

		_, ok = lock[v]
		if !ok {
			lock[v] = 0
		}
	}

	work := 0

	for {
		//加工完成
		if len(mp) == 0 {
			break
		}
		for k ,v := range mp {
			//如果零件不存在，则删除该零件
			if v == 0 {
				delete(mp, k)
				delete(lock, k)
				continue
			}

			//如果零件正处于加锁状态
			if lock[k] != 0 {
				continue
			}

			work++
			mp[k]--
			//fmt.Println("run:",k,work)
			lock[k] = n

			for k, v := range lock {
				if v == 0 {
					continue
				}
				lock[k]--
			}
		}

		//如何都没有到时间就等待
		flag := 1
		for _, v := range lock {
			if v == 0 {
				flag = 0
				break
			}
		}
		if flag == 1 {
			for k := range lock {
				lock[k]--
			}
			work++
			//fmt.Println("wait:", mp, lock, work)
		}

	}


	//fmt.Println(mp, lock, work)
	return work
}

func main(){
	leastWorkTime([]int{1,1,1,2,3,3,3}, 2)
}
