package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Cord struct {
	*SimpleCord
	ID int
}

type SimpleCord struct {
	X int
	Y int
}

type Corners struct {
	HighestX int
	HighestY int
	LowestX  int
	LowestY  int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getData() []Cord {
	file, err := os.Open("/Users/antonhagermalm/Projects/advent-of-code/6/data")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var cords []Cord
	id := 0
	for scanner.Scan() {
		for i, v := range scanner.Text() {
			if v == 44 {
				runes := []rune(scanner.Text())
				x, err := strconv.Atoi(string(runes[0:i]))
				check(err)
				y, err := strconv.Atoi(string(runes[i+2 : len(scanner.Text())]))
				check(err)
				cords = append(cords, Cord{SimpleCord: &SimpleCord{X: x, Y: y}, ID: id})
				break
			}
		}
		id++
	}
	return cords
}

func main() {
	one()
}

func two() {
	cords := getData()
	corners := calcCorners(cords)
	sum := calcDists(cords, corners)
	fmt.Println(sum)
}

func calcDists(cords []Cord, corners Corners) int {
	amount := 0
	for i := corners.LowestX; i <= corners.HighestX; i++ {
		for j := corners.LowestY; j <= corners.HighestY; j++ {
			simpleCord := SimpleCord{X: i, Y: j}
			if getHanhattanDistSum(cords, simpleCord) < 10000 {
				amount++
			}
		}
	}
	return amount
}

func one() {
	cords := getData()
	corners := calcCorners(cords)
	nrForID := calcNrForId(cords, corners)
	IDsForInf := getIdsForInf(cords, corners)
	largest := calcLargest(IDsForInf, nrForID, corners)
	fmt.Println(largest)
}

func calcLargest(infIDs map[int]bool, nrForID map[int]int, corners Corners) int {
	largestSum := 0
	for i, v := range nrForID {
		if !infIDs[i] && v > largestSum {
			largestSum = v
		}
	}
	return largestSum
}

func getIdsForInf(cords []Cord, corners Corners) map[int]bool {
	ids := make(map[int]bool)
	for i := corners.LowestX; i <= corners.HighestX; i++ {
		cord := SimpleCord{X: i, Y: corners.HighestY}
		id := getClosest(cords, cord)
		ids[id] = true
		cord = SimpleCord{X: i, Y: corners.LowestY}
		id = getClosest(cords, cord)
		ids[id] = true
	}
	for i := corners.LowestY; i <= corners.HighestY; i++ {
		cord := SimpleCord{X: corners.HighestX, Y: i}
		id := getClosest(cords, cord)
		ids[id] = true
		cord = SimpleCord{X: corners.LowestX, Y: i}
		id = getClosest(cords, cord)
		ids[id] = true
	}
	return ids
}

func calcCorners(cords []Cord) Corners {
	highestX, highestY, lowestX, lowestY := -1, -1, 10000, 10000
	for _, v := range cords {
		if v.X < lowestX {
			lowestX = v.X
		}
		if v.Y < lowestY {
			lowestY = v.Y
		}
		if v.X > highestX {
			highestX = v.X
		}
		if v.Y > highestY {
			highestY = v.Y
		}
	}
	return Corners{HighestX: highestX, HighestY: highestY, LowestX: lowestX, LowestY: lowestY}
}

func calcNrForId(cords []Cord, corners Corners) map[int]int {
	cordMap := make(map[int]int)
	for i := corners.LowestX; i <= corners.HighestX; i++ {
		for j := corners.LowestY; j <= corners.HighestY; j++ {
			simpleCord := SimpleCord{X: i, Y: j}
			closest := getClosest(cords, simpleCord)
			cordMap[closest]++
		}
	}
	return cordMap
}

func getHanhattanDistSum(cords []Cord, cord SimpleCord) int {
	sum := 0
	for _, v := range cords {
		sum += abs(v.X-cord.X) + abs(v.Y-cord.Y)
	}
	return sum
}

func getClosest(cords []Cord, cord SimpleCord) int {
	smallestDist := abs(cords[0].X-cord.X) + abs(cords[0].Y-cord.Y)
	smallestCord := cords[0].ID
	for i, v := range cords {
		dist := abs(cords[i].X-cord.X) + abs(cords[i].Y-cord.Y)
		// fmt.Println(dist)
		if dist < smallestDist {
			smallestDist = dist
			smallestCord = v.ID
		} else if dist == smallestDist {
			smallestCord = -1
		}
	}
	return smallestCord
}

func abs(i int) int {
	if i < 0 {
		i *= -1
	}
	return i
}
