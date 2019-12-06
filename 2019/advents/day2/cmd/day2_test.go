package main

import "testing"

func TestPart1(t *testing.T) {
	tests := []struct {
		data   []int
		result []int
	}{
		/*{
			data:   []int{1, 0, 0, 0, 99},
			result: []int{2, 0, 0, 0, 99},
		},*/
		{
			data:   []int{2, 3, 0, 3, 99},
			result: []int{2, 3, 0, 6, 99},
		},
		{
			data:   []int{2, 4, 4, 5, 99, 0},
			result: []int{2, 4, 4, 5, 99, 9802},
		},
		{
			data:   []int{1, 1, 1, 4, 99, 5, 6, 0, 99},
			result: []int{30, 1, 1, 4, 2, 5, 6, 0, 99},
		},
	}

	for _, test := range tests {
		res := Part1(test.data)
		for i, d := range res {
			if test.result[i] != d {
				t.Errorf("Test{%d} expected {%d} but ended with {%d}\n", test.data, test.result[i], d)
			}
		}
	}
}

/*func TestDay1Part2(t *testing.T) {
	tests := []struct {
		data   int
		result int
	}{
		{
			data:   14,
			result: 2,
		},
		{
			data:   1969,
			result: 966,
		},
		{
			data:   100756,
			result: 50346,
		},
	}

	for _, test := range tests {
		res := Part2(test.data)
		if res != test.result {
			t.Errorf("Test{%d} expected {%d} but ended with {%d}\n", test.data, test.result, res)
		}
	}
}
*/
