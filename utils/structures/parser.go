package structures

import (
	"strconv"
	"strings"
)

func parseArr(array string) []string {
	n := len(array)
	if array[0] != '[' || array[n-1] != ']' {
		return nil
	}
	array = array[1 : n-1]
	return strings.Split(array, ",")
}

func ParseIntArr(array string) []int {
	rawArr := parseArr(array)
	if rawArr == nil {
		return nil
	} else if len(rawArr[0]) == 0 {
		return []int{}
	}
	res := make([]int, len(rawArr))
	for i, v := range rawArr {
		res[i], _ = strconv.Atoi(v)
	}
	return res
}

func Parse2dIntArr(array string) (res [][]int) {
	n := len(array)
	if array[0] != '[' || array[n-1] != ']' {
		return nil
	}
	start := 1
	array = array[1 : n-1]
	for i, v := range array {
		if v == '[' {
			start = i
		} else if v == ']' {
			arr := ParseIntArr(array[start : i+1])
			res = append(res, arr)
		}
	}
	return
}

func ParseStringArr(array string) []string {
	return parseArr(array)
}

func IntArrToString(arr []int) string {
	n := len(arr)
	rawArr := make([]string, n)
	for i, v := range arr {
		rawArr[i] = strconv.Itoa(v)
	}
	return "[" + strings.Join(rawArr, ",") + "]"
}
