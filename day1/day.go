package main

import (
	"sort"

	tj "github.com/tjhowse/tjgo"
)

func main() {
	b := tj.FileToSlice("input_real")
	// b := tj.FileToIntSlice("input")
	// b := tj.FileTo2DIntSlice("input", ' ')

	total := 0
	max := 0
	totals := []int{}

	for _, line := range b {
		if line == "" {
			println(total)
			if total > max {
				max = total
			}
			totals = append(totals, total)
			total = 0
		} else {
			total += tj.Str2int(line)
			// println
		}
	}
	println(max)
	// Sort totals
	// sort.Ints(totals)
	sort.Sort(sort.Reverse(sort.IntSlice(totals)))
	// println(totals[len(totals)-1] + totals[len(totals)-2] + totals[len(totals)-3])
	println(totals[0] + totals[1] + totals[2])

}
