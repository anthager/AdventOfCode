package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func getData(test bool) []string {
	file := "/Users/antonhagermalm/Projects/advent-of-code/2k19/4/"
	if test {
		file += "data.test"
	} else {
		file += "data"
	}
	dat, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	input := strings.Split(string(dat), "-")

	return input
}

func main() {
	input := getData(false)
	// res := first(input)
	// fmt.Println(res)
	second(input)
}

func first(input []string) int {
	start, end := getValuesAsInts(input)
	numbers := 0
	for i := start; i < end; i++ {
		if meetCriteria(strconv.Itoa(i)) {
			fmt.Println(i)
			numbers++
		}
	}
	return numbers
}

func second(input []string) {
	start, end := getValuesAsInts(input)
	numbers := 0
	for i := start; i < end; i++ {
		if meetCriteria(strconv.Itoa(i)) {
			// fmt.Println(i)
			numbers++
		}
	}
	fmt.Println(numbers)
}

func meetCriteria(n string) bool {
	alwaysHigher := true
	for i := 1; i < len(n); i++ {
		alwaysHigher = alwaysHigher && n[i-1] <= n[i]
	}
	return alwaysHigher && hasSequenceOfExactly2(n)
}

func findLargestGroup(n string) (string, int) {
	largest := ""
	amount := 0
	for i := 0; i < len(n); i++ {
		for j := i + 1; j < len(n) && n[i] == n[j]; j++ {
			_amount := j - i + 1
			if string(n[i]) > largest || (string(n[i]) == largest && _amount > amount) {
				amount = _amount
				largest = string(n[i])
			}
		}
	}
	return largest, amount
}

func hasSequenceOfExactly2(n string) bool {
	for i := 0; i < len(n); {
		amount := 1
		j := i + 1
		for (j < len(n)) && (n[i] == n[j]) {
			amount++
			j++
		}
		if amount == 2 {
			return true
		}
		i = j
	}
	return false
}

func getValuesAsInts(input []string) (int, int) {
	start, err := strconv.Atoi(input[0])
	if err != nil {
		panic(err)
	}
	end, err := strconv.Atoi(input[1])
	if err != nil {
		panic(err)
	}
	return start, end
}
