package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Tile struct {
	X int
	Y int
}

type Id struct {
	Id    int   `json:"id"`
	Start []int `json:"start"`
	Size  []int `json:"size"`
}

type Data struct {
	Data []Id `json:"data"`
}

// Reading files requires checking most calls for errors.
// This helper will streamline our error checks below.
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getData() []Id {
	dat, err := ioutil.ReadFile("/Users/antonhagermalm/Projects/advent-of-code/3/data.json")
	check(err)
	var jsonData Data
	json.Unmarshal(dat, &jsonData)
	return jsonData.Data
}

func main() {
	two()
}

func two() {
	data := getData()
	tiles := make(map[Tile]int)
	for _, v := range data {
		for j := 0; j < v.Size[0]; j++ {
			for k := 0; k < v.Size[1]; k++ {
				tile := Tile{X: v.Start[0] + j, Y: v.Start[1] + k}
				tiles[tile]++
			}
		}
	}
	sum := 0
	for _, v := range tiles {
		if v > 1 {
			sum++
		}
	}
	for _, v := range data {
		shouldBreak := false
		for j := 0; j < v.Size[0]; j++ {
			for k := 0; k < v.Size[1]; k++ {
				if shouldBreak {
					break
				}
				tile := Tile{X: v.Start[0] + j, Y: v.Start[1] + k}
				if tiles[tile] > 1 {
					shouldBreak = true
				}
			}
			if shouldBreak {
				break
			}
		}
		if !shouldBreak {
			fmt.Println(v.Id)
		}
	}
}

func one() {

}
