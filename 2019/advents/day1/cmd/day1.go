package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func getData() ([]int, error) {
	dataFileName := "../data/input.txt"
	f, err := os.Open(dataFileName)
	if err != nil {
		panic(err)
	}
	nums := make([]int, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		tmpX, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		nums = append(nums, tmpX)
	}

	return nums, nil
}

// Part2 second part
func Part2(input int) int {
	results := make([]int, 0)
	cur := input

	for {
		cur = Part1(cur)
		if cur <= 0 {
			break
		}

		results = append(results, cur)

	}

	// sum the results slice
	all := 0
	for _, data := range results {
		all += data
	}

	return all
}

// Part1 compute result
func Part1(nums int) int {
	return nums/3 - 2
}

func main() {
	// test data
	input, err := getData()
	if err != nil {
		panic(err)
	}
	var results int64
	for _, in := range input {
		result := Part2(in)
		// result := Part1(in)

		results += int64(result)
	}
	// fmt.Printf("Part 1: %d\n", results)

	fmt.Printf("Part 2: %d\n", results)
}
