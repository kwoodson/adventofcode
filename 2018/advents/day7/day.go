package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"sync"
)

var mux sync.Mutex

type worker struct {
	id      int
	timer   int
	current string
}

type instruction struct {
	ID     string
	PreReq []string
}

func getData(fileName string) (map[string]instruction, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	instructions := make(map[string]instruction)
	// instructions := make([]instruction, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {

		line := scanner.Text()
		parts := strings.Split(line, " ")
		pre := parts[1]
		step := parts[7]
		if key, ok := instructions[step]; ok {
			key.PreReq = append(key.PreReq, pre)
			instructions[step] = key
		} else {
			instructions[step] = instruction{
				ID:     step,
				PreReq: []string{pre},
			}
		}
		if _, ok := instructions[pre]; ok {
			//key.PreReq = append(key.PreReq, pre)
		} else {
			instructions[pre] = instruction{
				ID:     pre,
				PreReq: []string{},
			}
		}
	}

	return instructions, nil
}

func clearPreReq(instructions map[string]instruction, current string) map[string]instruction {
	// mux.Lock()
	for key, ins := range instructions {
		for i := 0; i < len(ins.PreReq); i++ {
			if ins.PreReq[i] == current {
				ins.PreReq = append(ins.PreReq[:i], ins.PreReq[i+1:]...)
				instructions[key] = ins
			}
		}
	}
	// mux.Unlock()
	return instructions
}

func getInsructions(instructions map[string]instruction) string {
	exec := make([]string, 0)
	for step, ins := range instructions {
		if len(ins.PreReq) == 0 {
			exec = append(exec, step) //
		}
	}
	if len(exec) == 0 {
		return ""
	}
	sort.Strings(exec)
	return exec[0]
}

func getInsructions2(instructions map[string]instruction) []string {
	exec := make([]string, 0)
	for step, ins := range instructions {
		if len(ins.PreReq) == 0 {
			exec = append(exec, step) //
		}
	}
	// if len(exec) == 0 {
	// 	return []string{}
	// }
	sort.Strings(exec)
	return exec
}

func part1(instructions map[string]instruction) error {
	//
	executed := make([]string, 0)
	for len(instructions) > 0 { // while prereqs exist then we need to keep processing
		// clear prereqs that are in exec
		exec := getInsructions(instructions)
		delete(instructions, exec)
		instructions = clearPreReq(instructions, exec)
		executed = append(executed, exec)
	}
	fmt.Println(strings.Join(executed, ""))
	return nil
}

func isCurrentlyExecuting(currently []string, current string) bool {
	for _, letter := range currently {
		if letter == current {
			return true
		}
	}
	return false
}

func part2(instructions map[string]instruction, workerCount int) error {
	// string map
	values := map[string]int{
		"A": 1, "B": 2, "C": 3, "D": 4, "E": 5, "F": 6, "G": 7,
		"H": 8, "I": 9, "J": 10, "K": 11, "L": 12, "M": 13, "N": 14,
		"O": 15, "P": 16, "Q": 17, "R": 18, "S": 19, "T": 20, "U": 21,
		"V": 22, "W": 23, "X": 24, "Y": 25, "Z": 26,
	}
	executed := make([]string, 0)
	currentlyExecuting := make([]string, 0)
	workers := make([]worker, workerCount)
	for idx := range workers {
		workers[idx].id = idx + 1
	}
	timer := 0
	// available := workerCount
	exec := make([]string, 0)
	for len(instructions) > 0 { // while prereqs exist then we need to keep processing
		fmt.Printf("Timer: %d ", timer)

		for _, w := range workers {
			fmt.Printf("W%d: %v ", w.id, w.current)
		}
		fmt.Println()
		// get work
		exec = getInsructions2(instructions)

		for idx, w := range workers {
			for idy, letter := range exec {
				if w.current == "" && timer >= w.timer && !isCurrentlyExecuting(currentlyExecuting, letter) {
					fmt.Printf("Worker%d - %v  Second: %d\n", w.id, letter, timer)
					w.current = letter
					w.timer = timer + 60 + values[letter]
					workers[idx] = w
					currentlyExecuting = append(currentlyExecuting, letter)
					exec = append(exec[:idy], exec[idy+1:]...)
				}
			}
		}
		timer++

		for idx, w := range workers {
			if timer >= w.timer && w.current != "" {
				delete(instructions, w.current)
				instructions = clearPreReq(instructions, w.current)
				for idx, letter := range currentlyExecuting {
					if w.current == letter {
						fmt.Printf("Worker%d - Finished %v  Second: %d\n", w.id, letter, timer)
						currentlyExecuting = append(currentlyExecuting[:idx], currentlyExecuting[idx+1:]...)
						break
					}
				}
				executed = append(executed, w.current)
				w.current = ""
				w.timer = 0
				workers[idx] = w
			}
		}

	}
	fmt.Println(strings.Join(executed, ""))
	return nil
}

func main() {

	// dataFileName := "../../data/day7/test.txt"
	dataFileName := "../../data/day7/input.txt"

	ins, err := getData(dataFileName)
	if err != nil {
		panic(err)
	}

	// part1(ins)
	part2(ins, 5)
}
