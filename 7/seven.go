package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

// Worker for getting shit
type Worker struct {
	TimeLeft rune
	Job      rune
}

// Dec for getting shit
func (w *Worker) Dec() {
	if w.TimeLeft > 0 {
		w.TimeLeft--
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// ByValue for getting shit
type ByValue []rune

func (r ByValue) Len() int           { return len(r) }
func (r ByValue) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r ByValue) Less(i, j int) bool { return r[i] < r[j] }

func getData() (map[rune][]rune, []rune) {
	file, err := os.Open("/Users/antonhagermalm/Projects/advent-of-code/7/data")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	requirements := make(map[rune][]rune)
	stepsMap := make(map[rune]bool)
	for scanner.Scan() {
		requirements[rune(scanner.Text()[36])] = append(requirements[rune(scanner.Text()[36])], rune(scanner.Text()[5]))
		stepsMap[rune(scanner.Text()[5])] = true
		stepsMap[rune(scanner.Text()[36])] = true
	}
	var stepsArr []rune
	for i := range stepsMap {
		stepsArr = append(stepsArr, i)
	}
	sort.Sort(ByValue(stepsArr))
	return requirements, stepsArr
}

func main() {
	two()
}

func two() {
	requirements, stepsArr := getData()
	done := make(map[rune]bool)
	onGoing := make(map[rune]bool)
	workers := [...]*Worker{
		&Worker{Job: 0, TimeLeft: 0},
		&Worker{Job: 0, TimeLeft: 0},
		&Worker{Job: 0, TimeLeft: 0},
		&Worker{Job: 0, TimeLeft: 0},
		&Worker{Job: 0, TimeLeft: 0}}
	time := 0
	for len(done) < len(stepsArr) {
		var recentDone []rune
		for _, worker := range workers {
			if worker.TimeLeft == 0 {
				if worker.Job != 0 {
					recentDone = append(recentDone, worker.Job)
					worker.Job = 0
				}
				next := getNextJobWithOnGoing(requirements, done, onGoing, stepsArr)
				if next != 0 {
					fmt.Println(time)
					worker.TimeLeft = next - 5
					worker.Job = next
					onGoing[next] = true
				}
			}
			worker.Dec()
		}
		for _, v := range recentDone {
			done[v] = true
			onGoing[v] = false
		}
		time++
	}
	fmt.Println(time)
}

func getNextJobWithOnGoing(requirements map[rune][]rune, done map[rune]bool, onGoing map[rune]bool, stepsArr []rune) rune {
	next := rune(0)
	for _, v := range stepsArr {
		if done[v] {
			continue
		}
		if onGoing[v] {
			continue
		}
		reqs := requirements[v]
		meetReq := true
		for _, r := range reqs {
			if !done[r] {
				meetReq = false
				break
			}
		}
		if !meetReq {
			continue
		}
		next = v
		break
	}
	return next
}

func getNextJob(requirements map[rune][]rune, done map[rune]bool, stepsArr []rune) rune {
	next := rune(-1)
	for _, v := range stepsArr {
		if done[v] {
			continue
		}
		reqs := requirements[v]
		meetReq := true
		for _, r := range reqs {
			if !done[r] {
				meetReq = false
				break
			}
		}
		if !meetReq {
			continue
		}
		next = v
		break
	}
	return next
}

func one() {
	requirements, stepsArr := getData()
	var doneSequence []rune
	done := make(map[rune]bool)
	for len(done) < len(stepsArr) {
		next := getNextJob(requirements, done, stepsArr)
		doneSequence = append(doneSequence, next)
		done[next] = true
	}
	fmt.Println(string(doneSequence))
}
