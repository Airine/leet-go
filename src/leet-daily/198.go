package main

func max(a,b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	}
	res := 0
	curr, last := nums[0], 0
	for i := 1; i < len(nums); i ++ {
		res = max(curr, last + nums[i])
		curr, last = res, curr
	}
	return res
}

func main() {
}
