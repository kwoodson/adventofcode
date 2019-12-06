package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
		allNumbersString := strings.Split(line, ",")
		for _, n := range allNumbersString {
			tmpX, err := strconv.Atoi(n)
			if err != nil {
				panic(err)
			}
			nums = append(nums, tmpX)
		}
	}

	return nums, nil
}

// Part2 second part
func Part2(input []int) int {
	const result int = 19690720

	for {
		for i := 1; i < 100; i++ {
			for j := 1; j < 100; j++ {
				tmpData := make([]int, len(input))
				copy(tmpData, input)
				if result == Part1(tmpData, i, j)[0] {
					return i*100 + j
				}

			}
		}
	}

	return 0
}

// Part1 compute result
func Part1(nums []int, first, second int) []int {
	nums[1] = first
	nums[2] = second
	index := 0

	for {
		// start on index and process the command
		if nums[index] == 99 {
			// return nums
			break
		}

		// if op code is 1 then add
		if nums[index] == 1 {
			nums[nums[index+3]] = nums[nums[index+1]] + nums[nums[index+2]]
		} else if nums[index] == 2 {
			nums[nums[index+3]] = nums[nums[index+1]] * nums[nums[index+2]]
		}

		index += 4
	}

	return nums
}

func main() {
	// test data
	input, err := getData()
	if err != nil {
		panic(err)
	}
	// modify positions
	tmpData := make([]int, len(input))
	copy(tmpData, input)
	result := Part1(tmpData, 12, 2)
	fmt.Printf("Part 1: %d\n", result[0])
	res := Part2(input)
	fmt.Printf("Part 2: %d\n", res)
}
