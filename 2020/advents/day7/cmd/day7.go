package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type Bag struct {
	color string
	bags  map[string]int
}

type Sack struct {
	color  string
	count  int
	parent string
}

type Node struct {
	name     string
	value    int
	children []*Node
}

func getData() (map[string]*Bag, error) {
	dataFileName := "../data/input.txt"
	bags := make(map[string]*Bag)

	alldata, err := ioutil.ReadFile(dataFileName)
	if err != nil {
		return nil, err
	}
	re := regexp.MustCompile(`(?P<Count>\d) (?P<Bag>\w*\s\w*) bags?`)
	for _, p := range strings.Split(string(alldata), "\n") {

		s := strings.Split(p, "contain")
		b := Bag{
			color: s[0][:len(s[0])-6],
		}
		b.bags = make(map[string]int)
		for _, ss := range strings.Split(s[1], ",") {
			for _, v := range re.FindAllStringSubmatch(ss, 1) {
				num, err := strconv.Atoi(string(v[1]))
				if err != nil {
					return nil, err
				}
				b.bags[string(v[2])] = num
				// fmt.Println(z[i][1],
				// row = append(row, strings.Split(z, "")...)
			}
		}
		c := &b
		bags[c.color] = c
	}
	return bags, nil
}

func search(bags map[string]*Bag, colors map[string]int) map[string]int {
	colorsFound := make(map[string]int)
	for bag_name := range bags {
		b := bags[bag_name]
		for c := range colors {
			if v, ok := b.bags[c]; ok {
				colorsFound[b.color] = v * colors[c]
			}
		}
	}
	return colorsFound
}

// Part1 compute result
func Part1(bags map[string]*Bag) int {
	// load each group into a set and count results
	colorsFound := make(map[string]int)
	colorsSearch := make(map[string]int)
	colorsSearch["shiny gold"] = 0
	// each bag type check to see if it can hold a shiny gold bag
	for {
		// find all bags that have a color
		colorsSearch = search(bags, colorsSearch)
		for c := range colorsSearch {
			if _, ok := colorsFound[c]; ok {
				colorsFound[c]++
				delete(colorsSearch, c)
			} else {
				colorsFound[c] = 0
			}
		}
		if len(colorsSearch) == 0 {
			break
		}
	}

	return len(colorsFound)
}

func getChildren(allbags map[string]*Bag, color string, m int) []*Sack {
	children := make([]*Sack, 0)
	for c, v := range allbags[color].bags {
		fmt.Printf("\tAdding CHILD: Color: %v\tParent: %v\tm: %v\n", c, color, v*m)
		children = append(children, &Sack{
			color:  c,
			count:  v * m,
			parent: color,
		})
	}
	return children
}

func Part2(allbags map[string]*Bag) int {
	count := 0
	processColors := make([]*Sack, 0)
	next := make([]*Sack, 0)
	processColors = append(processColors, &Sack{
		// color:  "shiny gold",
		color:  "shiny gold",
		count:  1,
		parent: "",
	})

	for {
		fmt.Printf("Length of processColors: %v\n", len(processColors))
		for _, sack := range processColors {
			if len(allbags[sack.color].bags) == 0 {
				fmt.Printf("\t\t COUNTED: Color: %v\tParent: %v\tm: %v\n", sack.color, sack.parent, sack.count)
				count += sack.count // do we add the last bag as its count?
			} else {
				next = append(next, getChildren(allbags, sack.color, sack.count)...)
			}
		}

		if len(next) == 0 {
			break
		}
		processColors = append(processColors[:0], next[:]...)
		next = next[:0]
	}
	return count
}

func main() {
	// test data
	bags, err := getData()
	if err != nil {
		panic(err)
	}

	// fmt.Println(Part1(bags))
	// too low 1483
	fmt.Println(Part2(bags))
}
