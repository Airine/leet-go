package main

import "fmt"

// 前缀和

func subarraysDivByK(A []int, K int) int {
	record := map[int]int{0:1} // 需要注意的一个边界条件是，我们需要对哈希表初始化，记录 record[0]=1，这样就考虑了前缀和本身被 K 整除的情况。
	sum, res := 0, 0
	for _, elem := range A {
		sum += elem
		mod := ( sum % K + K ) % K
		res += record[mod]
		record[mod] ++
	}
	return res
}

func main() {
	arr := []int{4,5,0,-2,-3,1}
	k := 5
	fmt.Println(subarraysDivByK(arr, k))
}
