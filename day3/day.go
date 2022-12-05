package main

import (
	tj "github.com/tjhowse/tjgo"
)

func process(a string) string {
	return a
}

func get_value(a string) int {
	n := int([]rune(a)[0] - 96)
	if n <= 0 {
		n += 26*2 + 6
	}
	return n
}

func main() {
	main2()
}

func get_all_map() map[int]int {
	all := make(map[int]int)
	for i := 1; i < 53; i++ {
		all[i] = 1
	}
	return all
}

func get_char_map() map[rune]int {
	all := make(map[rune]int)
	for i := 65; i <= 66+26; i++ {
		all[rune(i)] = 0
	}
	for i := 97; i <= 97+26; i++ {
		all[rune(i)] = 0
	}
	// for k, v := range all {
	// 	println(string(k), v)
	// }
	return all
}

func main2() {
	// b := tj.FileToSlice("input")
	b := tj.FileToSlice("input_real")

	all := get_all_map()
	sums := get_char_map()
	total := 0
	count := 0
	for _, line := range b {
		for _, c := range line {
			// println("'", string(c), "'")
			sums[c]++
		}
		for k, v := range sums {
			// println(k, v)
			if v == 0 {
				all[get_value(string(k))] = 0
			}
		}
		sums = get_char_map()
		println(line)
		count++
		if (count % 3) == 0 {
			for k, v := range all {
				// println(k, v)
				if v != 0 {
					println(k)
					total += k
				}
			}
			all = get_all_map()
			println()
		}
	}
	println(total)

}

func main1() {
	b := tj.FileToSlice("input")
	// b := tj.FileToSlice("input_real")

	total := 0

	for _, line := range b {
		// total += d
		// println(process(line))
		first := line[0:int(len(line)/2)]
		second := line[int(len(line)/2):]
		println(first)
		println(second)
		dupe := make(map[rune]bool)
		for _, f := range first {
			for _, s := range second {
				if f == s {
					// dupe = append(dupe, f)
					dupe[f] = true
					// println(string(f))
					// total += get_value(string(f))
				}
			}
		}
		for k, _ := range dupe {
			total += get_value(string(k))
		}

	}
	println(total)

}
