package main

import (
	"math"

	tj "github.com/tjhowse/tjgo"
)

type computer struct {
	cycles int
	x      int
	signal int
	screen [6][40]string
}

// Returns true if i is equal to 20, 60, 100, 140, 180 or 220
// Else returns false
func check(i int) bool {
	if i == 20 || i == 60 || i == 100 || i == 140 || i == 180 || i == 220 {
		return true
	}
	return false
}

func (c *computer) clock(n int) {
	for i := 0; i < n; i++ {
		// Drawing pixels
		x := c.cycles % 40
		y := c.cycles / 40
		if math.Abs(float64(c.x-x)) < 2 {
			c.screen[y][x] = "#"
		} else {
			c.screen[y][x] = " "
		}

		c.cycles++
		if check(c.cycles) {
			c.signal += c.x * c.cycles
		}
	}
}

func (c *computer) cmd(i string, n int) {
	if i == "noop" {
		c.clock(1)
	} else if i == "addx" {
		c.clock(2)
		c.x += n
	}
}

func main() {
	// b := tj.FileTo2DSlice("input_real", ' ')
	// b := tj.FileTo2DSlice("input", ' ')
	b := tj.FileTo2DSlice("input_real", ' ')
	// b := tj.FileToSlice("input")

	c := computer{x: 1}
	for _, line := range b {
		if len(line) == 1 {
			c.cmd(line[0], 0)
		} else {
			c.cmd(line[0], tj.Str2int(line[1]))
		}
	}
	println("Cycles:", c.cycles)
	println("     x:", c.x)
	println("Signal:", c.signal)
	for _, line := range c.screen {
		for _, char := range line {
			print(char)
		}
		println()
	}

}
