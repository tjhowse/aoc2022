package main

import (
	"math"

	tj "github.com/tjhowse/tjgo"
)

func getLowest(openSet *map[[2]int]bool, fScore *map[[2]int]int) [2]int {
	min := 999999999
	lowest := [2]int{}
	for node, val := range *openSet {
		if !val {
			continue
		}
		fScore := getScore(*fScore, node)
		if fScore < min {
			min = fScore
			lowest = node
		}
	}
	return lowest
}

func getScore(m map[[2]int]int, p [2]int) int {
	score, ok := m[p]
	if !ok {
		return 99999999
	}
	return score
}

type heightmap struct {
	m [][]int
}

// Returns the approximate cost of getting from x,y to the destination
func (m *heightmap) heuristic(pos [2]int, end [2]int) int {
	// return int(math.Sqrt(math.Pow(float64(len(m.m[0])-pos[0]), 2)+math.Pow(float64(len(m.m)-pos[1]), 2)))
	return end[0] - pos[0] + end[1] - pos[1]
}
func (m *heightmap) reconstructPath(cameFrom map[[2]int][2]int, current [2]int) [][2]int {

	totalPath := [][2]int{current}
	for {
		_, ok := cameFrom[current]
		if !ok {
			break
		}
		current = cameFrom[current]
		totalPath = append(totalPath, current)
	}
	// current = cameFrom[current]

	// totalPath = append(totalPath, current)
	return totalPath
}

func (m *heightmap) getCost(pos [2]int) int {
	// return m.m[pos[1]][pos[0]]
	return 1
}

func (m *heightmap) checkInBounds(x, y int) bool {
	if x >= len(m.m[0]) || x < 0 || y >= len(m.m) || y < 0 {
		return false
	}
	return true
}

func (m *heightmap) getAdjacent2(pos [2]int) [][2]int {
	adjacent := [][2]int{}
	for _, oY := range []int{-1, 0, 1} {
		for _, oX := range []int{-1, 0, 1} {
			if !(oX == 0 || oY == 0) || ((oX == 0) && (oY == 0)) {
				continue
			}

			dX := pos[0] + oX
			dY := pos[1] + oY
			if !m.checkInBounds(dX, dY) {
				continue
			}
			// Only consider a cell adjacent if its value is no more than one
			// number higher thant he current position.
			destination_value := m.m[dY][dX]
			// println("pos[0]:", pos[0], "pos[1]:", pos[1])
			// if pos[0] >= len(m.m) || pos[0] < 0 || pos[1] >= len(m.m[0]) || pos[1] < 0 {
			// 	continue
			// }
			// println("We're at ", pos[0], pos[1])
			origin_value := m.m[pos[1]][pos[0]]
			if destination_value-origin_value <= 1 || destination_value < 0 || origin_value < 0 {
				adjacent = append(adjacent, [2]int{dX, dY})
				// println("We're at ", pos[0], pos[1], "and adding ", dX, dY, "from ", origin_value, "to ", destination_value)
			} else {
				// println("We're at ", pos[0], pos[1], "and rejecting ", dX, dY, "from ", origin_value, "to ", destination_value)

			}
			// adjacent = append(adjacent, [2]int{dX, dY})
		}
	}
	return adjacent
}

func checkOnPath(x, y int, path [][2]int) bool {
	for _, step := range path {
		if step[0] == x && step[1] == y {
			// This spot it alread on our path, no backtracking!
			return true
		}
	}
	return false
}
func (m *heightmap) printWithPath(path [][2]int) {

	// for y := 0; y < len(m.m); y++ {
	// 	for x := 0; x < len(m.m[0]); x++ {
	for y := 0; y < int(math.Min(50, float64(len(m.m)))); y++ {
		for x := 0; x < int(math.Min(50, float64(len(m.m[0])))); x++ {
			if checkOnPath(x, y, path) {
				print("+")
			} else {
				print(m.m[y][x])
			}
		}
		println()
	}
	println("--------------------")
}

func (m *heightmap) aStar(start [2]int, goal [2]int) int {
	// start := [2]int{0, 0}
	// goal := [2]int{len(m.m[0]) - 1, len(m.m) - 1}
	// goal := [2]int{5, 2}

	openSet := make(map[[2]int]bool)
	openSet[start] = true

	cameFrom := make(map[[2]int][2]int)

	gScore := make(map[[2]int]int)
	gScore[start] = 0

	fScore := make(map[[2]int]int)
	fScore[start] = m.heuristic(start, goal)

	for len(openSet) != 0 {

		current := getLowest(&openSet, &fScore)

		// if current[0] == len(m.m[0])-1 && current[1] == len(m.m)-1 {
		if current == goal {
			cost := 0
			path := m.reconstructPath(cameFrom, current)
			for _, s := range path {
				cost += m.getCost(s)
			}
			cost -= m.getCost(start)
			// m.printWithPath(path)
			return cost
		}
		delete(openSet, current)

		for _, neighbour := range m.getAdjacent2(current) {
			// tentativeGScore := m.heuristic(current) + m.getCost(neighbour)
			tentativeGScore := getScore(gScore, current) + m.getCost(neighbour)
			if tentativeGScore < getScore(gScore, neighbour) {
				cameFrom[neighbour] = current
				gScore[neighbour] = tentativeGScore
				fScore[neighbour] = tentativeGScore + m.heuristic(neighbour, goal)
				openSet[neighbour] = true
			}
		}
	}
	return -1
}

func rune2int(r rune) int {
	if r == 'S' {
		return -1
	}
	if r == 'E' {
		return int('z' - 'a')
	}
	return int(r - 'a')
}

func mainp1() {

	// b := tj.FileToSlice("input")
	b := tj.FileToSlice("input_real")
	// c := FileTo2DSliceRegex("input", " ")

	start := [2]int{0, 0}
	goal := [2]int{len(b[0]) - 1, len(b) - 1}
	m := heightmap{m: [][]int{}}
	for j, line := range b {
		m.m = append(m.m, []int{})
		for i, char := range line {
			m.m[j] = append(m.m[j], rune2int(char))
			if char == 'S' {
				start = [2]int{i, j}
			}
			if char == 'E' {
				goal = [2]int{i, j}
			}
		}
	}
	println(m.aStar(start, goal))

}

func mainp2() {

	// b := tj.FileToSlice("input")
	b := tj.FileToSlice("input_real")
	// c := FileTo2DSliceRegex("input", " ")

	starts := make([][2]int, 0)
	goal := [2]int{len(b[0]) - 1, len(b) - 1}
	m := heightmap{m: [][]int{}}
	for j, line := range b {
		m.m = append(m.m, []int{})
		for i, char := range line {
			m.m[j] = append(m.m[j], rune2int(char))
			if char == 'S' || char == 'a' {
				starts = append(starts, [2]int{i, j})
			}
			if char == 'E' {
				goal = [2]int{i, j}
			}
		}
	}
	min_distance := 1000000000
	for _, start := range starts {
		// println("Start: ", start[0], start[1])
		a := m.aStar(start, goal)
		if a < min_distance && a != -1 {
			min_distance = a
		}
	}
	println(min_distance)
	// println(m.aStar(start, goal))

}
func main() {
	mainp2()
}
