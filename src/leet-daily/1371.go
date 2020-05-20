package main

import (
	"fmt"
)

func findTheLongestSubstring(s string) int {
	ans, status := 0, 0
	/* status是一个5位二进制数，用来表示5位元音字母在s的每个位置的奇偶性
	 * 如：10011 表示 a,e,i,o,u 分别出现了 奇，偶，偶，奇，奇 次
	 */
	pos := make([]int, 1 << 5) // 1 << 5 == 32, pos用来记录32种状态第一次出现的位置
	for i := 0; i < len(pos); i++ {
		pos[i] = -1
	}
	pos[0] = 0
	for i := 0; i < len(s); i ++ {
		switch s[i] {
		case 'a':
			status ^= 1 << 0 // 00001
		case 'e':
			status ^= 1 << 1 // 00010
		case 'i':
			status ^= 1 << 2 // 00100
		case 'o':
			status ^= 1 << 3 // 01000
		case 'u':
			status ^= 1 << 4 // 10000
		}
		if pos[status] >= 0 {
			tempt := i + 1 - pos[status]
			if tempt > ans {
				ans = tempt
			}
		} else {
			pos[status] = i + 1
		}
	}
	return ans
}

func main() {
	fmt.Println(findTheLongestSubstring("eleetminicoworoep"))
	fmt.Println(findTheLongestSubstring("leetcodeisgreat"))
	fmt.Println(findTheLongestSubstring("bcbcbc"))
}
