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
	file := "/Users/antonhagermalm/Projects/advent-of-code/2k19/2/"
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
	for i := 0; input[i] != 99; i += 4 {
		instruction := input[i]
		input1 := input[input[i+1]]
		input2 := input[input[i+2]]
		outputAddr := input[i+3]
		if instruction == 1 {
			input[outputAddr] = input1 + input2
		} else if instruction == 2 {
			input[outputAddr] = input1 * input2
		} else {
			panic(fmt.Sprintf("instruction = %d", instruction))
		}
	}
	fmt.Println(input[0])
}

func second(input []int) int {
	noun := 0
	verb := 0
	for ; noun < 100; noun++ {
		verb = 0
		for ; verb < 100; verb++ {
			inputCopy := copySlice(input)
			inputCopy[1] = noun
			inputCopy[2] = verb
			for i := 0; inputCopy[i] != 99; i += 4 {
				instruction := inputCopy[i]
				input1 := inputCopy[inputCopy[i+1]]
				input2 := inputCopy[inputCopy[i+2]]
				outputAddr := inputCopy[i+3]
				if instruction == 1 {
					inputCopy[outputAddr] = input1 + input2
				} else if instruction == 2 {
					inputCopy[outputAddr] = input1 * input2
				} else {
					panic(fmt.Sprintf("instruction = %d", instruction))
				}
			}
			// if noun%100 == 0 && verb%100 == 0 {
			// 	fmt.Println(fmt.Sprintf("result: %d", inputCopy[0]))
			// 	fmt.Println(fmt.Sprintf("noun: %d", noun))
			// 	fmt.Println(fmt.Sprintf("verb: %d", verb))
			// }
			if inputCopy[0] == 19690720 {
				fmt.Println(fmt.Sprintf("correct result: %d", noun*100+verb))
				return noun*100 + verb
			}
		}
	}
	fmt.Println("not found")
	return -1
}

func copySlice(oldSlice []int) []int {
	newSlice := make([]int, len(oldSlice))
	copy(newSlice, oldSlice)
	return newSlice
}
