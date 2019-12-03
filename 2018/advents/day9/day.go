package main

import (
	"container/ring"
	"fmt"
)

func part1(end, players int, print bool) error {
	scores := make(map[int]int)
	gr := ring.New(1)
	gr.Value = 0
	warm := []int{2, 1, 3}
	for i := 0; i < len(warm); i++ {
		r := ring.New(1)
		r.Value = warm[i]
		gr = r.Link(gr)
		gr = gr.Next()
	}

	gr = gr.Next()
	for i := 4; i < end; i++ {
		if i%23 == 0 {
			// pop the item here
			gr = gr.Move(-10)
			tmpI := i
			// fmt.Printf("Popping %d  score %d  i=%v  tmpI=%v\n", gr.Next().Value.(int), gr.Next().Value.(int)+tmpI, i, tmpI)
			scores[i%players] += gr.Next().Value.(int) + tmpI
			gr.Unlink(1)
			gr = gr.Move(3)
		} else {
			r := ring.New(1)
			r.Value = i
			gr = r.Link(gr)
			gr = gr.Move(2)
		}

		if print {

			gr.Do(func(p interface{}) {
				if p != nil {
					fmt.Printf("%v ", p)
				}

			})
			fmt.Println()
		}

	}
	scores[end%players] += end
	score := 0
	player := 0
	for key, val := range scores {
		if val > score {
			score = val
			player = key
		}
	}

	fmt.Printf("Player: %d score %d\n", player, scores[player])
	return nil
}

func main() {
	part1(25, 9, true)
	part1(1618, 10, false)
	part1(7999, 13, false)
	part1(1104, 17, false)
	part1(6111, 21, false)
	part1(5807, 30, false)
	part1(71223*100, 455, false)

}
