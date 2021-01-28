package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"

	"github.com/airine/leet-go/utils"
)

// Start of the solution
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i+1; j < len(nums); j++ {
			if nums[i] + nums[j] == target {
				return []int{i,j}
			}
		}
	}
	return []int{-1,-1}
}
// End of the solution

func main() {
	var function interface{}
	// Replace with your function name
	// function = f1
	function = twoSum
	fv := reflect.ValueOf(function).Type()
	numIn, numOut := fv.NumIn(), fv.NumOut()

	inputs := readInput(numIn)
	expects := readOutput(numOut)

	for len(inputs) > len(expects) {
		expects = append(expects, []string{})
	}

	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		expect := expects[i]
		output := utils.Call(function, input...)
		if equal(expect, output) {
			fmt.Println("Pass test case", i)
		} else {
			fmt.Println("Fail test case", i)
			fmt.Printf("Expected:\t%s\nGot:\t\t%s", expect, output)
		}
	}
}

func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if !strings.EqualFold(v, b[i]) {
			return false
		}
	}
	return true
}

func readInput(numIn int) (res [][]string) {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal("failed to open input.txt")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var inputs []string
	cnt := 0
	for scanner.Scan() {
		inputs = append(inputs, scanner.Text())
		cnt++
		if cnt == numIn {
			cnt = 0
			res = append(res, inputs)
			inputs = []string{}
		}
	}
	return
}

func readOutput(numIn int) (res [][]string) {
	file, err := os.Open("output.txt")

	if err != nil {
		log.Fatal("failed to open output.txt")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var outputs []string
	cnt := 0
	for scanner.Scan() {
		outputs = append(outputs, scanner.Text())
		cnt++
		if cnt == numIn {
			cnt = 0
			res = append(res, outputs)
			outputs = []string{}
		}
	}
	return
}
