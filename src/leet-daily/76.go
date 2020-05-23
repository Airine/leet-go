package main

import "fmt"

func contains(m map[rune]int, t map[rune]int) bool {
	if len(m) > len(t) {
		return false
	}
	for k, v := range m {
		count, ok := t[k]
		if !ok || count < v {
			return false
		}
	}
	return true
}

func minWindow(s string, t string) string {
	count := map[rune]int{}
	for i := range t {
		char := rune(t[i])
		fmt.Println(string(char))
		v, ok := count[char]
		if ok {
			count[char] = v + 1
		} else {
			count[char] = 1
		}
	}

	res := map[rune]int{}
	flag := false
	low, high := 0, 0
	rl, rh := 0, len(s)+1
	for high < len(s) || contains(count, res) {
		if !contains(count, res) {
			char := rune(s[high])
			v, ok := res[char]
			if !ok {
				res[char] = 1
			} else {
				res[char] = v + 1
			}
			high ++
		} else {
			if high - low < rh - rl {
				rl, rh = low, high
			}
			if !flag {
				flag = true
			} else {
				char := rune(s[low])
				tempt := res
				v, _ := tempt[char]
				tempt[char] = v - 1
				if !contains(count, tempt) {
					flag = false
				}
				low ++
			}
		}
	}
	if rh == len(s)+1 {
		return ""
	} else {
		return s[rl:rh]
	}
}

func main() {
	fmt.Println(minWindow("ADOBECODEBANC","ABC"))
}
