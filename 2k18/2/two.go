package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	s "strings"
)

type Data struct {
	Data []string `json:"data"`
}

type HD struct {
	str1 string
	str2 string
	i1   int
	i2   int
	hd   int
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

func getData() []string {
	dat, err := ioutil.ReadFile("/Users/antonhagermalm/Projects/advent-of-code/2/data.json")
	fmt.Println(dat)
	check(err)
	var jsonData Data
	json.Unmarshal(dat, &jsonData)
	return jsonData.Data
}

func two() {
	data := getData()
	smallestHD := HD{hd: 27}
	for i1, str1 := range data {
		for i2, str2 := range data {
			if i1 == i2 {
				continue
			}
			hd := calcHammingDistance(str1, str2)
			if smallestHD.hd > hd {
				smallestHD.str1 = str1
				smallestHD.str2 = str2
				smallestHD.i1 = i1
				smallestHD.i2 = i2
				smallestHD.hd = hd
			}
		}
	}
	fmt.Println(smallestHD)
	fmt.Println(cleanStrings(smallestHD.str1, smallestHD.str2))
}

func calcHammingDistance(str1 string, str2 string) int {
	hd := 0
	for i := range str1 {
		if str1[i]-str2[i] != 0 {
			hd++
		}
	}
	return hd
}

func cleanStrings(str1 string, str2 string) string {
	hd := calcHammingDistance(str1, str2)
	var strB s.Builder
	strB.Grow(len(str1) - hd)
	for i := range str1 {
		if str1[i]-str2[i] == 0 {
			fmt.Fprintf(&strB, "%c", str1[i])
		}
	}
	return strB.String()
}

func one() {
	data := getData()
	ids := make(map[int]int)
	chSum := 1
	for _, v := range data {
		vals := make(map[rune]int)
		usedVals := make(map[int]bool)
		for _, val := range v {
			vals[val]++
		}
		for _, v := range vals {
			if v > 1 && !usedVals[v] {
				usedVals[v] = true
			}
		}
		for i := range usedVals {
			ids[i]++
		}
	}
	for _, v := range ids {
		chSum *= v
	}
	fmt.Println(chSum)
}
