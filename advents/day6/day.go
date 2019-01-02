package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type point struct {
	X        int
	Y        int
	shared   bool
	infinite bool
}

type gridLocation struct {
	p        point
	shortest int
	closest  []point
	ownedby  point
}

func getData(fileName string) ([]point, int, int, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, 0, 0, err
	}
	points := make([]point, 0)
	scanner := bufio.NewScanner(f)

	var xMax int
	var yMax int

	for scanner.Scan() {

		line := scanner.Text()
		parts := strings.Split(line, ", ")
		x, err := strconv.ParseInt(parts[0], 10, 16)
		if err != nil {
			return nil, 0, 0, err
		}

		y, err := strconv.ParseInt(parts[1], 10, 16)
		if err != nil {
			return nil, 0, 0, err
		}
		xint := int(x)
		yint := int(y)
		if xint > xMax {
			xMax = xint
		}
		if yint > yMax {
			yMax = yint
		}
		points = append(points, point{
			X: xint,
			Y: yint,
		})
	}

	return points, xMax, yMax, nil
}

func findDistance(a, b point) int {
	x := a.X - b.X
	if x < 0 {
		x = x * -1
	}
	y := a.Y - b.Y
	if y < 0 {
		y = y * -1
	}
	return x + y
}

func part1(points []point, xMax, yMax int) error {
	// create a dictionary to keep track of results
	results := make(map[point]int)

	for y := 0; y < yMax; y++ {
		for x := 0; x < xMax; x++ {
			p := point{
				X: y,
				Y: x,
			}
			shortest := yMax * xMax
			var ownedby int
			closest := 0

			for idx, curPoint := range points {
				// we are currently on the point we are checking so skip
				if p == curPoint {
					shortest = 0
					ownedby = idx
					break
				}

				tmpClosest := findDistance(p, curPoint)
				if tmpClosest < shortest {
					shortest = tmpClosest
					ownedby = idx
					closest = 1
				} else if shortest == tmpClosest {
					closest++
					ownedby = -1
				}
			}
			if ownedby < 0 {
				continue // we found a shared point
			}
			if y == 0 || x == 0 || y == yMax || x == xMax {
				points[ownedby].infinite = true
			}
			if closest <= 1 {
				results[points[ownedby]]++
			}
		}
	}

	largest := 0
	for p, count := range results {
		if p.infinite {
			continue
		}
		if count > largest {
			largest = count
			fmt.Printf("Point: (%d, %d)  Count: %d  Inf: %v\n", p.X, p.Y, count, p.infinite)
		}
	}
	return nil
}

func part2(points []point, xMax, yMax int) error {
	// create a dictionary to keep track of results
	grid := make([][]int, yMax)
	for i := 0; i < yMax; i++ {
		grid[i] = make([]int, xMax)
	}
	results := 0
	for y := 0; y < yMax; y++ {
		for x := 0; x < xMax; x++ {
			p := point{
				X: y,
				Y: x,
			}
			for _, curPoint := range points {
				tmpClosest := findDistance(p, curPoint)
				grid[y][x] += tmpClosest
			}
			if grid[y][x] < 10000 {
				results++
			}
		}
	}

	fmt.Printf("Results: %d\n", results)
	return nil
}

func main() {

	// dataFileName, size := "../../data/day6/test.txt"
	dataFileName := "../../data/day6/input.txt"

	points, x, y, err := getData(dataFileName)
	if err != nil {
		panic(err)
	}

	part1(points, x, y)
	part2(points, x, y)

}
