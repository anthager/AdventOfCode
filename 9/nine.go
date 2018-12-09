package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/anthager/AdventOfCode2k18"
)

// func (e *list.Element) CNext() *Element {
// 	if p := e.Next(); p != nil {
// 		return p
// 	}
// 	return e.list.Front()
// }

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getData() (int, int) {
	dat, err := ioutil.ReadFile("/Users/antonhagermalm/Projects/advent-of-code/9/data2")
	check(err)
	str := strings.Split(string(dat), " ")
	players, _ := strconv.Atoi(string(str[0]))
	lastMarble, _ := strconv.Atoi(string(str[6]))
	return players, lastMarble
}

func main() {
	one()
}

func one() {
	_, lastmarble := getData()
	createList(lastmarble)
}

func two() {
	// list := getData()
	// performTest(list)
}

func createList(lastMarble int) {
	l := coollist.New()
	current := l.Front()

	for i := 0; i <= lastMarble; i++ {
		l.InsertAfter(i, current.Next())
		current = current.Next().Next()
	}
	d := l.Front()
	for i := 0; i < l.Len(); i++ {
		fmt.Print(d.Value)
		d = d.Next()
	}

}

func addMarbleToList(e *coollist.Element) {
	e = e.Next().Next()
	e = e.Next()
}

// func removePoly(l *AdventOfCode2k18.List, poly rune) {
// 	for e := l.Front(); e != nil; {
// 		if e.Value.(rune) == poly || e.Value.(rune) == poly+32 {
// 			toBeRemoved := e
// 			e = e.Next()
// 			l.Remove(toBeRemoved)
// 		} else {
// 			e = e.Next()
// 		}
// 	}
// }

// func reduce(list *list.List) int {
// 	done := false
// 	for !done {
// 		done = true
// 		for e := list.Front(); e != nil; e = e.Next() {
// 			if e.Prev() != nil && (e.Value.(rune)-e.Prev().Value.(rune) == 32 || e.Value.(rune)-e.Prev().Value.(rune) == -32) {
// 				prev := e
// 				e = e.Next()
// 				list.Remove(prev.Prev())
// 				list.Remove(prev)
// 				done = false
// 			}
// 			if e == nil {
// 				break
// 			}
// 		}
// 	}
// 	return list.Len()
// }

// func makeCopy(old *list.List) *list.List {
// 	new := list.New()
// 	for e := old.Front(); e != nil; e = e.Next() {
// 		new.PushBack(e.Value.(rune))
// 	}
// 	return new
// }
