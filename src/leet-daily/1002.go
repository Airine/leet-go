package main

import "fmt"

func commonChars(A []string) []string {
	a := uint8('a')
	counters := make([][]uint8, len(A))
	var res []string
	for i := range counters {
		counters[i] = make([]uint8, 26)
	}

	for i := range A {
		for j := 0; j < len(A[i]); j++ {
			//fmt.Println(A[i][j] - a)
			char := A[i][j] - a
			counters[i][char] += 1
		}
	}

	for j := 0; j < 26; j++ {
		min := uint8(100)
		for i := range A {
			if counters[i][j] < min {
				min = counters[i][j]
			}
		}
		if min > 0 {
			for k := 0; k < int(min); k++ {
				res = append(res, string(rune(a+uint8(j))))
			}
		}
	}

	return res
}

func main() {
	A := []string{"bella","label","roller"}
	fmt.Println(commonChars(A))
}
