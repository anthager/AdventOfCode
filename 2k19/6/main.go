package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func getInput(test bool) [][]string {
	file := "/Users/antonhagermalm/Projects/advent-of-code/2k19/6/"
	if test {
		file += "data-2.test"
	} else {
		file += "data"
	}
	dat, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	lines := make([][]string, 0)
	for _, v := range strings.Split(string(dat), "\n") {
		lines = append(lines, strings.Split(v, ")"))
	}
	return lines
}

func main() {
	input := getInput(false)
	// first(input)
	second(input)
}

func first(input [][]string) {
	checkSums := make(map[string]int)
	currentSum := 0
	lastSum := 0
	firstIteration := true
	for currentSum != lastSum || firstIteration {
		firstIteration = false
		for i := 0; i < len(input); i++ {
			checkSums[input[i][1]] = checkSums[input[i][0]] + 1
		}
		sum := 0
		for _, v := range checkSums {
			sum += v
		}
		lastSum = currentSum
		currentSum = sum
	}
	fmt.Println(currentSum)
}

func second(input [][]string) {
	orbitTree := buildTree(input)
	ca := findCommonClosestCommonAncestor("YOU", "SAN", orbitTree)
	distanceToCommonForYOU := len(orbitTree["YOU"]) - 1 - len(orbitTree[ca])
	distanceToCommonForSAN := len(orbitTree["SAN"]) - 1 - len(orbitTree[ca])
	distanceFromYouToSanta := distanceToCommonForYOU + distanceToCommonForSAN
	fmt.Println(distanceFromYouToSanta)
}

func buildTree(input [][]string) map[string][]string {
	orbitTree := make(map[string][]string)
	currentSum := 0
	lastSum := 0
	firstIteration := true
	for currentSum != lastSum || firstIteration {
		firstIteration = false
		for i := 0; i < len(input); i++ {
			orbitTree[input[i][1]] = append(copySlice(orbitTree[input[i][0]]), input[i][0])
		}
		sum := 0
		for _, v := range orbitTree {
			sum += len(v)
		}
		lastSum = currentSum
		currentSum = sum
	}
	return orbitTree
}

func findCommonClosestCommonAncestor(first string, second string, orbitTree map[string][]string) string {
	firstAncestors := orbitTree[first]
	secondAncestors := orbitTree[second]
	commonAncestor := ""
	for i := 0; firstAncestors[i] == secondAncestors[i]; i++ {
		if i >= len(firstAncestors) || i >= len(secondAncestors) {
			panic(fmt.Sprintf("i (%d) is larger than or equal the length of first, not good", i))
		}
		commonAncestor = firstAncestors[i]
	}
	return commonAncestor
}

func copySlice(oldSlice []string) []string {
	newSlice := make([]string, len(oldSlice))
	copy(newSlice, oldSlice)
	return newSlice
}
