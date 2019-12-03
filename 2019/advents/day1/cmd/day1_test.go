package main

import "testing"

func TestDay1Part1(t *testing.T) {
	tests := []struct {
		data   int
		result int
	}{
		{
			data:   12,
			result: 2,
		},
		{
			data:   14,
			result: 2,
		},
		{
			data:   1969,
			result: 654,
		},
		{
			data:   100756,
			result: 33583,
		},
	}

	for _, test := range tests {
		res := Part1(test.data)
		if res != test.result {
			t.Errorf("Test{%d} expected {%d} but ended with {%d}\n", test.data, test.result, res)
		}
	}
}

func TestDay1Part2(t *testing.T) {
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
