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

// End of the solution

func main() {
	var function interface{}
	// Replace with your function name
	// function = f1

	// Step 1: use reflect to get the # of input parameters and # of return (output) value
	fv := reflect.ValueOf(function).Type()
	numIn, numOut := fv.NumIn(), fv.NumOut()

	// Step 2: read input from input.txt according to the # of input parameters
	inputs := readInput(numIn)
	// Step 3: read output from output.txt according to the # of output parameters
	expects := readOutput(numOut)

	for len(inputs) > len(expects) {
		// Step 4: we allow not specifing the output
		expects = append(expects, []string{})
	}

	// Iterate all the inputs
	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		expect := expects[i]
		// Step 5: call the solution function with given input args
		output := utils.Call(function, input...)
		if equal(expect, output) {
			fmt.Println("Pass test case", i)
		} else if len(expect) == 0 {
			fmt.Printf("Input: \t%s\nGot: \t%s\n", input, output)
		} else {
			fmt.Println("Fail test case", i)
			fmt.Printf("Expected:\t%s\nGot:\t\t%s\n", expect, output)
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

func readInput(numIn int) [][]string {
	return readFileByNum(numIn, "input.txt")
}

func readOutput(numOut int) [][]string {
	return readFileByNum(numOut, "output.txt")
}

func readFileByNum(num int, fileName string) (res [][]string) {
	file, err := os.Open(fileName)

	if err != nil {
		log.Fatal("failed to open " + fileName)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var outputs []string
	cnt := 0
	for scanner.Scan() {
		outputs = append(outputs, scanner.Text())
		cnt++
		if cnt == num {
			cnt = 0
			res = append(res, outputs)
			outputs = []string{}
		}
	}
	return
}
