package main

import (
	"container/list"
	"fmt"
	"io/ioutil"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getData() *list.List {
	dat, err := ioutil.ReadFile("/Users/antonhagermalm/Projects/advent-of-code/5/data")
	check(err)
	l := list.New()
	for _, v := range dat {
		l.PushBack(rune(v))
	}
	return l
}

func main() {
	two()
}

func two() {
	list := getData()
	performTest(list)
}

func performTest(list *list.List) {
	reduce(list)
	smallest := 60000
	for i := rune(65); i < 133; i++ {
		copy := makeCopy(list)
		removePoly(copy, i)
		curr := reduce(copy)
		if curr < smallest {
			smallest = curr
		}
	}
	fmt.Println(smallest)
}

func removePoly(l *list.List, poly rune) {
	for e := l.Front(); e != nil; {
		if e.Value.(rune) == poly || e.Value.(rune) == poly+32 {
			toBeRemoved := e
			e = e.Next()
			l.Remove(toBeRemoved)
		} else {
			e = e.Next()
		}
	}
}

func reduce(list *list.List) int {
	done := false
	for !done {
		done = true
		for e := list.Front(); e != nil; e = e.Next() {
			if e.Prev() != nil && (e.Value.(rune)-e.Prev().Value.(rune) == 32 || e.Value.(rune)-e.Prev().Value.(rune) == -32) {
				prev := e
				e = e.Next()
				list.Remove(prev.Prev())
				list.Remove(prev)
				done = false
			}
			if e == nil {
				break
			}
		}
	}
	return list.Len()
}

func makeCopy(old *list.List) *list.List {
	new := list.New()
	for e := old.Front(); e != nil; e = e.Next() {
		new.PushBack(e.Value.(rune))
	}
	return new
}
