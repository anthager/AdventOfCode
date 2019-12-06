package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Data struct {
	Data []int `json:"data"`
}

func getData(test bool) []int {
	file := "/Users/antonhagermalm/Projects/advent-of-code/2k19/5/"
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
	input := getData(false)
	// first(input)
	second(input)
}

func first(input []int) {
	for i := 0; input[i] != 99; {
		fmt.Println(fmt.Sprintf("index: %d gives instruction %d", i, input[i]))
		instruction := input[i]
		if input[i]%10 == 1 || input[i]%10 == 2 {
			i = handleTwoInput(input, i)
		} else if input[i]%10 == 4 {
			i = output(input, i)
		} else if input[i]%10 == 3 {
			input[input[i+1]] = 1
			i += 2
		} else {
			panic(fmt.Sprintf("invalid instruction %d", instruction))
		}
	}
}

func second(input []int) {
	for i := 0; input[i] != 99; {
		fmt.Println(fmt.Sprintf("index: %d gives instruction %d", i, input[i]))
		instruction := input[i]
		if input[i]%10 == 1 ||
			input[i]%10 == 2 ||
			input[i]%10 == 5 ||
			input[i]%10 == 6 ||
			input[i]%10 == 7 ||
			input[i]%10 == 8 {
			i = handleTwoInput(input, i)
		} else if input[i]%10 == 4 {
			i = output(input, i)
		} else if input[i]%10 == 3 {
			input[input[i+1]] = 5
			i += 2
		} else {
			panic(fmt.Sprintf("invalid instruction %d", instruction))
		}
	}
}

func handleTwoInput(input []int, i int) int {
	instruction := input[i]
	modeOf1st := (instruction / 100) % 10
	modeOf2nd := (instruction / 1000) % 10
	input1 := 0
	input2 := 0
	fmt.Println(fmt.Sprintf("mode1: %d, mode2: %d", modeOf1st, modeOf2nd))
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
	fmt.Println(fmt.Sprintf("input1: %d, input2: %d", input1, input2))
	outputAddr := input[i+3]
	if instruction%10 == 1 {
		input[outputAddr] = input1 + input2
	} else if instruction%10 == 2 {
		input[outputAddr] = input1 * input2
	} else if instruction%10 == 5 {
		// jump if true
		if input1 != 0 {
			return input2
		} else {
			return i + 3
		}
	} else if instruction%10 == 6 {
		// jump if false
		if input1 == 0 {
			return input2
		} else {
			return i + 3
		}
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

func output(input []int, i int) int {
	mode := input[i] / 100
	val := 0
	if mode == 0 {
		val = input[input[i+1]]
	} else {
		val = input[i+1]
	}
	fmt.Println(fmt.Sprintf("-- output: %d", val))
	return i + 2
}

func copySlice(oldSlice []int) []int {
	newSlice := make([]int, len(oldSlice))
	copy(newSlice, oldSlice)
	return newSlice
}
