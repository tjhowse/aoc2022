package main

import (
	"strings"

	tj "github.com/tjhowse/tjgo"
)

type stack struct {
	s []string
}

func (s *stack) push(letter string) {
	s.s = append(s.s, letter)
}

func (s *stack) pop() string {
	if len(s.s) == 0 {
		return ""
	}
	letter := s.s[len(s.s)-1]
	s.s = s.s[:len(s.s)-1]
	return letter
}

func (s *stack) push_lots(letters string) {
	for _, letter := range letters {
		s.push(string(letter))
	}
}

func (s *stack) pop_lots(i int) string {
	if len(s.s) == 0 {
		return ""
	}
	temp := ""
	for j := 0; j < i; j++ {
		temp += s.pop()
	}
	// Reverse the order
	result := ""
	for j := len(temp) - 1; j >= 0; j-- {
		result += string(temp[j])
	}
	return result
	// return temp
}

func (s *stack) String() {
	for i := len(s.s) - 1; i >= 0; i-- {
		println(s.s[i])
	}
}

func main() {
	// b := tj.FileToSlice("input")
	b := tj.FileToSlice("input_real")
	// b := tj.FileToIntSlice("input")
	// b := tj.FileTo2DSlice("input", ',')
	// b := tj.FileTo2DSlice("input_real", ',')
	// b := tj.FileTo2DIntSlice("input")

	// stacks := make([]stack, 0)
	// stacks := [3]stack{}
	stacks := [9]stack{}
	// stacks_mode := true

	for i := 7; i >= 0; i-- {
		// for i := 2; i >= 0; i-- {
		println(b[i])
		line := b[i]
		for j := 1; j <= len(line); j += 4 {
			if string(line[j]) != " " {
				println("adding", string(line[j]), " to stack", (j-1)/4)
				stacks[(j-1)/4].push(string(line[j]))
			}
		}
	}
	// println("Start")
	// for i, stack := range stacks {
	// 	println("Stack", i+1)
	// 	stack.String()
	// 	println("---")
	// }

	// for i := 5; i < len(b); i++ {
	for i := 10; i < len(b); i++ {
		// Get numbers from string "move 1 from 2 to 1", splitting on spaces
		println(b[i])
		nums := strings.Split(b[i], " ")
		crate_count := tj.Str2int(nums[1])
		from := tj.Str2int(nums[3])
		to := tj.Str2int(nums[5])
		// for j := 0; j < crate_count; j++ {
		// 	stacks[to-1].push(stacks[from-1].pop())
		// }
		for _, stack := range stacks {
			stack.String()
		}
		println("---")
		stacks[to-1].push_lots(stacks[from-1].pop_lots(crate_count))
	}
	for _, stack := range stacks {
		stack.String()
	}
	println("---")

	print("Result: ")
	for _, stack := range stacks {
		// stack.String()
		print(stack.pop())
		// println("---")
	}

}
