package main

import (
	"fmt"
	"strings"

	tj "github.com/tjhowse/tjgo"
)

type line struct {
	s tj.Vec2
	e tj.Vec2
}

func (l *line) load(start string, end string) {
	s := strings.Split(start, ",")
	l.s = tj.Vec2{X: tj.Str2int(s[0]), Y: tj.Str2int(s[1])}
	e := strings.Split(end, ",")
	l.e = tj.Vec2{X: tj.Str2int(e[0]), Y: tj.Str2int(e[1])}
	if l.s.X > l.e.X {
		l.s.X, l.e.X = l.e.X, l.s.X
	}
	if l.s.Y > l.e.Y {
		l.s.Y, l.e.Y = l.e.Y, l.s.Y
	}
}

// This returns a string representation of the line
func (l *line) String() string {
	return fmt.Sprintf("%s %s", l.s.String(), l.e.String())
}

type cave struct {
	c      map[tj.Vec2]string
	min    tj.Vec2
	max    tj.Vec2
	abyssY int
}

func (c *cave) Print() {
	for y := c.min.Y; y <= c.max.Y; y++ {
		for x := c.min.X; x <= c.max.X; x++ {
			if _, ok := c.c[tj.Vec2{X: x, Y: y}]; !ok {
				c.c[tj.Vec2{X: x, Y: y}] = "."
			}
			print(c.c[tj.Vec2{X: x, Y: y}])
		}
		println()
	}
	println()
}

func (c *cave) AddSand() bool {
	if c.c[tj.Vec2{X: 500, Y: 0}] == "+" {
		return false
	}
	c.c[tj.Vec2{X: 500, Y: 0}] = "+"
	c.min.Y = 0
	return true
}
func (c *cave) AddSandP2() bool {
	if c.c[tj.Vec2{X: 500, Y: 0}] == "+" {
		return false
	}
	c.min.Y = 0
	// c.c[tj.Vec2{X: 500, Y: 0}] = "+"
	pos := tj.Vec2{X: 500, Y: 0}
	for {
		temp_pos := pos
		temp_pos.Y++
		if t, ok := c.c[temp_pos]; !ok || t == "." {
			// Fall straight down
			pos = temp_pos
			continue
		}
		temp_pos.X--
		if t, ok := c.c[temp_pos]; !ok || t == "." {
			// Fall down left
			pos = temp_pos
			continue
		}
		temp_pos.X += 2
		if t, ok := c.c[temp_pos]; !ok || t == "." {
			// Fall down right
			pos = temp_pos
			continue
		}
		// Couldn't fall.
		c.c[pos] = "+"
		break
	}
	return true
}

func (c *cave) UpdateMinMax() {
	for k := range c.c {
		if k.X < c.min.X {
			c.min.X = k.X
		}
		if k.X > c.max.X {
			c.max.X = k.X
		}
		if k.Y < c.min.Y {
			c.min.Y = k.Y
		}
		if k.Y > c.max.Y {
			c.max.Y = k.Y
		}
	}
}

func (c *cave) Tick() bool {
	moved := false
	for y := c.min.Y; y <= c.max.Y; y++ {
		for x := c.min.X; x <= c.max.X; x++ {
			if c.c[tj.Vec2{X: x, Y: y}] == "+" {
				// Check if it can fall 0, -1
				if t, ok := c.c[tj.Vec2{X: x, Y: y + 1}]; !ok || t == "." {
					// Fall straight down
					c.c[tj.Vec2{X: x, Y: y + 1}] = "+"
					c.c[tj.Vec2{X: x, Y: y}] = "."
					moved = true
					continue
				}
				// Check if it can fall -1, -1
				if t, ok := c.c[tj.Vec2{X: x - 1, Y: y + 1}]; !ok || t == "." {
					// Fall left
					c.c[tj.Vec2{X: x - 1, Y: y + 1}] = "+"
					c.c[tj.Vec2{X: x, Y: y}] = "."
					moved = true
					continue
				}
				// Check if it can fall 1, -1
				if t, ok := c.c[tj.Vec2{X: x + 1, Y: y + 1}]; !ok || t == "." {
					// Fall right
					c.c[tj.Vec2{X: x + 1, Y: y + 1}] = "+"
					c.c[tj.Vec2{X: x, Y: y}] = "."
					moved = true
					continue
				}
			}
		}
	}
	// This is needed for part 1
	if moved {
		c.UpdateMinMax()
		if c.max.Y > c.abyssY {
			return false
		}
	}
	return moved
}

func main_p1() {

	// b := tj.FileToSlice("input")
	b := tj.FileToSlice("input_real")
	// c := FileTo2DSliceRegex("input", " ")

	// packets := make([]packet, 0)
	lines := make([]line, 0)
	for _, row := range b {
		s := strings.Split(row, " ")
		for i := 0; i < len(s)-1; i += 2 {
			l := line{}
			l.load(s[i], s[i+2])
			lines = append(lines, l)
		}
	}
	c := cave{}
	c.min.X = 100000
	c.min.Y = 100000
	c.max.X = -100000
	c.max.Y = -100000
	c.c = make(map[tj.Vec2]string)
	c.min = tj.Vec2{X: 100000, Y: 100000}
	c.max = tj.Vec2{X: -100000, Y: -100000}
	for _, line := range lines {
		for y := line.s.Y; y <= line.e.Y; y++ {
			for x := line.s.X; x <= line.e.X; x++ {
				c.c[tj.Vec2{X: x, Y: y}] = "#"
			}
		}
	}
	c.UpdateMinMax()
	c.abyssY = c.max.Y
	rest := 0
	for {
		// c.Print()
		c.AddSand()
		for c.Tick() {
		}
		rest++
		if c.max.Y > c.abyssY {
			break
		}
	}
	println("Rest:", rest-1)
}

// P{art 2
func main() {

	// b := tj.FileToSlice("input")
	b := tj.FileToSlice("input_real")
	// c := FileTo2DSliceRegex("input", " ")

	// packets := make([]packet, 0)
	lines := make([]line, 0)
	for _, row := range b {
		s := strings.Split(row, " ")
		for i := 0; i < len(s)-1; i += 2 {
			l := line{}
			l.load(s[i], s[i+2])
			lines = append(lines, l)
		}
	}
	c := cave{}
	c.c = make(map[tj.Vec2]string)
	c.min = tj.Vec2{X: 100000, Y: 100000}
	c.max = tj.Vec2{X: -100000, Y: -100000}
	for _, line := range lines {
		for y := line.s.Y; y <= line.e.Y; y++ {
			for x := line.s.X; x <= line.e.X; x++ {
				c.c[tj.Vec2{X: x, Y: y}] = "#"
			}
		}
	}
	c.UpdateMinMax()
	// lines = append(lines, line{s: tj.Vec2{X: -n, Y: c.max.Y + 2}, e: tj.Vec2{X: n, Y: c.max.Y + 2}})
	// l := line{}
	// l.load(fmt.Sprintf("%d,%d", -n, c.max.Y+2), fmt.Sprintf("%d,%d", n, c.max.Y+2))
	// lines = append(lines, l)

	n := 30000
	for x := -n; x < n; x++ {
		v := tj.Vec2{X: x + ((c.max.X-c.min.X)/2 + c.min.X), Y: c.max.Y + 2}
		c.c[v] = "#"
		// println("Adding", v.String())
	}

	c.UpdateMinMax()
	// c.abyssY = c.max.Y
	rest := 0
	// for i := 0; i < 10; i++ {
	for {
		// c.Print()
		// break
		if !c.AddSandP2() {
			break
		}
		// for c.Tick() {
		// }
		rest++
	}
	println("Rest:", rest)
}
