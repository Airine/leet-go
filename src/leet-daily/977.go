package main

func sortedSquares(A []int) []int {
	left, right := 0, len(A)
	res := make([]int, len(A))
	idx := len(A)
	for left <= right {
		if A[left] < 0 && A[left] + A[right] < 0 {
			res[idx] = A[left] * A[left]
			left ++
		} else {
			res[idx] = A[right] * A[right]
			right --
		}
		idx --
	}
	return res
}

func main() {
	
}
