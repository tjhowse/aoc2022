package main

import (
	tj "github.com/tjhowse/tjgo"
)

func check_visibility(treemap [][]int, i, j int) bool {
	// Edge always visible
	if i == 0 || j == 0 || i == len(treemap)-1 || j == len(treemap[i])-1 {
		return true
	}

	visible := false
	// check left
	all_lower := true
	for k := 0; k < j; k++ {
		if treemap[i][k] >= treemap[i][j] {
			all_lower = false
		}
	}
	if all_lower {
		visible = true
	}
	all_lower = true
	// check Right
	for k := j + 1; k < len(treemap[i]); k++ {
		if treemap[i][k] >= treemap[i][j] {
			all_lower = false
		}
	}
	if all_lower {
		visible = true
	}
	all_lower = true
	// check up
	for k := 0; k < i; k++ {
		if treemap[k][j] >= treemap[i][j] {
			all_lower = false
		}
	}
	if all_lower {
		visible = true
	}
	all_lower = true
	// check down
	for k := i + 1; k < len(treemap); k++ {
		if treemap[k][j] >= treemap[i][j] {
			all_lower = false
		}
	}
	if all_lower {
		visible = true
	}
	all_lower = true
	return visible
}

func check_visibility2(treemap [][]int, i, j int) int {

	result := make([]int, 0)
	// check left
	for k := j - 1; k >= 0; k-- {
		if treemap[i][k] >= treemap[i][j] || k == 0 {
			result = append(result, j-k)
			// println("left", j-k)
			break
		}
	}
	// check Right
	for k := j + 1; k < len(treemap[i]); k++ {
		if treemap[i][k] >= treemap[i][j] || k == len(treemap[i])-1 {
			result = append(result, k-j)
			// println("Right", k-j)
			break
		}
	}
	// check up
	for k := i - 1; k >= 0; k-- {
		if treemap[k][j] >= treemap[i][j] || k == 0 {
			result = append(result, i-k)
			// println("Up", i-k)
			break
		}
	}
	// check down
	for k := i + 1; k < len(treemap); k++ {
		if treemap[k][j] >= treemap[i][j] || k == len(treemap)-1 {
			result = append(result, k-i)
			// println("down", k-i)
			break
		}
	}
	total := 1
	for _, i := range result {
		// println(i)
		total *= i
	}
	return total
}

func main() {
	// b := tj.FileToSlice("input")
	b := tj.FileToSlice("input_real")

	treemap := make([][]int, 0)
	for i, line := range b {
		// println(line)
		treemap = append(treemap, make([]int, 0))
		for _, tree := range line {
			n := tj.Str2int(string(tree))
			// print(n)
			treemap[i] = append(treemap[i], n)
		}
	}
	// println("Tree:", treemap[3][2])
	// println(check_visibility2(treemap, 3, 2))
	total := 0
	for i, row := range treemap {
		for j, _ := range row {
			// print(tree)
			// if check_visibility(treemap, i, j) {
			// 	total++
			// 	print("X")
			// }
			// print("  ")
			a := check_visibility2(treemap, i, j)
			if a > total {
				total = a
			}
		}
		println()
	}
	println(total)
	// c := tj.DoRegex("123-45", `(\d+)-(\d+)`)
	// for _, i := range c {
	// 	println(i)
	// }
	// return
	// for r := range b[0] {
	// 	println(string(rune(r)))
	// }
	// input := b[0]

}
