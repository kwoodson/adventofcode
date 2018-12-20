package main

import (
	"bufio"
	"fmt"
	"os"
)

func getData(fileName string) ([]string, error) {
	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	lines := make([]string, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	return lines, nil
}

func processLine(inc string) map[string]int {
	chars := make(map[string]int, 10000)
	for _, char := range inc {
		chars[string(char)]++
	}

	return chars
}

func part1(lines []string) int {
	ids := make(map[int]int, 10000)

	for _, line := range lines {
		results := processLine(line)
		twoFound, threeFound := false, false
		for key, val := range results {
			if twoFound && threeFound {
				break
			}
			if val == 2 && !twoFound {
				ids[val] += 1
				twoFound = true
				fmt.Printf("(%v had %d)", key, val)

			}
			if val == 3 && !threeFound {
				ids[val] += 1
				threeFound = true
				fmt.Printf("(%v had %d)", key, val)
			}
		}
		fmt.Println()
	}

	total := 1
	for key, val := range ids {
		if val == 0 {
			continue
		}
		fmt.Printf("%d had %d\n", key, val)
		total *= val
	}

	return total
}

func main() {
	dataFileName := "../../data/day2/day_2_input.txt"
	lines, err := getData(dataFileName)
	if err != nil {
		panic(err)
	}
	fmt.Println(part1(lines))
}
