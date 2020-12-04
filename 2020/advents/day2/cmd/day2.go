package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Password struct {
	min  int
	max  int
	pass string
	seq  string
}

func getData() ([]Password, error) {
	dataFileName := "../data/input.txt"
	f, err := os.Open(dataFileName)
	if err != nil {
		panic(err)
	}
	passwords := make([]Password, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		// 9-14 c: clckqtcpcgpccckcwcc
		parts := strings.Split(line, ":")

		countSeq := strings.Split(parts[0], " ")
		s := countSeq[1]

		count := strings.Split(countSeq[0], "-")
		min, err := strconv.Atoi(count[0])
		if err != nil {
			panic(err)
		}
		max, err := strconv.Atoi(count[1])
		if err != nil {
			panic(err)
		}
		p := strings.TrimLeft(parts[1], " ")

		passwords = append(passwords, Password{
			min:  min,
			max:  max,
			pass: p,
			seq:  s,
		})
	}

	return passwords, nil
}

// Part1 compute result
func Part1(passwords []Password) int {
	// circle through the numbers and verify 2 numbers that add up to 2020 and
	// then multiply
	count := 0
	for _, p := range passwords {
		n := strings.Count(p.pass, p.seq)
		if n >= p.min && n <= p.max {
			count++
		}
	}
	return count
}

func Part2(passwords []Password) int {
	count := 0
	for _, p := range passwords {
		if (string(p.pass[p.min-1]) == p.seq && string(p.pass[p.max-1]) != p.seq) ||
			(string(p.pass[p.max-1]) == p.seq && string(p.pass[p.min-1]) != p.seq) {
			count++
		}
	}
	return count
}

func main() {
	// test data
	passwords, err := getData()
	if err != nil {
		panic(err)
	}

	fmt.Println(Part1(passwords))
	fmt.Println(Part2(passwords))

}
