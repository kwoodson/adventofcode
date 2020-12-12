package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Group struct {
	report [][]string
}

func getData() ([]Group, error) {
	dataFileName := "../data/input.txt"
	groups := make([]Group, 0)

	alldata, err := ioutil.ReadFile(dataFileName)
	if err != nil {
		return nil, err
	}

	for _, p := range strings.Split(string(alldata), "\n\n") {
		g := Group{}
		g.report = make([][]string, 0)
		for _, z := range strings.Split(p, "\n") {
			row := make([]string, 0)
			row = append(row, strings.Split(z, "")...)
			g.report = append(g.report, row)
		}
		groups = append(groups, g)
	}
	return groups, nil
}

// Part1 compute result
func Part1(report []Group) int {
	// load each group into a set and count results
	count := 0
	// each group in report
	for _, group := range report {
		results := make(map[string]int)
		// each person in group
		for _, person := range group.report {
			for _, choice := range person {
				results[choice] = 1
			}
		}
		for _, k := range results {
			count += k
		}
	}
	return count
}

// Part1 compute result
func Part2(report []Group) int {
	// load each group into a set and count results
	count := 0
	// each group in report
	for _, group := range report {
		results := make(map[string]int)
		// each person in group
		for _, person := range group.report {
			for _, choice := range person {
				results[choice] += 1
			}
		}
		for _, v := range results {
			if v == len(group.report) {
				count += 1
			}
		}
	}
	return count
}

func main() {
	// test data
	report, err := getData()
	if err != nil {
		panic(err)
	}

	fmt.Println(Part1(report))
	// 11810 too high
	fmt.Println(Part2(report))
}
