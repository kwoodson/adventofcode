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

// Part1 compute result
func Part1(nums []int) int {
	// circle through the numbers and verify 2 numbers that add up to 2020 and
	// then multiply
	for idx, x := range nums {
		if idx == len(nums) {
			break
		}
		for _, y := range nums[idx+1:] {
			if x+y == 2020 {
				return x * y
			}
		}
	}
	return 0
}

func Part2(nums []int) int {
	for idx, x := range nums {
		for idy, y := range nums[idx+1:] {
			for _, z := range nums[idy+1:] {
				if x+y+z == 2020 {
					return x * y * z
				}
			}
		}
	}
	return 0
}

func main() {
	// test data
	input, err := getData()
	if err != nil {
		panic(err)
	}
	fmt.Println(Part1(input))
	fmt.Println(Part2(input))

	// var results int64
	// for _, in := range input {
	// 	result := Part2(in)
	// 	// result := Part1(in)

	// 	results += int64(result)
	// }
	// // fmt.Printf("Part 1: %d\n", results)

	// fmt.Printf("Part 2: %d\n", results)
}
