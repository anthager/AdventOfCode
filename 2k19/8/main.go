package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

const imagesSize = 150
const rowSize = 25

func getInput() [][]int {
	file := "/Users/antonhagermalm/Projects/advent-of-code/2k19/8/data"
	dat, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	input := make([][]int, len(dat)/imagesSize)
	for i := 0; i < len(dat)/imagesSize; i++ {
		input[i] = make([]int, imagesSize)
		for j := 0; j < imagesSize; j++ {
			nr, err := strconv.Atoi(string(dat[i*imagesSize+j]))
			if err != nil {
				panic(err)
			}
			input[i][j] = nr
		}
	}
	return input
}

func main() {
	input := getInput()
	// first(input)

	second(input)
}

func first(input [][]int) {
	minZeroes := -1
	product := 0
	for _, layer := range input {
		zeroes := 0
		ones := 0
		twos := 0
		for _, v := range layer {
			if v == 0 {
				zeroes++
			}
			if v == 1 {
				ones++
			}
			if v == 2 {
				twos++
			}
		}
		if zeroes < minZeroes || minZeroes == -1 {
			minZeroes = zeroes
			product = ones * twos
		}
	}
	fmt.Println(product)
}

func second(input [][]int) {
	finalLayer := make([]int, imagesSize)
	for i := range finalLayer {
		finalLayer[i] = -1
	}
	for _, layer := range input {
		for j, v := range layer {
			if finalLayer[j] == -1 && v != 2 {
				finalLayer[j] = v
			}
		}
	}
	for i := 0; i < len(finalLayer)/rowSize; i++ {
		for j := 0; j < rowSize; j++ {
			if finalLayer[i*rowSize+j] == 0 {
				fmt.Print(" ")
			} else {
				fmt.Print("X")
			}
		}
		fmt.Print("\n")
	}
}
