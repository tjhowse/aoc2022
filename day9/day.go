package main

import (
	"math"

	tj "github.com/tjhowse/tjgo"
)

const MapSize = 6

type vec2 struct {
	x int
	y int
}

func (v *vec2) distance(v2 *vec2) float64 {
	return math.Sqrt(math.Pow(math.Abs(float64(v.x-v2.x)), 2) + math.Pow(math.Abs(float64(v.y-v2.y)), 2))
}

type snake struct {
	head vec2
	tail vec2
}

func (s *snake) move(dir string, dist int, tp *map[vec2]bool) {
	// println("move ", dir, " ", dist, " ")
	for i := 0; i < dist; i++ {
		head_start := s.head
		switch dir {
		case "U":
			s.head.y++
		case "D":
			s.head.y--
		case "L":
			s.head.x--
		case "R":
			s.head.x++
		}
		// println(s.head.distance(&s.tail))
		if s.head.distance(&s.tail) > 1.5 {
			s.tail = head_start
		}
		(*tp)[s.tail] = true
		// println(s.head.distance(&s.tail))
	}

}

func print_snakemap(s *snake, sm *[MapSize][MapSize]string) {
	for y := MapSize - 1; y >= 0; y-- {
		for x := 0; x < MapSize; x++ {
			char := "."
			if x == s.head.x && y == s.head.y {
				char = "H"
			} else if x == s.tail.x && y == s.tail.y {
				char = "T"
			}
			print(char)
		}
		println()
	}
}

func main() {
	b := tj.FileTo2DSlice("input_real", ' ')
	// b := tj.FileTo2DSlice("input", ' ')

	snakemap := [MapSize][MapSize]string{}
	for i := 0; i < MapSize; i++ {
		for j := 0; j < MapSize; j++ {
			snakemap[i][j] = "."
		}
	}
	s := snake{vec2{0, 0}, vec2{0, 0}}

	// print_snakemap(&s, &snakemap)

	tail_positions := map[vec2]bool{}

	for _, line := range b {
		s.move(line[0], tj.Str2int(line[1]), &tail_positions)
		// print_snakemap(&s, &snakemap)
	}

	// Count the number of visited tail positions
	count := 0
	for _, v := range tail_positions {
		if v {
			count++
		}
	}
	println(count)
}
