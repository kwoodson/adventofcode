package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Password struct {
	min  int
	max  int
	pass string
	seq  string
}

func getData() ([][]string, error) {
	dataFileName := "../data/input.txt"
	f, err := os.Open(dataFileName)
	if err != nil {
		panic(err)
	}
	treeMap := make([][]string, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]string, 0)
		parts := strings.Split(line, "")
		row = append(row, parts...)
		treeMap = append(treeMap, row)
	}

	return treeMap, nil
}

func treeCount(xInc, yInc int, treeMap [][]string) int {
	count := 0
	x := 0
	y := 0

	for {
		if y >= len(treeMap) {
			break
		}
		if string(treeMap[y%len(treeMap)][x%len(treeMap[y])]) == "#" {
			count++
		}
		x += xInc
		y += yInc
	}

	return count
}

// Part1 compute result
func Part1(treeMap [][]string) int {
	return treeCount(3, 1, treeMap)
}

type Slope struct {
	x int
	y int
}

func Part2(treeMap [][]string) int {
	nums := []Slope{
		{
			x: 1,
			y: 1,
		},
		{
			x: 3,
			y: 1,
		},
		{
			x: 5,
			y: 1,
		},
		{
			x: 7,
			y: 1,
		},
		{
			x: 1,
			y: 2,
		},
	}
	results := make([]int, 0)
	for _, s := range nums {
		results = append(results, treeCount(s.x, s.y, treeMap))
	}

	answer := results[0]
	for _, r := range results[1:] {
		answer *= r
	}

	return answer
}

func main() {
	// test data
	treeMap, err := getData()
	if err != nil {
		panic(err)
	}

	fmt.Println(Part1(treeMap))
	fmt.Println(Part2(treeMap))
}
