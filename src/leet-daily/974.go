package main

import "fmt"

func subarraysDivByK(A []int, K int) int {
	res := 0
	for i, _ := range A {
		sum := 0
		for j := i ; j < len(A); j ++ {
			sum += A[j]
			if sum % K == 0 {
				res += 1
			}
		}
	}
	return res
}

func main() {
	arr := []int{4,5,0,-2,-3,1}
	k := 5
	fmt.Println(subarraysDivByK(arr, k))
}
