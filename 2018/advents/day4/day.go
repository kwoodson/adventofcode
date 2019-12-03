package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	BEGIN  recordType = 0
	WAKEUP recordType = 1
	ASLEEP recordType = 2
)

type recordType int32

type record struct {
	date       time.Time
	data       string
	rt         recordType
	owner      int
	cal        []int
	allRecords []record
}

type ByDate []record

func (r ByDate) Len() int           { return len(r) }
func (r ByDate) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r ByDate) Less(i, j int) bool { return r[i].date.Before(r[j].date) }

func getData(fileName string) ([]record, error) {
	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	records := make([]record, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {

		line := scanner.Text()
		// get date
		data := strings.Split(line, "]")
		date := strings.TrimPrefix(data[0], "[")
		/*
			allDate := strings.Split(date, " ")
			ymd := strings.Split(allDate[0], "-")
			hm := strings.Split(allDate[1], ":")

			y, err := strconv.ParseUInt(ymd[0], 10, 32)
			if err != nil {
				return nil, err
			}
			mo, err := strconv.ParseUInt(ymd[1], 10, 32)
			if err != nil {
				return nil, err
			}
			d, err := strconv.ParseUInt(ymd[2], 10, 32)
			if err != nil {
				return nil, err
			}
			h, err := strconv.ParseUInt(hm[0], 10, 32)
			if err != nil {
				return nil, err
			}
			mi, err := strconv.ParseUInt(hm[1], 10, 32)
			if err != nil {
				return nil, err
			}
		*/
		var rt recordType
		var id int
		et, err := time.Parse("2006-01-02 15:04", date)
		if err != nil {
			return nil, err
		}
		if strings.Contains(data[1], "Guard #") {
			rt = BEGIN
			tids := strings.Split(strings.Split(data[1], "#")[1], " ")
			tid, err := strconv.ParseUint(tids[0], 10, 16)
			if err != nil {
				return nil, err
			}
			id = int(tid)
		} else if strings.Contains(data[1], "wakes") {
			rt = WAKEUP
		} else if strings.Contains(data[1], "asleep") {
			rt = ASLEEP
		}

		r := record{
			date:  et,
			data:  data[1],
			rt:    rt,
			owner: id,
		}

		records = append(records, r)

	}
	sort.Sort(ByDate(records))

	return records, nil
}

func associateRecords(records []record) ([]record, map[int][]record) {
	recordsByOwner := make(map[int][]record)
	var currentGuard int
	for i := 0; i < len(records); i++ {
		if currentGuard < 1 && records[i].rt != BEGIN {
			panic("Could not find an associated gaurd")
		}
		if records[i].rt == BEGIN {
			currentGuard = records[i].owner
			recordsByOwner[records[i].owner] = append(recordsByOwner[records[i].owner], records[i])
			continue
		}
		records[i].owner = currentGuard
		recordsByOwner[records[i].owner] = append(recordsByOwner[records[i].owner], records[i])
	}

	return records, recordsByOwner
}

func recordSleep(records []record, sleep []int) []int {
	// current := records[0].date.Minute()
	// walk each minute and check next record to see
	for i := 1; i < len(records)-1; i += 2 {
		asleep := records[i].date.Minute()
		wake := records[i+1].date.Minute()
		for j := asleep; j < wake; j++ {
			sleep[j]++
		}
	}

	return sleep
}

func processSleep(recordMap map[int][]record) int {
	// recordsByOwner := make(map[int][]record)
	sleepMap := make(map[int][]int, 25)

	for owner, records := range recordMap {

		if len(sleepMap[owner]) == 0 {
			sleepMap[owner] = make([]int, 60)
		}

		//
		cap := false
		begin := 0
		for idx, r := range records {
			if r.rt == BEGIN {

				// first time through
				if begin == 0 && !cap {
					cap = true
					continue
				}
				cap = false
				// call processSleep and move to next
				sleepMap[r.owner] = recordSleep(records[begin:idx], sleepMap[owner])
				// sleepMap[r.owner] = append(sleepMap[r.owner], sleep)
				begin = idx
			}
		}
	}

	// TODO: analyze each row and determine who sleeps the most and which one overlaps
	// return it!
	most := 0
	freq := 0
	sleepTime := make(map[int][]int, len(sleepMap))
	for owner, sleep := range sleepMap {
		if most == 0 {
			most = owner
		}
		if freq == 0 {
			freq = owner
		}
		sleepTime[owner] = make([]int, 4)
		count := 0
		high := 0
		minute := 0
		total := 0
		for i := 0; i < len(sleep); i++ {
			// fmt.Printf("i=%d ", i)
			total += sleep[i]
			if sleep[i] != 0 {
				count++
			}
			if sleep[i] > high {
				high = sleep[i]
				minute = i
			}
		}
		sleepTime[owner][0] = count
		sleepTime[owner][1] = high
		sleepTime[owner][2] = minute
		sleepTime[owner][3] = total
		// fmt.Printf("\n[%d] Total[%d] - High[%d] Min[%d] Count[%d] %v\n", owner, total, high, minute, count, sleep)
		// part 1 answer
		if sleepTime[owner][3] >= sleepTime[most][3] {
			most = owner
		}
		// part 2 answer
		if sleepTime[owner][1] >= sleepTime[freq][1] {
			freq = owner
		}
	}
	// PART 1
	fmt.Printf("Part 1 ID: %d\n", most)
	fmt.Println(most * sleepTime[most][2])
	fmt.Printf("Part 2 ID: %d\n", freq)
	fmt.Println(freq * sleepTime[freq][2])
	return 0

}

func part1(records map[int][]record) int {
	result := processSleep(records)
	return result
}

func main() {

	dataFileName := "../../data/day4/sorted.txt"
	// dataFileName := "../../data/day4/input.txt"

	records, err := getData(dataFileName)
	if err != nil {
		panic(err)
	}
	records, recordMap := associateRecords(records)
	part1(recordMap)
}
