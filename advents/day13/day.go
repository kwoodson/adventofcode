package main

import (
	"bufio"
	"container/ring"
	"fmt"
	"os"
	"sort"
	"strings"
)

var (
	turns *ring.Ring
	// turns = []string{"<", "^", ">", "v"}
)

type point struct {
	Y int
	X int
}
type car struct {
	direction string
	turn      *ring.Ring
	loc       point
	id        int
}

func getData(fileName string) ([][]string, []car, error) {

	f, err := os.Open(fileName)
	if err != nil {
		return nil, nil, err
	}
	defer f.Close()

	cars := make([]car, 0)
	// cars := make(map[point]car, 0)

	game := make([][]string, 0)
	scanner := bufio.NewScanner(f)
	lineRow := 0
	carID := 0
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "")
		row := make([]string, 0)
		row = append(row, parts...)
		game = append(game, row)
		for idx, part := range parts {
			if part == "<" || part == "^" || part == ">" || part == "v" {
				cturns := ring.New(3)
				cturns.Value = 0
				cturns = cturns.Next()
				cturns.Value = 1
				cturns = cturns.Next()
				cturns.Value = 2
				cturns = cturns.Next()

				loc := point{
					Y: lineRow,
					X: idx,
				}
				cars = append(cars, car{
					direction: part,
					turn:      cturns,
					loc:       loc,
					id:        carID,
				})
				carID++
				// these are assumptions but not sure what the game would look like without
				if part == "<" || part == ">" {
					game[lineRow][idx] = "-"
				} else if part == "^" || part == "v" {
					game[lineRow][idx] = "|"
				}
			}
		}
		lineRow++
	}

	return game, cars, nil
}

func dumpGame(game [][]string, cars []car) {
	for y := 0; y < len(game); y++ {
		for x := 0; x < len(game[y]); x++ {
			carFound := false
			for _, car := range cars {
				if car.loc.Y == y && car.loc.X == x {
					fmt.Printf("%v%d", car.direction, car.id)
					carFound = true
					break
				}
			}
			if !carFound {
				fmt.Printf("%v", game[y][x])
			}
		}
		fmt.Println()
	}
	fmt.Println()
	fmt.Println()
}

func advance(game [][]string, c car) car {
	//

	if c.direction == "<" {
		c.loc.X--
	} else if c.direction == "^" {
		c.loc.Y--
	} else if c.direction == ">" {
		c.loc.X++
	} else if c.direction == "v" {
		c.loc.Y++
	}

	if game[c.loc.Y][c.loc.X] == "\\" && c.direction == ">" { // right turn
		c.direction = "v"
	} else if game[c.loc.Y][c.loc.X] == "/" && c.direction == ">" { // left turn
		c.direction = "^"
	} else if game[c.loc.Y][c.loc.X] == "/" && c.direction == "v" { // right turn
		c.direction = "<"
	} else if game[c.loc.Y][c.loc.X] == "\\" && c.direction == "v" { // left turn
		c.direction = ">"
	} else if game[c.loc.Y][c.loc.X] == "/" && c.direction == "^" { // right turn
		c.direction = ">"
	} else if game[c.loc.Y][c.loc.X] == "\\" && c.direction == "^" { // right left
		c.direction = "<"
	} else if game[c.loc.Y][c.loc.X] == "/" && c.direction == "<" { // left turn
		c.direction = "v"
	} else if game[c.loc.Y][c.loc.X] == "\\" && c.direction == "<" { // right turn
		c.direction = "^"
	} else if game[c.loc.Y][c.loc.X] == "+" {
		// turn left, go striaght, or turn right
		if c.turn.Value.(int) != 1 {
			c.direction = nextPosition(c.direction, c.turn.Value.(int))
		}
		c.turn = c.turn.Next() // move to next 0, 1, 2
	} else if game[c.loc.Y][c.loc.X] == "-" && (c.direction == "<" || c.direction == ">") {
		// do nothing because path is valid

	} else if game[c.loc.Y][c.loc.X] == "|" && (c.direction == "^" || c.direction == "v") {
		// do nothing because path is valid
	} else {
		panic("DID NOT HIT ANYTHING VALID")
	}

	return c
}

func nextPosition(current string, next int) string {
	var rval string
	for i := 0; i < turns.Len(); i++ {
		if current == turns.Value.(string) {
			break
		}
		turns = turns.Next()
	}
	if next == 0 {
		rval = turns.Prev().Value.(string)
	} else if next == 2 {
		rval = turns.Next().Value.(string)
	}
	return rval
}

func detectCollision(cars []car, c car) int {
	// detect collision
	for i := 0; i < len(cars); i++ {
		if cars[i].id == c.id {
			continue
		}
		if cars[i].loc.Y == c.loc.Y && cars[i].loc.X == c.loc.X {
			fmt.Printf("Collision detected: (%d, %d)\n", cars[i].loc.X, cars[i].loc.Y)
			return i
		}
	}
	return -1
}

func part1(game [][]string, cars []car) error {
	tick := 0
	for {
		fmt.Printf("Tick: %d\n", tick)

		for idx, car := range cars {
			// update each car and see if we get a collision
			cars[idx] = advance(game, car)
			crashed := detectCollision(cars, cars[idx])
			if crashed >= 0 {
				return nil
			}
		}
		tick++
		if tick >= 5000 {
			break
		}
	}
	return nil
}

func part2(game [][]string, cars []car) error {
	// dumpGame(game, cars)
	tick := 0
	for {
		skip := []int{}
		// if tick > 5255 && tick < 5258 {
		// 	dumpGame(game, cars)
		// }
		// need to sort cars based upon top to bottom and left to right
		sort.SliceStable(cars, func(i, j int) bool {
			if cars[i].loc.Y < cars[j].loc.Y {
				return true
			} else if cars[i].loc.Y == cars[j].loc.Y {
				return cars[i].loc.X < cars[j].loc.X
			}
			return false
		})
	carloop:
		for idx := 0; idx < len(cars); idx++ {
			for _, s := range skip {
				if idx == s {
					continue carloop
				}
			}
			// update each car and see if we get a collision
			cars[idx] = advance(game, cars[idx])
			crashed := detectCollision(cars, cars[idx])

			// could have a problem if 3 cars land on same spot but don't believe this is a case
			if crashed != -1 {
				skip = append(skip, idx, crashed)
				// fmt.Printf("tick %d\n", tick)
				// fmt.Printf(" first: %d  %v  %v\n", cars[idx].id, cars[idx].direction, cars[idx].loc)
				// fmt.Printf("second: %d  %v  %v\n", cars[crashed].id, cars[crashed].direction, cars[crashed].loc)
				// dumpGame(game, cars)
			}
		}

		if len(skip) > 0 {
			// remove the two cars that crashed
			sort.Ints(skip) // sort them and then remove them from highest to lowest
			for r := len(skip) - 1; r >= 0; r-- {
				fmt.Printf("Removing car: %d\n", cars[skip[r]].id)
				if skip[r] == len(cars)-1 {
					cars = cars[:len(cars)-1]
				} else {
					cars = append(cars[:skip[r]], cars[skip[r]+1:]...) // collision
				}
			}
		}

		if len(cars) == 1 {
			fmt.Printf("Last card location (%d, %d)\n", cars[0].loc.X, cars[0].loc.Y)
			return nil
		}

		tick++
		if tick >= 15000 {
			panic("Too far")
		}
	}
	return nil
}

func main() {
	dataFileName := "../../data/day13/input.txt"
	// dataFileName := "../../data/day13/test.1.txt"
	// dataFileName := "../../data/day13/test.txt"

	// build turns ring
	d := []string{"<", "^", ">", "v"}
	turns = ring.New(1)
	for i := 0; i < len(d); i++ {
		turns.Value = d[i]
		if i < len(d)-1 {
			t := ring.New(1)
			turns.Link(t)
		}
		turns = turns.Next()
	}

	// turns.Do(func(p interface{}) { fmt.Println(p) })
	game, cars, err := getData(dataFileName)
	if err != nil {
		panic(err)
	}

	// part1(game, cars)
	part2(game, cars)
}
