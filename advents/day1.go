package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func getData() ([]int, error) {
	dataFileName := "../data/day1/day_1_input.txt"
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

func part2(start int, nums []int) (int, error) {
	freq := make(map[int]bool, 10000)

        freq[start] = true
	count := 0
	for {
		for _, x := range nums {
			start += x
			if freq[start] {
				fmt.Printf("Repeated frequency: %d\n", start)
				return start, nil
			}
			freq[start] = true

		}
		count += 1
		if count > 1000 {
			panic("ENDLESS?? 1000")
		}
	}

	return 0, nil
}

func part1(nums []int) (int, error) {
	var start int

	for _, x := range nums {
		start += x
	}

	return start, nil
}

func main() {
	nums, err := getData()
	if err != nil {
		panic(err)
	}
	start, err := part1(nums)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Frequence Part 1: %d\n", start)
	part2(0, nums)
}
