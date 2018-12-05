package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sort"
)

type TimeEntry struct {
	Hour   int
	Minute int
	Init   bool
}

func (te TimeEntry) Decrement() TimeEntry {
	minute := (te.Minute - 1) % 60
	hour := te.Hour
	if minute == 59 {
		hour = (hour - 1) % 24
	}
	return TimeEntry{Hour: hour, Minute: minute, Init: true}
}

type Entry struct {
	Hour   int    `json:"hour"`
	Month  int    `json:"month"`
	Day    int    `json:"day"`
	Minute int    `json:"minute"`
	Guard  int    `json:"guard,omitempty"`
	Event  string `json:"event,omitempty"`
}

func (te Entry) getTime() int {
	return te.Minute + te.Hour*60 + te.Day*60*24 + te.Month*60*24*31
}

type ByTime []Entry

func (a ByTime) Len() int           { return len(a) }
func (a ByTime) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByTime) Less(i, j int) bool { return a[i].getTime() < a[j].getTime() }

type Data struct {
	Data []Entry `json:"data"`
}

// Reading files requires checking most calls for errors.
// This helper will streamline our error checks below.
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getData() []Entry {
	dat, err := ioutil.ReadFile("/Users/antonhagermalm/Projects/advent-of-code/4/data.json")
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
	sort.Sort(ByTime(data))
	guards := make(map[int]map[TimeEntry]int)
	currentGuard := 0
	var fellAsleep TimeEntry
	for _, v := range data {
		// fmt.Println(data[i].Month, data[i].Day, data[i].Hour, data[i].Minute, data[i].Event, data[i].Guard)
		if v.Guard != 0 {
			currentGuard = v.Guard
			if guards[v.Guard] == nil {
				guards[v.Guard] = make(map[TimeEntry]int)
			}
		} else {
			if v.Event == "falls asleep" && !fellAsleep.Init {
				fellAsleep = TimeEntry{Hour: v.Hour, Minute: v.Minute, Init: true}
			} else if v.Event == "wakes up" && fellAsleep.Init {
				currentTime := TimeEntry{Hour: v.Hour, Minute: v.Minute, Init: true}
				for currentTime != fellAsleep {
					currentTime = currentTime.Decrement()
					// fmt.Println(currentTime)
					guards[currentGuard][currentTime]++
				}
				fellAsleep = TimeEntry{Hour: 0, Minute: 0, Init: false}
			} else {
				fmt.Println(v)
				fmt.Println(fellAsleep)
				panic("something bad mate")
			}
		}
	}
	// fmt.Println(guards)
	largest := struct {
		Id     int
		Minute int
		Times  int
	}{Id: 0, Minute: -1, Times: 0}

	for id, guard := range guards {
		for te, times := range guard {
			fmt.Println(id, times)
			if times >= largest.Times {
				largest.Id = id
				largest.Minute = te.Minute
				largest.Times = times
			}
		}
	}
	fmt.Println(largest.Times, largest.Id, largest.Minute)
	fmt.Println(largest.Minute * largest.Id)
}
