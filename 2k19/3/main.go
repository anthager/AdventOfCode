package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Data struct {
	Data [][]Direction `json:"data"`
}

type Direction struct {
	Direction string `json:"direction"`
	Distance  int    `json:"distance"`
}

type Cord struct {
	Right int
	Up    int
}

func getData(test bool) Data {
	file := "/Users/antonhagermalm/Projects/advent-of-code/2k19/3/"
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
	err = json.Unmarshal(dat, &jsonData)
	if err != nil {
		panic(err)
	}
	return jsonData
}

func main() {
	input := getData(false)
	// first(input)
	second(input)
}

func first(input Data) {
	path1 := createEntirePath(input.Data[0], []Cord{Cord{Up: 0, Right: 0}})
	path2 := createEntirePath(input.Data[1], []Cord{Cord{Up: 0, Right: 0}})
	common := findCommonElement(path1, path2)
	shortest := findShortestManhattan(common)
	fmt.Println(shortest)
}

func second(input Data) {
	path1 := createEntirePath(input.Data[0], []Cord{Cord{Up: 0, Right: 0}})
	path2 := createEntirePath(input.Data[1], []Cord{Cord{Up: 0, Right: 0}})
	_, dists := findCommonElementWithDist(path1, path2)
	earliest := findEarliest(dists)
	fmt.Println(earliest)
}

func createEntirePath(input []Direction, path []Cord) []Cord {
	if len(input) == 0 {
		return path
	}
	for i := 0; i < input[0].Distance; i++ {
		lastCord := path[len(path)-1]
		path = append(path, getNextCord(input[0].Direction, lastCord))
	}
	return createEntirePath(input[1:], path)
}

func getNextCord(direction string, lastCord Cord) Cord {
	switch direction {
	case "R":
		return Cord{
			Right: lastCord.Right + 1,
			Up:    lastCord.Up,
		}
	case "L":
		return Cord{
			Right: lastCord.Right - 1,
			Up:    lastCord.Up,
		}
	case "U":
		return Cord{
			Right: lastCord.Right,
			Up:    lastCord.Up + 1,
		}
	case "D":
		return Cord{
			Right: lastCord.Right,
			Up:    lastCord.Up - 1,
		}
	}
	panic(fmt.Sprintf("direction: '%s' is not valid", direction))
}

func findCommonElementWithDist(first []Cord, second []Cord) ([]Cord, []int) {
	common := make([]Cord, 0)
	dists := make([]int, 0)
	for fi, f := range first {
		for si, s := range second {
			if isEqual(f, s) {
				common = append(common, f)
				dists = append(dists, fi+si)
			}
		}
	}
	return common, dists
}

func findCommonElement(first []Cord, second []Cord) []Cord {
	common := make([]Cord, 0)
	for _, f := range first {
		for _, s := range second {
			if isEqual(f, s) {
				common = append(common, f)
			}
		}
	}
	return common
}

func isEqual(first Cord, second Cord) bool {
	return first.Right == second.Right && first.Up == second.Up
}

func findShortestManhattan(common []Cord) int {
	shortest := -1
	for _, v := range common {
		manhattan := calcManhattan(v)
		if manhattan != 0 && (shortest == -1 || manhattan < shortest) {
			shortest = manhattan
		}
	}
	return shortest
}

func findEarliest(dists []int) int {
	smallest := 0
	for _, dist := range dists {
		if smallest == 0 || dist < smallest {
			smallest = dist
		}
	}
	return smallest
}

func calcManhattan(c Cord) int {
	right := c.Right
	up := c.Up
	res := 0
	if right < 0 {
		res += (right * -1)
	} else {
		res += right
	}
	if up < 0 {
		res += (up * -1)
	} else {
		res += up
	}

	return res
}
