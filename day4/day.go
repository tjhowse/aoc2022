package main

import (
	"strings"

	tj "github.com/tjhowse/tjgo"
)

// all := make(map[int]int)

// totals := []int{}
func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func in(slice []int, item int) bool {
	for _, i := range slice {
		if i == item {
			return true
		}
	}
	return false
}

// This returns true if every element of int slice a is contained within
// int slice b, or vice versa
func isSubset(a, b []int) bool {
	for _, i := range a {
		if !in(b, i) {
			return false
		}
	}
	return true
}

func isOverlap(a, b []int) bool {
	for _, i := range a {
		if in(b, i) {
			return true
		}
	}
	return false
}

func check_pair_overlap(a, b []string) int {
	a_range := makeRange(tj.Str2int(a[0]), tj.Str2int(a[1]))
	b_range := makeRange(tj.Str2int(b[0]), tj.Str2int(b[1]))

	// Check if a_range is wholly within b_range
	if isSubset(a_range, b_range) {
		return 1
	}
	if isSubset(b_range, a_range) {
		return 1
	}
	return 0
}

func check_overlap(a, b []string) int {
	a_range := makeRange(tj.Str2int(a[0]), tj.Str2int(a[1]))
	b_range := makeRange(tj.Str2int(b[0]), tj.Str2int(b[1]))

	// Check if a_range is wholly within b_range
	if isOverlap(a_range, b_range) {
		return 1
	}
	return 0
}

func main() {
	// b := tj.FileToSlice("input")
	// b := tj.FileToIntSlice("input")
	// b := tj.FileTo2DSlice("input", ',')
	b := tj.FileTo2DSlice("input_real", ',')
	// b := tj.FileTo2DIntSlice("input")

	total := 0

	for _, line := range b {
		// For each pair
		a := strings.Split(line[0], "-")
		b := strings.Split(line[1], "-")
		// total += check_pair_overlap(a, b)
		total += check_overlap(a, b)
	}
	println(total)

}
