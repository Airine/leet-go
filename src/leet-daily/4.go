package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1, len2 := len(nums1), len(nums2)
	res := .0
	i, j , cnt, tempt:= 0, 0, 0, .0
	var n1, n2 float64
	if (len1 + len2) % 2 == 0 {
		for ; cnt < (len1 + len2) / 2 + 1; cnt ++ {
			if i >= len1 {
				n1, n2 = tempt, float64(nums2[j])
				j ++
				tempt = n2
			} else if  j >= len2 {
				n1, n2 = tempt, float64(nums2[i])
				tempt = n2
				i ++
			} else {
				n1, n2 = float64(nums1[i]), float64(nums2[j])
				if n1 < n2 {
					n2 = tempt
					tempt = n1
					i++
				} else {
					n1 = tempt
					tempt = n2
					j++
				}
			}
			res = (n1 + n2) / 2
		}
	} else {
		for ; cnt < (len1 + len2) / 2; cnt ++ {
			if i >= len1 {
				n1, n2 = tempt, float64(nums2[j])
				j ++
				tempt = n2
			} else if  j >= len2 {
				n1, n2 = tempt, float64(nums2[i])
				tempt = n2
				i ++
			} else {
				n1, n2 = float64(nums1[i]), float64(nums2[j])
				if n1 < n2 {
					tempt = n1
					i++
				} else {
					tempt = n2
					j++
				}
			}
			res = tempt
		}
	}
	return res
}

func main() {
	nums1 := []int{1,3}
	nums2 := []int{2}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}
