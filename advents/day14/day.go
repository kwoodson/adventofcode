package main

import "fmt"

var (
	input = 513401
)

type elf struct {
	curPos     int
	lastPlayed int
}

func getPosition(pos, glength int) int {
	// get the correct position
	// if pos == glength-1 {
	// return 0
	// } else
	if pos > glength-1 {
		pos -= glength - 2 // array length -1 - 1(to account for 0 based array)
	}
	return pos
}

func part1() error {
	// recipeCount := 0
	game := []int{3, 7}

	e1 := elf{
		curPos: 0,
		// lastPlayed: 1,
	}
	e2 := elf{
		curPos: 1,
		// lastPlayed: 0,
	}

	for i := 0; i < 100000000; i++ {
		// calculate recipe

		r := game[e1.curPos] + game[e2.curPos]
		if r >= 10 {
			game = append(game, r/10, r%10)
		} else {
			game = append(game, r)
		}
		// fmt.Println(game)
		e1.curPos = (1 + e1.curPos + game[e1.curPos]) % len(game) // needs to move us appropriately
		e2.curPos = (1 + e2.curPos + game[e2.curPos]) % len(game) // needs to move us appropriately

	}

	fmt.Printf("After 18: %v\n", game[18:18+10])
	fmt.Printf("After 2018: %v\n", game[2018:2018+10])
	fmt.Printf("After 513401: %v\n", game[513401:513401+10])

	for i := 0; i <= len(game)-6; i++ {
		if game[i] == 5 &&
			game[i+1] == 1 &&
			game[i+2] == 3 &&
			game[i+3] == 4 &&
			game[i+4] == 0 &&
			game[i+5] == 1 {
			fmt.Printf("i=%d\n", i)
			break
		}

	}

	return nil
}

func main() {
	part1()
}
