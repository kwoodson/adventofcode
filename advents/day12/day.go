package main

import (
	"bufio"
	ring "container/ring"
	"fmt"
	"os"
	"strings"
)

func addDots(r *ring.Ring, size int) *ring.Ring {
	for i := 0; i < size; i++ {
		n := ring.New(1)
		n.Value = "."
		r.Link(n)
		r = r.Next()
	}
	return r
}

func getData(fileName string) (*ring.Ring, []string, error) {
	r := ring.New(1)
	r.Value = "."
	r = addDots(r, 2)
	n := ring.New(1)
	r.Link(n)
	r = r.Next()
	f, err := os.Open(fileName)
	if err != nil {
		return r, nil, err
	}

	rules := make([]string, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "initial") {
			parts := strings.Split(line, ": ")
			for idx, part := range parts[1] {
				r.Value = string(part)
				if idx == len(parts[1])-1 {
					break
				}
				next := ring.New(1)
				r.Link(next)
				r = r.Next()
			}
		} else if line != "" {
			if strings.HasSuffix(line, "#") {
				parts := strings.Split(line, " =>")
				rules = append(rules, parts[0])
			}
		}
	}

	// move us back to the beginning
	r = r.Next()
	r = addDots(r, 2000) // not sure why we add dots on end??
	r.Do(func(p interface{}) {
		if p != nil {
			fmt.Printf("%v ", p)
		} else {
			fmt.Println("FOUND NIL")
		}
	})
	fmt.Println()

	return r, rules, nil
}

func doGeneration(pots *ring.Ring, rules []string) *ring.Ring {
	nextGen := ring.New(1)
	for i := 0; i < pots.Len(); i++ {
		curPat := pots.Prev().Prev().Value.(string) + pots.Prev().Value.(string) + pots.Value.(string) +
			pots.Next().Value.(string) + pots.Next().Next().Value.(string)

		for _, rule := range rules {
			// check to see if current matches
			if strings.Compare(curPat, rule) == 0 {
				nextGen.Value = "#"
				break
			}
		}
		if nextGen.Value != "#" {
			nextGen.Value = "."
		}
		if i != pots.Len()-1 {
			n := ring.New(1)
			nextGen.Link(n)
			nextGen = nextGen.Next()
		}
		pots = pots.Next()
	}
	nextGen = nextGen.Next()
	pots = nextGen

	return pots
}

func part1(pots *ring.Ring, rules []string) error {
	last := 0
	for i := 1; i <= 200; i++ {
		pots = doGeneration(pots, rules)
		count := 0
		for i := 0; i < pots.Len(); i++ {
			if pots.Value == "#" {
				count += i - 3
			}
			pots = pots.Next()
		}
		fmt.Printf("i: %d  Count: %d  diff: %d\n", i, count, count-last)
		last = count
	}

	// manual intervention for part 2
	// look at output and see when the output stabilizes, this happens at around 184
	// calculate 200 generations which then changes by 194 each iteration
	// Take (50_000_000_000 - iterations (200)) * difference_per_iteration (194) + last (count at end of 200 iterations)
	fmt.Println((50000000000-200)*194 + last)
	return nil
}

func main() {
	dataFileName := "../../data/day12/input.txt"
	// dataFileName := "../../data/day12/test.txt"

	initial, rules, err := getData(dataFileName)
	if err != nil {
		panic(err)
	}

	part1(initial, rules)
}
