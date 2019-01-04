package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type point struct {
	velX int // left to right
	velY int // up or down
	posX int // left or right
	posY int // up or down
}

func getData(fileName string) ([]point, int, int, int, int, error) {
	maxX := 0
	maxY := 0
	lowX := 0
	lowY := 0
	f, err := os.Open(fileName)
	if err != nil {
		return nil, maxX, maxY, lowX, lowY, err
	}
	points := make([]point, 0)
	// instructions := make([]instruction, 0)
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {

		line := scanner.Text()
		parts := strings.Split(line, "> velocity=<")
		//position=< 3,  6>
		posSplit := strings.Split(parts[0], "<")
		posTrim := strings.TrimLeft(posSplit[1], " ")
		posComma := strings.Split(posTrim, ", ")
		posX, err := strconv.ParseInt(posComma[0], 10, 32)
		if err != nil {
			return nil, maxX, maxY, lowX, lowY, err
		}
		posYTrim := strings.TrimLeft(posComma[1], " ")
		posY, err := strconv.ParseInt(posYTrim, 10, 32)
		if err != nil {
			return nil, maxX, maxY, lowX, lowY, err
		}

		velTrim := strings.TrimLeft(parts[1], " ")
		velTrim = strings.TrimRight(velTrim, ">")
		velSplit := strings.Split(velTrim, ", ")
		velX, err := strconv.ParseInt(velSplit[0], 10, 32)
		if err != nil {
			return nil, maxX, maxY, lowX, lowY, err
		}

		velYTrim := strings.TrimLeft(velSplit[1], " ")
		velY, err := strconv.ParseInt(velYTrim, 10, 32)

		if int(posX) > maxX {
			maxX = int(posX)
		}
		if int(posY) > maxY {
			maxY = int(posY)
		}
		if int(posX) < lowX {
			lowX = int(posX)
		}
		if int(posY) < lowY {
			lowY = int(posY)
		}

		points = append(points, point{
			posX: int(posX),
			posY: int(posY),
			velX: int(velX),
			velY: int(velY),
		})

	}

	return points, maxX, maxY, lowX, lowY, nil
}

func calcVelocity(points []point) []point {
	for idx, p := range points {
		p.posX += p.velX
		p.posY += p.velY
		points[idx] = p
	}
	return points
}

func part1(points []point, maxX, maxY, lowX, lowY int) error {
	timer := 0
	update := false
	for timer < maxY {
		//are all y-points within 10?
		bigY, lilY := 0, maxY
		bigX, lilX := 0, maxX
		for _, p := range points {
			if p.posY < lilY {
				lilY = p.posY
			}
			if p.posY > bigY {
				bigY = p.posY
			}
			if p.posX < lilX {
				lilX = p.posX
			}
			if p.posX > bigX {
				bigX = p.posX
			}
		}
		// if bigX-lilX < 100 {
		if bigY-lilY <= 10 && bigX-lilX < 200 {
			fmt.Printf("Timer: %d\n", timer)
			for y := lilY; y < bigY+2; y++ {
				for x := lilX; x < bigX+2; x++ {
					update = false
					for _, p := range points {
						if p.posY == y && p.posX == x {
							fmt.Printf("#")
							update = true
							break
						}
					}
					if !update {
						fmt.Printf(" ")
					}
				}
				fmt.Println()
			}
			fmt.Println()
			fmt.Println()
		}
		// update points with velocity
		points = calcVelocity(points)
		timer++
	}
	return nil
}

func main() {
	// dataFileName := "../../data/day10/test.txt"
	dataFileName := "../../data/day10/input.txt"

	points, mx, my, lx, ly, err := getData(dataFileName)
	if err != nil {
		panic(err)
	}

	// for _, p := range points {
	// 	fmt.Printf("X:%d Y:%d  VX:%d VY:%d\n", p.posX, p.posY, p.velX, p.velY)
	// }
	part1(points, mx, my, lx, ly)
}
