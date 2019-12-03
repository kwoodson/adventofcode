package main

import (
	"fmt"
	"strconv"
)

type point struct {
	Y int
	X int
}

func calcPowerLevel(x, y, serial int) (int, error) {
	rackID := x + 10
	powerLevel := (rackID*y + serial) * rackID
	hundredDigit := 0
	if powerLevel > 99 && powerLevel < 1000 {
		hundredDigit = powerLevel / 100
	} else {
		powerLevelString := strconv.Itoa(powerLevel)
		hd := powerLevelString[len(powerLevelString)-3 : len(powerLevelString)-2]
		plt, err := strconv.ParseInt(hd, 10, 32)
		if err != nil {
			return 0, err
		}
		hundredDigit = int(plt)
	}
	return hundredDigit - 5, nil
}

func dumpArea(grid [][]int, x, y int) {
	fmt.Printf("Grid at %d, %d\n", x, y)
	total := 0
	y--
	x--
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i > 1 && i < 4 && j > 1 && j < 4 {
				//tmpi := grid[i+y-1][j+x-1]
				total += grid[i+y][j+x]
				//fmt.Println(tmpi)
			}
			fmt.Printf("%d ", grid[i+y-1][j+x-1])
		}
		fmt.Println()
	}
	// fmt.Printf("Score: %d\n", total)
}

func part1(serial int) ([][]int, error) {

	grid := make([][]int, 300)
	for i := 0; i < len(grid); i++ {
		grid[i] = make([]int, 300)
	}

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			val, err := calcPowerLevel(x, y, serial)
			if err != nil {
				return nil, err
			}
			grid[y][x] = val
		}
	}

	return grid, nil
}

// room for improvement here:
// score by column and cache results then move across and just add the next row
func findScores(grid [][]int, square int) (map[point]int, point, int) {
	var tmpP point
	high := 0
	scores := make(map[point]int)
	for y := 0; y < len(grid)-square; y++ {
		for x := 0; x < len(grid[y])-square; x++ {
			p := point{Y: y, X: x}
			total := 0
			for i := 0; i < square; i++ {
				for j := 0; j < square; j++ {
					total += grid[y+i][x+j]
				}
			}
			scores[p] = total
			if total > high {
				high = total
				tmpP = p
			}
		}
	}
	return scores, tmpP, scores[tmpP]
}

func part2(grid [][]int, serial int) {
	// we need a function that searches by a squares
	// 1x1, 2x2, 3x3, ..., 300,300  this seems perfect for dynamic programming
	high := 0
	sqSize := 0
	var p point
	// for i := 4; i < len(grid)-i; i++ {
	for i := 3; i < 50; i++ {
		_, tmp, score := findScores(grid, i)

		if score > high {
			high = score
			sqSize = i
			p = tmp
		}
	}
	fmt.Printf("High score: %d,%d with size %d with serial %d scores %d\n", p.X, p.Y, sqSize, serial, high)

}

func main() {

	// fmt.Println(calcPowerLevel(3, 5, 8))
	// fmt.Println(calcPowerLevel(122, 79, 57))
	// fmt.Println(calcPowerLevel(217, 196, 39))
	// fmt.Println(calcPowerLevel(101, 153, 71))
	// test1
	grid, err := part1(18)
	if err != nil {
		panic(err)
	}
	_, tmp, high := findScores(grid, 3)

	fmt.Printf("    REAL: %d,%d with serial %d scores %d\n", tmp.X, tmp.Y, 18, high)
	fmt.Printf("EXPECTED: %d,%d with serial %d scores %d\n", 33, 45, 18, 29)
	// test 2
	grid, err = part1(42)
	if err != nil {
		panic(err)
	}
	_, tmp, high = findScores(grid, 3)

	fmt.Printf("    REAL: %d,%d with serial %d scores %d\n", tmp.X, tmp.Y, 18, high)
	fmt.Printf("EXPECTED: %d,%d with serial %d scores %d\n", 21, 61, 42, 30)

	// Part 1:
	grid, err = part1(2866)
	if err != nil {
		panic(err)
	}
	_, tmp, high = findScores(grid, 3)
	fmt.Printf("High score: %d,%d with size %d with serial %d scores %d\n", tmp.X, tmp.Y, 3, 2866, high)
	// Part 2:
	part2(grid, 2866)
}
