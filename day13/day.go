package main

import (
	"fmt"
	"sort"

	tj "github.com/tjhowse/tjgo"
)

type packet struct {
	val int
	str string
	sub []packet
}

// This parses a string like [[1],[2,3,4]] and returns the next outermost
// complete bracketed chunk. I.E.
// [[1],[2,3,4]] -> [1],[2,3,4]
// [1],[2,3,4] -> 1
// [2,3,4] -> 2,3,4
// or
// [22,33,44] -> 22,33,44

func peel(i string) string {
	if i[0] != '[' {
		// println("Peeling", i)
		for j, c := range i {
			if c == ',' {
				return i[:j]
			}
		}
		return i
	}
	depth := 0
	for j, c := range i {
		if c == '[' {
			depth++
		}
		if c == ']' {
			depth--
		}
		if depth == 0 {
			return i[:j+1]
		}
	}
	return i
}

func (p *packet) print() {
	if p.val >= 0 {
		print(p.val)
		return
	}
	print("[")
	for i, sub := range p.sub {
		if i > 0 {
			print(",")
		}
		sub.print()
	}
}

// Accepts a string in the format [...]
func (p *packet) load(i string) {
	// println("Loading:", i)
	p.sub = make([]packet, 0)
	p.val = -1
	p.str = i
	if i[0] != '[' {
		p.val = tj.Str2int(i)
		return
	}
	i = i[1 : len(i)-1]
	for {
		if i == "" {
			break
		}
		s := peel(i)
		// println("peeled:", s)
		i = i[len(s):]
		if i != "" {
			i = i[1:]
		}
		packet := packet{}
		packet.load(s)
		p.sub = append(p.sub, packet)
	}

}

// This returns true if p comes before r
func (lhs *packet) compare(rhs packet) rune {
	if lhs.val >= 0 && rhs.val >= 0 {
		// Both values are numbers, not lists
		if lhs.val < rhs.val {
			return '<'
		} else if lhs.val > rhs.val {
			return '>'
		} else {
			return '='
		}
	}
	// if len(lhs.sub) > 0 && len(rhs.sub) == 0 {
	if lhs.val < 0 && rhs.val >= 0 {
		// println("lhs is list, rhs is not")
		temp := packet{}
		temp.load(fmt.Sprintf("[%d]", rhs.val))
		return lhs.compare(temp)
	}
	// if len(rhs.sub) > 0 && len(lhs.sub) == 0 {
	if lhs.val >= 0 && rhs.val < 0 {
		// println("lhs is list, rhs is not")
		temp := packet{}
		temp.load(fmt.Sprintf("[%d]", lhs.val))
		return temp.compare(rhs)
	}
	i := 0
	for ; i < len(lhs.sub) && i < len(rhs.sub); i++ {
		result := lhs.sub[i].compare(rhs.sub[i])
		// println(string(result))
		if result != '=' {
			return result
		}
	}
	if len(lhs.sub) < len(rhs.sub) {
		return '<'
	}
	if len(lhs.sub) == len(rhs.sub) {

		return '='
	}
	return '>'

}

func main() {

	// b := tj.FileToSlice("input")
	b := tj.FileToSlice("input_real")
	// c := FileTo2DSliceRegex("input", " ")
	packets := make([]packet, 0)
	for _, line := range b {
		if line != "" {
			packet := packet{}
			packet.load(line)
			packets = append(packets, packet)
			// packet.print()
		}
		// println()
	}
	answers := map[int]bool{}
	answers[1] = true
	answers[2] = true
	answers[3] = false
	answers[4] = true
	answers[5] = false
	answers[6] = true
	answers[7] = false
	answers[8] = false
	correct := 0
	checked := 0
	// for i := 0; i < 20; i += 2 {
	for i := 0; i < len(packets); i += 2 {
		println("=== Pair", int(i/2)+1, "===")
		ans := packets[i].compare(packets[i+1]) == '<'
		if ans {
			println("Correct order")
			correct += (i+1)/2 + 1
		} else {
			println("Incorrect order")
		}
		checked++
		if ans != answers[int(i/2)+1] {
			println("               Wrong")
		} else {
			println("               Right")
		}

	}
	println("Correct:", correct)
	println("checked:", checked)
	mark_a := packet{}
	mark_a.load("[[2]]")
	packets = append(packets, mark_a)
	mark_b := packet{}
	mark_b.load("[[6]]")
	packets = append(packets, mark_b)

	sort.Slice(packets, func(i, j int) bool {
		return packets[i].compare(packets[j]) == '<'
	})
	res := 1
	for i, p := range packets {
		// println(p.str)
		if p.str == "[[2]]" {
			println("Mark A:", i+1)
			res *= i + 1
		}
		if p.str == "[[6]]" {
			println("Mark B:", i+1)
			res *= i + 1
		}
		// println()
	}
	println(res)

}

// 5626 too high
