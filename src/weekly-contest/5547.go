package main

import (
    "fmt"
    "sort"
)

func isArithmeticSubarrays(nums []int) bool {
    if len(nums) == 0 {
        return false
    } else if len(nums) <= 2 {
        return true
    }
    sorted := make([]int, len(nums))
    copy(sorted, nums)
    sort.Ints(sorted)

    diff := sorted[1] - sorted[0]
    for i := 2; i < len(sorted); i++ {
        if sorted[i] - sorted[i-1] != diff {
            return false
        }
    }
    return true

}

func checkArithmeticSubarrays(nums []int, l []int, r []int) []bool {
    res := make([]bool, len(l))
    for i := range l {
        res[i] = isArithmeticSubarrays(nums[l[i]:r[i]+1])
    }
    return res
}

func main() {
    fmt.Println(checkArithmeticSubarrays([]int{4, 6, 5, 9, 3, 7}, []int{0, 0, 2}, []int{2, 3, 5}))
}
