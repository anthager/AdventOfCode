package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

// Worker for getting shit
type Node struct {
	Children []Node
	Meta     []int
}

// Dec for getting shit

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getData() []int {
	dat, err := ioutil.ReadFile("/Users/antonhagermalm/Projects/advent-of-code/8/data")
	check(err)
	var str []int
	for i := 0; len(dat) > 0; i++ {
		if dat[i] == 32 {
			val, _ := strconv.Atoi(string(dat[:i]))
			str = append(str, val)
			dat = dat[i+1:]
			i = 0
		}
	}
	return str
}

func main() {
	two()
}
func two() {
	str := getData()
	sum, _ := calcAdvancedScore(str)
	fmt.Println(sum)
}

func calcAdvancedScore(stream []int) (int, int) {
	sum := 0
	index := 2
	metaValueToOccurences := make(map[int]int)
	childrenIDToScore := make(map[int]int)
	// if stream[0] == 0 {
	for i := 0; i < stream[0]; i++ {
		res, end := calcAdvancedScore(stream[index:])
		childrenIDToScore[i+1] = res
		index += end
	}
	for i := 0; i < stream[1]; i++ {
		if stream[0] != 0 {
			metaValueToOccurences[stream[index]]++
		} else {
			sum += stream[index]
		}
		index++
	}
	fmt.Println(childrenIDToScore)
	for child, score := range childrenIDToScore {
		sum += score * metaValueToOccurences[child]
	}
	// }
	return sum, index
}

func one() {
	str := getData()
	sum, _ := calcSimpleScore(str)
	fmt.Println(sum)
}

func calcSimpleScore(stream []int) (int, int) {
	sum := 0
	index := 2
	for i := 0; i < stream[0]; i++ {
		res, end := calcSimpleScore(stream[index:])
		sum += res
		index += end
	}
	for i := 0; i < stream[1]; i++ {
		sum += stream[index]
		index++
	}
	return sum, index
}

func getNode(stream []int) Node {
	var nodes []Node
	fmt.Println(stream)
	if stream[0] > 0 {
		nodes = getChildren(stream[2:len(stream)-1], stream[0])
	}
	var meta []int
	for j := 0; j < stream[1]; j++ {
		meta = append(meta, stream[len(stream)-1-j])
	}
	return Node{Meta: meta, Children: nodes}
}

func getChildren(stream []int, amount int) []Node {
	currStart := 0
	currEnd := 0
	var nodes []Node
	// fmt.Println(amount)
	for i := 0; i < amount; i++ {
		// fmt.Println(stream[currStart])
		if stream[currStart] == 0 {
			currEnd = stream[currStart+1] + 1 + currStart
			nodes = append(nodes, getNode(stream[currStart:currEnd+1]))
			// fmt.Println(stream[currEnd])
			currStart = currEnd + 1
			// fmt.Println(currStart)
			// fmt.Println(len(stream))
		}
	}
	return nodes
}
