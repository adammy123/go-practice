package penalty

import (
	"testing"
)

// get_best_closing_times("BEGIN Y Y END \nBEGIN N N END")
//
//	should return an array: [2, 0]
//
// get_best_closing_times("BEGIN BEGIN \nBEGIN N N BEGIN Y Y\n END N N END")
//
//	should return an array: [2]
func TestGet_best_closing_times(t *testing.T) {
	tests := []struct {
		input    string
		expected []int
	}{
		{"BEGIN Y Y END \nBEGIN N N END", []int{2, 0}},
		{"BEGIN BEGIN \nBEGIN N N BEGIN Y Y\n END N N END", []int{2}},
	}

	for _, test := range tests {
		result := Get_best_closing_times(test.input)
		if !testEq(result, test.expected) {
			t.Errorf("Loudly(%s) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func testEq(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestFind_best_closing_time(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"Y Y N N", 2},
		{"Y Y Y Y", 4},
		{"N N N N", 0},
		{"Y N Y N", 1},
		{"Y Y Y Y Y", 5},
		{"N N N N N", 0},
		{"N N Y Y", 0},
	}

	for _, test := range tests {
		result := find_best_closing_time(test.input)
		if result != test.expected {
			t.Errorf("Loudly(%s) = %d; want %d", test.input, result, test.expected)
		}
	}
}

// compute_penalty("Y Y N Y", 0) should return 3
// compute_penalty("N Y N Y", 2) should return 2
// compute_penalty("Y Y N Y", 4) should return 1
func TestCompute_penalty(t *testing.T) {
	tests := []struct {
		input    string
		k        int
		expected int
	}{
		{"Y Y N Y", 0, 3},
		{"N Y N Y", 2, 2},
		{"Y Y N Y", 4, 1},
		{"Y Y Y Y", 4, 0},
		{"N N N N", 4, 4},
		{"Y Y Y Y", 0, 4},
		{"N N N N", 0, 0},
	}

	for _, test := range tests {
		result := compute_penalty(test.input, test.k)
		if result != test.expected {
			t.Errorf("Loudly(%s, %d) = %d; want %d", test.input, test.k, result, test.expected)
		}
	}
}
