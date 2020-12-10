package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type Passport struct {
	byr int
	iyr int
	eyr int
	hgt string
	hcl string
	ecl string
	pid string
	cid int
}

func getData() ([]*Passport, error) {
	dataFileName := "../data/input.txt"
	// f, err := os.Open(dataFileName)
	// if err != nil {
	// 	panic(err)
	// }
	pass := make([]*Passport, 0)

	alldata, err := ioutil.ReadFile(dataFileName)
	if err != nil {
		return nil, err
	}

	// scanner := bufio.NewScanner(f)
	for _, p := range strings.Split(string(alldata), "\n\n") {
		pp := Passport{}
		p := strings.ReplaceAll(p, "\n", " ")
		parts := strings.Split(p, " ")
		for _, part := range parts {
			k_v := strings.Split(part, ":")
			switch v := k_v[0]; v {
			case "byr":
				pp.byr, err = strconv.Atoi(k_v[1])
				if err != nil {
					return nil, err
				}
			case "iyr":
				pp.iyr, err = strconv.Atoi(k_v[1])
				if err != nil {
					return nil, err
				}
			case "eyr":
				pp.eyr, err = strconv.Atoi(k_v[1])
				if err != nil {
					return nil, err
				}
			case "hgt":
				pp.hgt = k_v[1]
			case "hcl":
				pp.hcl = k_v[1]
			case "ecl":
				pp.ecl = k_v[1]
			case "pid":
				pp.pid = k_v[1]
			case "cid":
				pp.cid, err = strconv.Atoi(k_v[1])
				if err != nil {
					return nil, err
				}
			}
		}
		e := pp
		pass = append(pass, &e)
	}

	return pass, nil
}

func countPassports(passports []*Passport) int {
	count := 0
	for _, p := range passports {
		if p.byr == 0 || p.iyr == 0 || p.eyr == 0 || p.hcl == "" || p.ecl == "" || p.pid == "" || p.hgt == "" {
			continue
		} else {
			count++
		}
	}
	return count
}

func validHeight(h string) (bool, error) {
	resplit := regexp.MustCompile(`(\d{2,3})(in|cm)`)
	if !resplit.MatchString(h) {
		return false, nil
	}

	if strings.Contains(h, "in") {
		hs := strings.Split(h, "in")
		hInt, err := strconv.Atoi(hs[0])
		if err != nil {
			return false, err
		}
		if hInt >= 59 && hInt <= 76 {
			return true, nil
		}
	} else if strings.Contains(h, "cm") {
		hs := strings.Split(h, "cm")
		hInt, err := strconv.Atoi(hs[0])
		if err != nil {
			return false, err
		}
		if hInt >= 150 && hInt <= 193 {
			return true, nil
		}
	}

	return false, nil
}

func countValidPassports(passports []*Passport) (int, error) {
	count := 0
	for _, p := range passports {
		height, err := validHeight(p.hgt)
		if err != nil {
			continue
		}
		// results := resplit.FindAllString(p.hgt, 1)
		if p.byr < 1920 || p.byr > 2002 {
			// fmt.Printf("Incorrect: byr: with %v NOT valid\n", p.byr)
		} else if p.iyr < 2010 || p.iyr > 2020 {
			// fmt.Printf("Incorrect: iyr: with %v NOT valid\n", p.iyr)
		} else if p.eyr < 2020 || p.eyr > 2030 {
			// fmt.Printf("Incorrect: eyr: with %v NOT valid\n", p.eyr)
		} else if !height {
			// fmt.Printf("Incorrect: hgt: with %v NOT valid\n", p.hgt)
		} else if !regexp.MustCompile(`^(amb|blu|brn|gry|grn|hzl|oth)$`).MatchString(p.ecl) {
			// fmt.Printf("Incorrect: ecl: with %v NOT valid\n", p.ecl)
		} else if !regexp.MustCompile(`^#[0-9a-f]{6}$`).MatchString(p.hcl) {
			// fmt.Printf("Incorrect: hcl: with %v NOT valid\n", p.hcl)
		} else if !regexp.MustCompile(`^\d{9}$`).MatchString(p.pid) {
			// fmt.Printf("Incorrect: pid: with %v NOT valid\n", p.pid)
		} else {
			count++
		}
	}
	return count, nil
}

// Part1 compute result
func Part1(passports []*Passport) int {
	return countPassports(passports)
}

func Part2(passports []*Passport) (int, error) {
	return countValidPassports(passports)
}

func main() {
	// test data
	p, err := getData()
	if err != nil {
		panic(err)
	}

	fmt.Println(Part1(p))
	// 173 too low
	// 195 too high
	res, err := Part2(p)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
	// fmt.Println(Part2(treeMap))
}
