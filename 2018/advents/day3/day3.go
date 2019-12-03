package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type plan struct {
	ID        int32
	startLeft int32
	startTop  int32
	width     int32
	height    int32
	pattern   []tupple
	overlap   bool
}

type tupple struct {
	row    int32
	column int32
}

func getData(fileName string) ([]plan, error) {
	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	plans := make([]plan, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {

		line := scanner.Text()
		parts := strings.Split(line, " ")
		id, err := strconv.ParseInt(strings.Split(parts[0], "#")[1], 10, 32)
		if err != nil {
			return nil, err
		}
		middle := strings.Split(parts[2], ",")
		left, err := strconv.ParseInt(middle[0], 10, 32)
		if err != nil {
			return nil, err
		}

		top, err := strconv.ParseInt(strings.Split(middle[1], ":")[0], 10, 32)
		if err != nil {
			return nil, err
		}
		wh := strings.Split(parts[3], "x")
		wid, err := strconv.ParseInt(wh[0], 10, 32)
		if err != nil {
			return nil, err
		}
		hei, err := strconv.ParseInt(wh[1], 10, 32)
		if err != nil {
			return nil, err
		}
		tmpPlan := plan{
			ID:        int32(id),
			startTop:  int32(top),
			startLeft: int32(left),
			width:     int32(wid),
			height:    int32(hei),
		}
		tmpPlan.pattern = generatePattern(tmpPlan)
		plans = append(plans, tmpPlan)
	}

	return plans, nil
}

func generatePattern(p plan) []tupple {
	pattern := make([]tupple, 0)
	for i := p.startLeft; i < p.startLeft+p.width; i++ {
		for j := p.startTop; j < p.startTop+p.height; j++ {
			pattern = append(pattern, tupple{
				row:    j,
				column: i,
			})
		}
	}

	return pattern
}

func placePattern(p plan, ws [][]int) [][]int {

	for _, point := range p.pattern {
		ws[point.row][point.column]++
	}

	return ws
}

func part1(plans []plan, workspace [][]int) (int, [][]int) {
	// create a multidemnsional array
	// place down all of the elves patterns 1by1
	// mark +1 at each overlapping sqaure

	for _, plan := range plans {
		workspace = placePattern(plan, workspace)
	}

	overlap := 0
	for i := 0; i < len(workspace)-1; i++ {
		for j := 0; j < len(workspace)-1; j++ {
			if workspace[i][j] > 1 {
				overlap++
			}
		}
	}
	return overlap, workspace
}

func part2(plans []plan, ws [][]int) plan {
	// for each plan
	found := -1
	for idx, plan := range plans {
		// fmt.Printf("Checking plan [%d]\n", plan.ID)
		// check overlap
		for _, point := range plan.pattern {
			if ws[point.row][point.column] > 1 {
				plan.overlap = true
				break
			}
		}
		if !plan.overlap {
			found = idx
			break
		}
	}
	fmt.Println(found)
	return plans[found]
}

func main() {

	dataFileName := "../../data/day3/day_3_input.txt"
	plans, err := getData(dataFileName)
	if err != nil {
		panic(err)
	}
	workspace := make([][]int, 1000)
	for i := range workspace {
		workspace[i] = make([]int, 1000)
	}

	overlap, workspace := part1(plans, workspace)
	fmt.Println(overlap)

	p := part2(plans, workspace)
	fmt.Printf("Plan ID: %d\n", p.ID)

}
