package main

import (
	"log"
	"strings"
	"time"

	tj "github.com/tjhowse/tjgo"
)

type stack struct {
	s string
}

func (s *stack) push(letter string) {
	s.s = s.s[1:14] + letter
}
func (s *stack) print() {
	println(s.s)
}

// Check if any character appears more than once in s.s
func (s *stack) dupe() bool {
	for _, r := range s.s {
		if strings.Count(s.s, string(r)) > 1 {

			return true
		}
	}
	return false
}

func better_dupe(s string) bool {
	for _, r := range s {
		if strings.Count(s, string(r)) > 1 {
			return false
		}
	}
	return true
}

func main() {
	// b := tj.FileToSlice("input")
	b := tj.FileToSlice("input_real")
	start := time.Now()
	// println(len(b))
	// c := tj.DoRegex("123-45", `(\d+)-(\d+)`)
	// for _, i := range c {
	// 	println(i)
	// }
	// return
	// for r := range b[0] {
	// 	println(string(rune(r)))
	// }
	input := b[0]

	// First, dumb, solution
	// s := stack{s: "000000000000000"}
	// for i := 0; i < len(input); i++ {
	// 	s.push(string(input[i]))
	// 	if !s.dupe() && i >= 4 {
	// 		print(i + 1)
	// 		return
	// 	}
	// }

	// Much clearer solution
	n := 14
	for i := n; i < len(input); i++ {
		if better_dupe(input[i-n : i]) {
			elapsed := time.Since(start)
			println(i)
			log.Printf("%s", elapsed)
			return
		}
	}
}
