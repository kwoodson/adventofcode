package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func getData(fileName string) (string, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	lines := make([]string, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {

		line := scanner.Text()
		// get date
		lines = append(lines, line)
	}

	return lines[0], nil
}

func react(puzzle []rune) string {
	i := 0
	for i+1 < len(puzzle) {
		first := puzzle[i]
		second := puzzle[i+1]
		if unicode.ToLower(first) == unicode.ToLower(second) && first != second {
			puzzle = append(puzzle[:i], puzzle[i+2:]...)
			if i > 0 {
				i--
			}
			continue
		}
		i++
	}
	return string(puzzle)
}

func part1(puzzle string) string {
	rPuzzle := []rune(puzzle)
	return react(rPuzzle)
}

func part2(puzzle string) int {
	chars := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p',
		'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}

	results := make(map[rune]int)

	shortest := 100000
	for _, c := range chars {
		lower := strings.Replace(puzzle, string(c), "", -1)
		both := strings.Replace(lower, string(unicode.ToUpper(c)), "", -1)

		results[c] = len(react([]rune(both)))
		if shortest > results[c] {
			shortest = results[c]
			fmt.Printf("Char[%v] Rune[%v] Shortest[%d]\n", string(c), c, shortest)
		}
	}
	return shortest
}

func main() {

	dataFileName := "../../data/day5/input.txt"

	puzzle, err := getData(dataFileName)
	if err != nil {
		panic(err)
	}
	fmt.Println(len(part1(puzzle)))

	fmt.Println(part2(puzzle))
}
