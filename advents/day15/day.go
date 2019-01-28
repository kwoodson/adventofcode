package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type point struct {
	Y int
	X int
}
type player struct {
	loc         point
	hp          int
	elf         bool
	ID          int
	attackPower int
	opp         string
	self        string
}

func getData(fileName string) ([][]string, []player, error) {

	f, err := os.Open(fileName)
	if err != nil {
		return nil, nil, err
	}
	defer f.Close()

	players := make([]player, 0)
	game := make([][]string, 0)
	scanner := bufio.NewScanner(f)
	lineRow := 0
	playerID := 0
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "")
		row := make([]string, 0)
		row = append(row, parts...)
		game = append(game, row)
		for idx, part := range parts {
			if part == "E" || part == "G" {
				loc := point{
					Y: lineRow,
					X: idx,
				}
				opp := "E"
				if part == "E" {
					opp = "G"
				}
				players = append(players, player{
					loc:         loc,
					hp:          200,
					attackPower: 3,
					elf:         part == "E",
					ID:          playerID,
					opp:         opp,
					self:        part,
				})
				playerID++
				game[lineRow][idx] = "."
			}
		}
		lineRow++
	}

	return game, players, nil
}

func dumpGame(game [][]string, players []player) {
	for y := 0; y < len(game); y++ {
	innerloop:
		for x := 0; x < len(game[y]); x++ {
			for _, player := range players {
				if player.loc.Y == y && player.loc.X == x {
					c := "E"
					if !player.elf {
						c = "G"
					}
					fmt.Printf(c)
					continue innerloop
				}
			}
			fmt.Printf("%s", game[y][x])
		}
		fmt.Println()
	}
}

func findPath(game [][]string, start, end point, path []point, valid [][]point) ([]point, [][]point, bool) {
	if start.X == end.X && start.Y == end.Y {
		return path, nil, true
	}
	// if distance(start, end) == 1 {
	// 	return path, nil, true
	// }

	neighbors := getNeighbors(start, game)

	candidates := make([]point, 0)
	lowest := len(game) * 2
	for _, n := range neighbors {
		dist := distance(n, end)
		if dist <= lowest {
			if dist == lowest {
				candidates = append(candidates, n)
			} else {
				candidates = []point{n}
			}
			lowest = dist
		}
	}
	vpath := []point{}
	// tmpPath := []point{}
	result := false
cloop:
	for _, c := range candidates {
		for _, n := range path {
			if c == n {
				continue cloop // safegaurd against looping over same path
			}
		}
		tmpPath := append(path, c)
		vpath, _, result = findPath(game, c, end, tmpPath, valid)
		if result && len(vpath) > 0 {
			valid = append(valid, vpath)
		}
	}
	if result && len(vpath) > 0 {
		return nil, valid, true
	} else {
		return nil, valid, false
	}
}

func movePlayers(game [][]string, players []player) []player {
	players = sortPlayers(players)
	for cpIndex, player := range players {
		nextTo := false
		for _, p := range players {
			if player.ID == p.ID {
				continue
			}
			if player.opp == p.opp {
				continue
			}
			if player.loc.Y == p.loc.Y && (player.loc.X-1 == p.loc.X || player.loc.X+1 == p.loc.X) {
				nextTo = true
				break
			}
			if player.loc.X == p.loc.X && (player.loc.Y-1 == p.loc.Y || player.loc.Y+1 == p.loc.Y) {
				nextTo = true
				break
			}
		}
		if !nextTo {
			// path := make([]point, 0)
			tmploc := nextMove(player, players, game)
			// _, valid, result := findPath(game, player, opp, path)
			if player.loc != tmploc {
				// move happened, turn over
				players[cpIndex].loc = tmploc // persist move
				continue
			}
			fmt.Printf("player(%d) didn't move!\n", player.ID)
			// panic("DIDNT MOVE!")
		}
	}
	return players
}

func war(players []player, game [][]string) []player {
	// TODO: EVERYONE MAKE THEIR MOVE FIRST, IF MOVED, THEN REMOVE FROM TURN.  IF NOT MOVED, THEN ATTACK.
	// each player needs to take a turn, if a player dies then remove him
	// find the next player in order that hasn't moved
	// MIGHT NEED TO SORT HERE!
	// var cp player
	playerMap := make(map[int]player)
	round := make([]int, 0)
	for _, p := range players {
		round = append(round, p.ID)
		playerMap[p.ID] = p
	}
	action := make([]int, 0)
	killed := make([]int, 0)
loopround:
	for rid, id := range round {
		for _, killid := range killed {
			if id == killid {
				continue loopround
			}
		}
		cp := playerMap[id]
		action = append(action, rid)
		// round = append(round[:rid], round[rid+1:]...) // remove player from round
		// attack here
		attackPlayer, valid := attack(cp, players, game)
		if valid { // no one to attack
			attackPlayer.hp -= cp.attackPower

			if attackPlayer.hp <= 0 {
				if true {
					fmt.Println("I died")
				}
				// remove attackPlayer
				// player died, remove him
				removeID := -1
				for pid, p := range players {
					if attackPlayer.ID == p.ID {
						removeID = pid
						break
					}
				}
				players = append(players[:removeID], players[removeID+1:]...) // remove player from players

				// also remove him from the round
				killed = append(killed, attackPlayer.ID)

			} else {
				playerMap[attackPlayer.ID] = attackPlayer
				for idx, p := range players {
					if attackPlayer.ID == p.ID {
						players[idx] = attackPlayer // persist the hitpoints
						break
					}
				}

			}
		}
	}
	// account for each players action
	sort.Ints(action)
	for i := len(action) - 1; i >= 0; i-- {
		round = append(round[:action[i]], round[action[i]+1:]...)
	}
	// Remove killed players
	for _, kid := range killed {
		for rid, id := range round {
			if kid == id {
				round = append(round[:rid], round[rid+1:]...)
				break
			}
		}
	}
	if len(round) > 0 {
		panic("Someone didn't move")
	}

	for _, player := range players {
		fmt.Printf("AFTER: %v: id(%d) loc(%d, %d) hp(%d)\n", player.self, player.ID, player.loc.Y, player.loc.X, player.hp)
	}
	return players
}

func attack(p player, players []player, game [][]string) (player, bool) {
	var opponent player
	valid := true
	candidates := make([]player, 0)
	lowestHP := 200
	for _, n := range getNeighbors(p.loc, game) {
		for _, opp := range players {
			if p.ID == opp.ID {
				continue
			}
			if p.opp == opp.opp { // same team?
				continue
			}
			if n == opp.loc {
				if opp.hp == lowestHP {
					candidates = append(candidates, opp)
				} else if opp.hp < lowestHP {
					candidates = []player{opp}
					lowestHP = opp.hp
				}
			}
		}
	}
	// more than 1 candidate, sort and then select
	if len(candidates) > 1 {
		// sort and select based upon t->b l->r
		candidates = sortPlayers(candidates)
		opponent = candidates[0]
	} else if len(candidates) == 1 {
		opponent = candidates[0]
	} else {
		// this is probably ok, someone died??
		fmt.Printf("Someone probably died: no candidates to attack for %d at (%d,%d)\n", p.ID, p.loc.Y, p.loc.X)
		valid = false
		// panic("No candidates to attack")
	}

	return opponent, valid
}
func sortPlayers(players []player) []player {
	// must sort each time through since players are moving
	sort.SliceStable(players, func(i, j int) bool {
		if players[i].loc.Y < players[j].loc.Y {
			return true
		} else if players[i].loc.Y == players[j].loc.Y {
			return players[i].loc.X < players[j].loc.X
		}
		return false
	})
	return players
}
func occupied(loc point, players []player) bool {
	for _, player := range players {
		if loc.X == player.loc.X && loc.Y == player.loc.Y {
			return true
		}
	}

	return false
}
func distance(l1 point, l2 point) int {
	// find closest person and a valid path
	x := l2.X - l1.X
	y := l2.Y - l1.Y
	if x < 0 {
		x *= -1
	}
	if y < 0 {
		y *= -1
	}
	return x + y
}

func getNeighbors(loc point, game [][]string) []point {
	maxY := len(game)
	maxX := len(game[loc.Y])
	points := make([]point, 0)
	if loc.X-1 >= 0 && game[loc.Y][loc.X-1] == "." {
		points = append(points, point{X: loc.X - 1, Y: loc.Y})
	}
	if loc.Y-1 >= 0 && game[loc.Y-1][loc.X] == "." {
		points = append(points, point{X: loc.X, Y: loc.Y - 1})
	}
	if loc.X+1 <= maxX && game[loc.Y][loc.X+1] == "." {
		points = append(points, point{X: loc.X + 1, Y: loc.Y})
	}
	if loc.Y+1 <= maxY && game[loc.Y+1][loc.X] == "." {
		points = append(points, point{X: loc.X, Y: loc.Y + 1})
	}
	return points
}

func nextMove(pStart player, players []player, game [][]string) point {
	distanceMap := make(map[point]map[point]int)
	candidates := make([]point, 0)
	lowest := len(game)
	// loop through all of my neighbor positions (starting positions)

	// If the space is a wall/cave or is occupied, skip
	distanceMap[pStart.loc] = make(map[point]int)
	// loop through each player and find an opponent
	for _, cp := range players {
		c := cp.loc
		if pStart.opp == cp.opp {
			continue
		}
		// for each opponent's neighbor, find a valid location in which currentPlayer will attempt to move
		for _, cn := range getNeighbors(c, game) {
			// find the distance between current Player and the opponent's neighbor
			dist := distance(pStart.loc, cn)
			distanceMap[pStart.loc][cn] = dist

			// find the lowest distance for all moves
			if dist <= lowest {
				if dist == lowest {
					candidates = append(candidates, cn) // tied
				} else {
					candidates = []point{cn} // single lowest path
				}
				lowest = dist
			}
		}

	}
	allResults := make(map[string][]point)
	for _, c := range candidates {
		path := []point{}
		valid := make([][]point, 0)
		_, valid, result := findPath(game, pStart.loc, c, path, valid)
		if result {
			fmt.Println(valid)
			// we need to combine all valid results and make a move based upon the shortest *valid* path
			for _, result := range valid {
				allResults[fmt.Sprintf("%v%v", result[0].Y, result[0].X)] = result
			}
		}
	}

	// now loop over allResults and find the shortest distance
	// shortest := []int{}
	candidates = []point{}
	lowest = len(game) * 2
	for _, result := range allResults {
		if len(result) <= lowest {
			if len(result) == lowest {
				candidates = append(candidates, result[0])
			} else {
				candidates = []point{result[0]}
			}
			lowest = len(result)
		}
	}

	sort.SliceStable(candidates, func(i, j int) bool {
		if candidates[i].Y < candidates[j].Y {
			return true
		} else if candidates[i].Y == candidates[j].Y {
			return candidates[i].X < candidates[j].X
		}
		return false
	})
	if len(candidates) == 0 {
		return pStart.loc
	}

	return candidates[0]
}

func continueGame(players []player) bool {
	p1 := players[0]
	for i := 1; i < len(players); i++ {
		if players[i].opp != p1.opp {
			return true
		}
	}

	return false
}

func getHP(players []player) int {
	total := 0
	for _, p := range players {
		total += p.hp
	}
	return total
}

func part1(game [][]string, players []player) error {
	dumpGame(game, players)

	// perform round
	round := 1
	for len(players) != 1 {

		fmt.Printf("ROUND: %d\n", round)
		// move players first
		if round == 24 {
			fmt.Println("24")
		}
		players = movePlayers(game, players)
		players = war(players, game)
		// game over?

		// perform attacks
		fmt.Println()
		fmt.Println()
		dumpGame(game, players)
		fmt.Println()
		fmt.Println()
		fmt.Println()

		if !continueGame(players) {
			// game summary
			hp := getHP(players)
			fmt.Printf("Game ended in %d rounds.  HP: %d  Score: %d", round-1, hp, hp*round-1)
			break
		}
		round++

	}
	return nil
}

func main() {
	// dataFileName := "../../data/day15/input.txt"
	// dataFileName := "../../data/day15/test.txt"
	dataFileName := "../../data/day15/test.1.txt"

	game, players, err := getData(dataFileName)
	if err != nil {
		panic(err)
	}

	part1(game, players)
}
