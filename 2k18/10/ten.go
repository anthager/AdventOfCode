package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

func (c Cord) add(d Cord) Cord {
	return Cord{X: c.X + d.X, Y: c.Y + d.Y}
}

type Point struct {
	Cord   Cord
	Change Cord
}

func (p *Point) iterate() {
	p.Cord = p.Cord.add(p.Change)
}

type Message []*Point

func (m Message) interate() {
	for _, v := range m {
		v.iterate()
	}
}

func (m Message) print(i int) {
	minX, minY := 10000, 10000
	maxX, maxY := -10000, -10000
	cordToPoint := make(map[Cord]*Point)
	for _, p := range m {
		if p.Cord.X < minX {
			minX = p.Cord.X
		}
		if p.Cord.Y < minY {
			minY = p.Cord.Y
		}
		if p.Cord.X > maxX {
			maxX = p.Cord.X
		}
		if p.Cord.Y > maxY {
			maxY = p.Cord.Y
		}
		cord := p.Cord
		cordToPoint[cord] = p
	}
	if (maxX-minX)*(maxY-minY) > 10000 {
		return
	}
	fmt.Println(i)
	for j := minY; j <= maxY; j++ {
		for i := minX; i <= maxX; i++ {
			if cordToPoint[Cord{X: i, Y: j}] != nil {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println("----------------------------------------------------------------------")
}

func (p Point) time() Point {
	return Point{Cord: p.Cord.add(p.Change), Change: p.Change}
}

func getData() Message {
	file, err := os.Open("/Users/antonhagermalm/Projects/advent-of-code/10/data")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var points Message

	for scanner.Scan() {
		s := strings.Replace(strings.Replace(strings.Replace(scanner.Text(), " ", "", 20), ">", "", 2), "<", "", 2)
		sA := strings.Split(s, "=")

		p := strings.Split(strings.Split(sA[1], "v")[0], ",")
		c := strings.Split(sA[2], ",")
		p1, _ := strconv.Atoi(p[0])
		p2, _ := strconv.Atoi(p[1])
		c1, _ := strconv.Atoi(c[0])
		c2, _ := strconv.Atoi(c[1])
		cord := Cord{X: p1, Y: p2}
		point := Point{Cord: cord, Change: Cord{X: c1, Y: c2}}
		points = append(points, &point)
	}
	return points
}

func main() {
	one()
}

func one() {
	points := getData()
	for i := 0; i < 100000; i++ {
		points.interate()
		points.print(i)
	}

	// cord := Cord{X: 0, Y: 0}
	// point := &Point{Cord: cord, Change: Cord{X: 1, Y: 1}}
	// point2 := point
	// point.iterate()
	// fmt.Println(point2)
}

func two() {
	// list := getData()
	// performTest(list)
}
