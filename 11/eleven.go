package main

import "fmt"

// func (e *list.Element) CNext() *Element {
// 	if p := e.Next(); p != nil {
// 		return p
// 	}
// 	return e.list.Front()
// }

type Cord struct {
	X int
	Y int
}

func main() {
	two()
}

func one() {
	gridID := 5535
	largest := 0
	largestX := 0
	largestY := 0
	for i := 1; i <= 298; i++ {
		for j := 1; j <= 298; j++ {
			curr := calcForThree(i, j, gridID)
			if curr > largest {
				largest = curr
				largestX = i
				largestY = j
			}
		}
	}
	fmt.Println(largestX, largestY)
}

func two() {
	gridID := 5535
	largest := 0
	largestX := 0
	largestY := 0
	size := 0

	for i := 1; i <= 298; i++ {
		fmt.Println(i, "%", "  size:", size, "    largest:", largest)
		for j := 1; j <= 298; j++ {
			curr, currSize := calcForCords(i, j, gridID)
			if curr > largest {
				size = currSize
				largest = curr
				largestX = i
				largestY = j
			}
		}
	}
	fmt.Println(largestX, largestY, size)
}

func calcForCords(x int, y int, gridID int) (int, int) {
	var max int
	if x < y {
		max = 301 - y
	} else {
		max = 301 - x
	}
	val := 0
	size := 0
	for i := 1; i <= max; i++ {
		_val := calcForArb(x, y, gridID, i)
		if _val > val {
			val = _val
			size = i
		}
	}
	return val, size
}

func calcForArb(x int, y int, gridID int, size int) int {
	sum := 0
	for i := x; i < x+size; i++ {
		for j := y; j < y+size; j++ {
			sum += calcForCord(i, j, gridID)
		}
	}
	return sum
}

func calcForThree(x int, y int, gridID int) int {
	sum := 0
	for i := x; i < x+3; i++ {
		for j := y; j < y+3; j++ {
			sum += calcForCord(i, j, gridID)
		}
	}
	return sum
}

func calcForCord(x int, y int, gridID int) int {
	rackID := x + 10
	return getHundVal((rackID*y+gridID)*rackID) - 5
}

func getHundVal(v int) int {
	return v % 1000 / 100
}
