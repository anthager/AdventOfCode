package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Data struct {
	Data []int `json:"data"`
}

func getData() []int {
	dat, err := ioutil.ReadFile("/Users/antonhagermalm/Projects/advent-of-code/2k19/1/data.json")
	if err != nil {
		panic(err)
	}
	var jsonData Data
	json.Unmarshal(dat, &jsonData)
	return jsonData.Data
}

func mapToFuels(masses []int) []int {
	fuels := make([]int, len(masses))
	for i, v := range masses {
		fuel := v/3 - 2
		if fuel > 0 {
			fuels[i] = fuel
		} else {
			fuels[i] = 0
		}
	}
	return fuels
}

func sumAllFuels(fuel []int) int {
	fuelSum := 0
	for _, v := range fuel {
		fuelSum += v
	}
	return fuelSum
}

func withoutZeros(fuels []int) []int {
	nonZeroFuels := make([]int, 0)
	for _, v := range fuels {
		if v > 0 {
			nonZeroFuels = append(nonZeroFuels, v)
		}
	}
	return nonZeroFuels
}

func main() {
	masses := getData()
	fuelSum := 0
	for len(masses) > 0 {
		fuels := mapToFuels(masses)
		fuelSum += sumAllFuels(fuels)
		masses = withoutZeros(fuels)
		fmt.Println(len(masses))
	}
	fmt.Println(fuelSum)
}
