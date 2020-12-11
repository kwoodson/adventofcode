package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

func getData() ([][]string, error) {
	dataFileName := "../data/input.txt"
	seats := make([][]string, 0)

	alldata, err := ioutil.ReadFile(dataFileName)
	if err != nil {
		return nil, err
	}

	for _, p := range strings.Split(string(alldata), "\n") {
		row := make([]string, 0)
		row = append(row, strings.Split(p, "")...)
		seats = append(seats, row)
	}
	return seats, nil
}

func determineSeat(data []string) int {
	low := 0
	dLen := len(data)
	high := int(math.Pow(2, float64(dLen))) - 1
	for i, r := range data {
		m := math.Pow(2, float64(dLen-1-i))
		if r == "F" || r == "L" {
			high = high - int(m)
		} else {
			low = low + int(m)
		}
	}
	return high
}

// Part1 compute result
func Part1(seats [][]string) int {
	highest := 0
	for _, s := range seats {
		r := determineSeat(s[:7])
		c := determineSeat(s[7:])
		id := r*8 + c
		if id > highest {
			highest = id
		}
	}

	return highest
}

func Part2(seats [][]string) int {
	allSeats := make(map[int]struct{})
	for _, s := range seats {
		r := determineSeat(s[:7])
		c := determineSeat(s[7:])
		allSeats[r*8+c] = struct{}{}
	}

	for r := 1; r <= 126; r++ {
		for c := 0; c <= 7; c++ {
			seatID := r*8 + c
			if _, ok := allSeats[seatID]; !ok {
				_, prev := allSeats[seatID-1]
				_, next := allSeats[seatID+1]
				if prev && next {
					return seatID
				}
			}
		}
	}
	return 0
}

func main() {
	// test data
	seats, err := getData()
	if err != nil {
		panic(err)
	}

	fmt.Println(Part1(seats))
	// 735 too low

	fmt.Println(Part2(seats))

}
