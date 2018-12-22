package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strings"
)

type State map[int]bool

func (s *State) update(c *Changes) *State {
	tempS := make(State)
	newS := &tempS
	smallest, largest := s.getEdgeKeys()
	for i := smallest - 2; i <= largest+2; i++ {
		var localState [5]bool
		localState = [5]bool{(*s)[i-2], (*s)[i-1], (*s)[i], (*s)[i+1], (*s)[i+2]}
		// if (i < smallest || i > largest) && !(*c)[localState] {
		// 	continue
		// }
		if (*c)[localState] {
			(*newS)[i] = true
		}
	}
	return newS
}

func (s *State) getEdgeKeys() (int, int) {
	smallest := int((reflect.ValueOf(*s).MapKeys()[0].Int()))
	largest := smallest
	for i := range *s {
		if i < smallest {
			smallest = i
		} else if i > largest {
			largest = i
		}
	}
	return smallest, largest
}

func (s *State) print() {
	smallest, largest := (*s).getEdgeKeys()
	for i := smallest; i < largest+1; i++ {
		if (*s)[i] {
			fmt.Print("#")
		} else {
			fmt.Print(".")
		}
	}
	fmt.Println()
}

func (s *State) calcSum() {
	sum := 0
	for i, v := range *s {
		if v {
			sum += i
		}
	}
	fmt.Println(sum)
}

type Changes map[[5]bool]bool

func (c *Changes) print() {
	for i, v := range *c {
		for _, val := range i {
			if val {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print(" => ", v)
		fmt.Println()
	}
}

func getData() (*State, *Changes) {
	file, err := os.Open("/Users/antonhagermalm/Projects/advent-of-code/12/data")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	initialState := strings.Split(scanner.Text(), " ")[2]
	fstate := make(State)
	state := &fstate
	for i, val := range initialState {
		if val == 35 {
			(*state)[i] = true
		} else {
			// (*state)[i] = false
		}
	}
	scanner.Scan()
	fchanges := make(Changes)
	changes := &fchanges
	for scanner.Scan() {
		s := scanner.Text()
		str := strings.Split(s, " ")[0]
		var change [5]bool
		for i, v := range str {
			if v == 35 {
				change[i] = true
			}
		}
		if []byte(strings.Split(s, " ")[2])[0] == 35 {
			(*changes)[change] = true
		} else {
			(*changes)[change] = false
		}
	}
	return state, changes
}

func main() {
	one()
}

func one() {
	state, changes := getData()
	// state.print()
	// fmt.Println(state)
	for i := 0; i < 100000; i++ {
		if i%10000 == 0 {
			fmt.Println(i)
			// fmt.Println(len(*state))
			state.calcSum()
		}
		// state.print()
		state = state.update(changes)
		// fmt.Println(i)
		// state.print()
	}
}

func two() {
	// list := getData()
	// performTest(list)
}
