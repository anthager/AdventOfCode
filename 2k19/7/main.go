package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Data struct {
	Data []int `json:"data"`
}

const debug = false
const test = false

func getData(test bool) []int {
	file := "/Users/antonhagermalm/Projects/advent-of-code/2k19/7/"
	if test {
		file += "data.test.json"
	} else {
		file += "data.json"
	}
	dat, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	var jsonData Data
	json.Unmarshal(dat, &jsonData)
	return jsonData.Data
}

func main() {
	input := getData(test)
	second(input)
	// second(input)
}

func first(input []int) {
	largest := 0
	a := []int{0, 1, 2, 3, 4}
	var sequences [][]int
	heapPermutation(a, len(a), &sequences)
	var largestSequence []int
	for _, sequence := range sequences {
		inputCopy := copySlice(input)
		output := runSequence(inputCopy, sequence)
		if output > largest {
			largest = output
			largestSequence = sequence
		}
	}

	fmt.Println(largest)
	fmt.Println(largestSequence)
	// fmt.Println(runSequence(input, [5]int{4, 3, 2, 1, 1}))
}
func second(input []int) {
	largest := 0
	list := []int{5, 6, 7, 8, 9}
	var sequences [][]int
	heapPermutation(list, len(list), &sequences)
	var largestSequence []int
	for _, sequence := range sequences {
		inputCopy := copySlice(input)
		output := runSequence(inputCopy, sequence)
		if output > largest {
			largest = output
			largestSequence = sequence
		}
	}
	fmt.Println(largest)
	fmt.Println(largestSequence)
	// fmt.Println(runSequence(input, []int{9, 7, 8, 5, 6}))
}

func runSequence(program []int, sequence []int) int {
	output := 0
	isDone := false
	nrOfAmplifiers := 5
	programs := make([][]int, nrOfAmplifiers)
	programCounters := make([]int, nrOfAmplifiers)
	for i := 0; i < nrOfAmplifiers; i++ {
		programs[i] = make([]int, len(program))
		copy(programs[i], program)
		programCounters[i] = 0
	}
	for i := 0; i < 5; i++ {
		params := []int{sequence[i%nrOfAmplifiers], output}
		output, programCounters[i%nrOfAmplifiers], isDone = runProgram(programs[i%nrOfAmplifiers], params, programCounters[i%nrOfAmplifiers])
	}
	for i := 0; !isDone; i++ {
		params := []int{output}
		var out int
		out, programCounters[i%nrOfAmplifiers], isDone = runProgram(programs[i%nrOfAmplifiers], params, programCounters[i%nrOfAmplifiers])
		if !isDone {
			output = out
		}
	}
	fmt.Println(output)
	return output
}

func runProgram(input []int, params []int, i int) (int, int, bool) {
	inputCount := 0
	output := 0
	for input[i] != 99 {
		if debug {
			fmt.Println(fmt.Sprintf("index: %d gives instruction %d", i, input[i]))
		}
		instruction := input[i]
		if input[i]%10 == 1 ||
			input[i]%10 == 2 ||
			input[i]%10 == 5 ||
			input[i]%10 == 6 ||
			input[i]%10 == 7 ||
			input[i]%10 == 8 {
			i = handleTwoInput(input, i)
		} else if input[i]%10 == 4 {
			i, output = getOutput(input, i)
			return output, i, false
		} else if input[i]%10 == 3 {
			if inputCount > 1 {
				panic("bad input count")
			}
			if debug {
				fmt.Println(fmt.Sprintf("input is: %d", params[inputCount]))
				fmt.Println(fmt.Sprintf("saving %d to position %d", params[inputCount], input[i+1]))
			}
			input[input[i+1]] = params[inputCount]
			inputCount++
			i += 2
		} else {
			panic(fmt.Sprintf("invalid instruction %d", instruction))
		}
	}
	// fmt.Println(input)
	// fmt.Println(i)
	return output, -1, true
}

func handleTwoInput(input []int, i int) int {
	instruction := input[i]
	modeOf1st := (instruction / 100) % 10
	modeOf2nd := (instruction / 1000) % 10
	input1 := 0
	input2 := 0
	if debug {
		fmt.Println(fmt.Sprintf("mode1: %d, mode2: %d", modeOf1st, modeOf2nd))
	}
	if modeOf1st == 0 {
		input1 = input[input[i+1]]
	} else {
		input1 = input[i+1]
	}
	if modeOf2nd == 0 {
		input2 = input[input[i+2]]
	} else {
		input2 = input[i+2]
	}
	if debug {
		fmt.Println(fmt.Sprintf("input1: %d, input2: %d", input1, input2))
	}
	outputAddr := input[i+3]
	if instruction%10 == 1 {
		input[outputAddr] = input1 + input2
	} else if instruction%10 == 2 {
		input[outputAddr] = input1 * input2
	} else if instruction%10 == 5 {
		// jump if true
		if input1 != 0 {
			return input2
		}
		return i + 3
	} else if instruction%10 == 6 {
		// jump if false
		if input1 == 0 {
			return input2
		}
		return i + 3
	} else if instruction%10 == 7 {
		// less than
		if input1 < input2 {
			input[outputAddr] = 1
		} else {
			input[outputAddr] = 0

		}
	} else if instruction%10 == 8 {
		// equals
		if input1 == input2 {
			input[outputAddr] = 1
		} else {
			input[outputAddr] = 0
		}
	} else {
		panic("instruction not supported yet")
	}
	if instruction != input[i] {
		return i
	}
	return i + 4
}

func getOutput(input []int, i int) (int, int) {
	mode := input[i] / 100
	val := 0
	if mode == 0 {
		val = input[input[i+1]]
	} else {
		val = input[i+1]
	}
	if debug {
		fmt.Println(fmt.Sprintf("-- output: %d", val))
	}
	return i + 2, val
}

func copySlice(oldSlice []int) []int {
	newSlice := make([]int, len(oldSlice))
	copy(newSlice, oldSlice)
	return newSlice
}

func heapPermutation(a []int, size int, result *[][]int) {
	if size == 1 {
		list := make([]int, len(a))
		copy(list, a)
		*result = append(*result, list)
	}
	for i := 0; i < size; i++ {
		heapPermutation(a, size-1, result)

		if size%2 == 1 {
			a[0], a[size-1] = a[size-1], a[0]
		} else {
			a[i], a[size-1] = a[size-1], a[i]
		}
	}
}
