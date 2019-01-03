package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getData(fileName string) ([]int, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	numbers := make([]int, 0)
	// instructions := make([]instruction, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {

		line := scanner.Text()
		parts := strings.Split(line, " ")
		for _, part := range parts {
			num, err := strconv.ParseInt(part, 10, 16)
			if err != nil {
				return nil, err
			}
			numbers = append(numbers, int(num))
		}
	}

	return numbers, nil
}

func parse(nums []int) (int, int, []int) {
	// process the current node
	nodes := nums[0]
	metadata := nums[1]
	nums = append(nums[:0], nums[2:]...)
	values := make([]int, 0)
	count := 0
	tmpCount := 0
	tmpValue := 0

	// process this nodes children 1-by-1
	for i := 0; i < nodes; i++ {
		tmpCount, tmpValue, nums = parse(nums)
		count += tmpCount
		values = append(values, tmpValue)
	}

	// once children have processed, process the metadata by adding to count
	tmpMeta := 0
	for i := 0; i < metadata; i++ {
		tmpMeta += nums[i]
	}

	count += tmpMeta
	if nodes == 0 {
		// remove the metadata that was just processed
		nums = append(nums[:0], nums[metadata:]...)
		return count, tmpMeta, nums
	} else {
		//must count the child metadata to provide root node value
		tv := 0
		// for each metadata
		for _, m := range nums[:metadata] {
			// if the metadata is > than 0 then it represents the child number and
			// the metadata is <= than length of the values then we are referencing a child
			if m > 0 && m <= len(values) {
				tv += values[m-1] // since values slice is 0 based we must remove -1
			}
		}
		// remove the metadata that was just processed
		nums = append(nums[:0], nums[metadata:]...)
		return count, tv, nums
	}
}

func part1(numbers []int) error {
	//
	count, data, _ := parse(numbers)
	fmt.Println(count)
	fmt.Println(data)
	return nil
}

// after struggling with something close I peaked and found this solution :(
// https://www.reddit.com/r/adventofcode/comments/a47ubw/2018_day_8_solutions/ebc7ol0/
func main() {

	// dataFileName := "../../data/day8/test.txt"
	dataFileName := "../../data/day8/input.txt"

	numbers, err := getData(dataFileName)
	if err != nil {
		panic(err)
	}

	part1(numbers)
}
