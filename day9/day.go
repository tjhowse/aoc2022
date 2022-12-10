package main

import (
	"math"

	tj "github.com/tjhowse/tjgo"
)

const MapSize = 21

type vec2 struct {
	x int
	y int
}

func (v *vec2) distance(v2 *vec2) float64 {
	return math.Sqrt(math.Pow(math.Abs(float64(v.x-v2.x)), 2) + math.Pow(math.Abs(float64(v.y-v2.y)), 2))
}
func (v *vec2) distance2(v2 *vec2) int {
	return int(math.Abs(float64(v.x-v2.x)) + math.Abs(float64(v.y-v2.y)))
}

const SnakeLength = 10

type snake struct {
	body [SnakeLength]vec2
}

func (s *snake) move(dir string, dist int, tp *map[vec2]bool) {
	// println("move ", dir, " ", dist, " ")
	for i := 0; i < dist; i++ {
		// prev_segment_original := s.body[0]
		switch dir {
		case "U":
			s.body[0].y++
		case "D":
			s.body[0].y--
		case "L":
			s.body[0].x--
		case "R":
			s.body[0].x++
		}
		for j := 1; j < SnakeLength; j++ {
			// temp := s.body[j]
			dist := s.body[j].distance(&s.body[j-1])
			if dist > 1.5 {
				if s.body[j-1].x > s.body[j].x {
					s.body[j].x++
				} else if s.body[j-1].x < s.body[j].x {
					s.body[j].x--
				}
				if s.body[j-1].y > s.body[j].y {
					s.body[j].y++
				} else if s.body[j-1].y < s.body[j].y {
					s.body[j].y--
				}
			}
			// println(dist)
			// switch int(dist) {
			// case 0:
			// 	// Overlap, Do nothing
			// case 1:
			// 	// Orthogonally or diagonally adjacent, do nothing
			// case 2:
			// 	fallthrough
			// case 3:
			// 	// Horizontally or vertically distant
			// 	println("Horiz/Vert")
			// 	s.body[j] = prev_segment_original
			// case 4:
			// 	// Diagonally distant move diagonally towards body[j-1]
			// 	println("Diag")
			// 	if s.body[j-1].x > s.body[j].x {
			// 		s.body[j].x++
			// 	} else {
			// 		s.body[j].x--
			// 	}
			// 	if s.body[j-1].y > s.body[j].y {
			// 		s.body[j].y++
			// 	} else {
			// 		s.body[j].y--
			// 	}
			// }
			// prev_segment_original = temp
		}
		(*tp)[s.body[SnakeLength-1]] = true
		// println(s.head.distance(&s.tail))
	}

}

func print_snakemap(s *snake, sm *[MapSize][MapSize]string) {
	for y := MapSize - 1; y >= 0; y-- {
		for x := 0; x < MapSize; x++ {
			char := "."
			for i, v := range s.body {
				if x == v.x && y == v.y {
					if i == 0 {
						char = "H"
					} else {
						char = tj.Int2str(i)
					}
				}
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
	s := snake{}

	print_snakemap(&s, &snakemap)

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
