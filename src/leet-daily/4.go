package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1, len2 := len(nums1), len(nums2)
	//res := .0
	i, j , cnt := 0, 0, 0
	arr := make([]int, (len1+len2)/2 + 1)
	for ; cnt < (len1 + len2) / 2 + 1; cnt ++ {
		if i < len1 && j < len2 {
			if nums1[i] < nums2[j] {
				arr[cnt] = nums1[i]
				i++
			} else {
				arr[cnt] = nums2[j]
				j++
			}
		} else if j < len2 {
			arr[cnt] = nums2[j]
			j++
		} else {
			arr[cnt] = nums1[i]
			i++
		}
	}

	if (len1 + len2) % 2 == 0 {
		return (float64(arr[cnt-2]) + float64(arr[cnt-1])) / 2
	} else {
		return float64(arr[cnt-1])
	}
}

func main() {
	nums1 := []int{1,3}
	nums2 := []int{2,4}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}
