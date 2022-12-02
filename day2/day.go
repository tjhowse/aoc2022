package main

import (
	tj "github.com/tjhowse/tjgo"
)

func score(them, me string) int {
	// A, X  Rock
	// B, Y  Paper
	// C, Z  scissors
	score := 0
	if me == "X" {
		score += 1
	} else if me == "Y" {
		score += 2
	} else {
		score += 3
	}
	if (them == "A" && me == "X") || (them == "B" && me == "Y") || (them == "C" && me == "Z") {
		score += 3
	} else if them == "A" && me == "Y" {
		score += 6
	} else if them == "B" && me == "Z" {
		score += 6
	} else if them == "C" && me == "X" {
		score += 6
	}
	return score
}

func strat(them, result string) string {
	if result == "X" {
		// Lose
		if them == "A" {
			return "Z"
		} else if them == "B" {
			return "X"
		} else {
			return "Y"
		}
	} else if result == "Y" {
		// Draw
		if them == "A" {
			return "X"
		}
		if them == "B" {
			return "Y"
		}
		if them == "C" {
			return "Z"
		}
	} else {
		// Win
		if them == "A" {
			return "Y"
		} else if them == "B" {
			return "Z"
		} else {
			return "X"
		}
	}
	return "X"
}

func main() {
	// b := tj.FileTo2DSlice("input", ' ')
	b := tj.FileTo2DSlice("input_real", ' ')

	total := 0

	for _, line := range b {
		// Modify strat
		line[1] = strat(line[0], line[1])
		d := score(line[0], line[1])
		// println(d)
		total += d
	}
	println(total)
}
