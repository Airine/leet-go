package main

import "fmt"

func minWindow(s string, t string) string {
	var sFreq [128]int
	var tFreq [128]int
	for _, v := range t {
		tFreq[v] ++
	}
	distance := 0
	tLen := len(t)
	low, high := 0, 0
	rl, rh := 0, len(s)+1
	for high < len(s) {
		char := s[high]
		sFreq[char]++
		if tFreq[char] >= sFreq[char] {
			distance ++
		}
		high ++
		for distance == tLen {
			if high - low < rh - rl {
				rh, rl = high, low
			}
			char := s[low]
			if  tFreq[char] == sFreq[char] {
				distance --
			}
			sFreq[char]--
			low++
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
