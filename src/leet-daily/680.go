package main

import "fmt"

func palindrome(s string) bool {
	var low, high int
	high = len(s) - 1
	for low < high {
		if s[low] != s[high] {
			return false
		}
		low ++
		high --
	}
	return true
}

func validPalindrome(s string) bool {
	var low, high int
	high = len(s) - 1
	for low < high {
		if s[low] != s[high] {
			return palindrome(s[low+1:high]) || palindrome(s[low:high])
		}
		low ++
		high --
	}
	return true
}

func main() {
	fmt.Println(validPalindrome("aba"))
	fmt.Println(validPalindrome("abca"))
}
