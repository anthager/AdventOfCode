package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Data struct {
	Data []int `json:"data"`
}

// Reading files requires checking most calls for errors.
// This helper will streamline our error checks below.
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	two()
}

func getData() []int {
	dat, err := ioutil.ReadFile("/Users/antonhagermalm/Projects/advent-of-code/first/data.json")
	check(err)
	var jsonData Data
	json.Unmarshal(dat, &jsonData)
	return jsonData.Data
}

func two() {
	var data = getData()
	used := make(map[int]bool)
	var sum int
	i := 0
	for true {
		sum += data[i%len(data)]
		i++
		if !used[sum] {
			used[sum] = true
		} else {
			fmt.Println(sum)
			break
		}
	}
}

func one() {
	var sum int
	data := getData()
	for _, v := range data {
		sum += v
		fmt.Println(sum)
	}
}
