package main

type monkey struct {
	holding       []int
	operation     func(int) int
	testDivisible int
	trueAction    int
	falseAction   int
	inspections   int
}

func (m *monkey) turn(monkeys []monkey) {
	for len(m.holding) > 0 {
		item := m.holding[0]
		if len(m.holding) == 1 {
			m.holding = m.holding[:0]
		} else {
			m.holding = m.holding[1:]
		}
		item = m.operation(item)
		m.inspections++
		item = int(item / 3)
		destination := 0
		if item%m.testDivisible == 0 {
			destination = m.trueAction
		} else {
			destination = m.falseAction
		}
		monkeys[destination].throw(item)
	}
}
func (m *monkey) throw(item int) {
	m.holding = append(m.holding, item)
}

func (m *monkey) turnpart2(monkeys []monkey, divisor int) {
	for len(m.holding) > 0 {
		item := m.holding[0]
		if len(m.holding) == 1 {
			m.holding = m.holding[:0]
		} else {
			m.holding = m.holding[1:]
		}
		item %= divisor
		item = m.operation(item)
		m.inspections++
		// item = int(item / 3)
		destination := 0
		if item%m.testDivisible == 0 {
			destination = m.trueAction
		} else {
			destination = m.falseAction
		}
		monkeys[destination].throw(item)
	}
}

func main_demo() {

	monkeys := []monkey{}

	monkeys = append(monkeys, monkey{
		holding: []int{79, 98},
		operation: func(i int) int {
			return i * 19
		},
		testDivisible: 23,
		trueAction:    2,
		falseAction:   3,
	})
	monkeys = append(monkeys, monkey{
		holding: []int{54, 65, 75, 74},
		operation: func(i int) int {
			return i + 6
		},
		testDivisible: 19,
		trueAction:    2,
		falseAction:   0,
	})
	monkeys = append(monkeys, monkey{
		holding: []int{79, 60, 97},
		operation: func(i int) int {
			return i * i
		},
		testDivisible: 13,
		trueAction:    1,
		falseAction:   3,
	})
	monkeys = append(monkeys, monkey{
		holding: []int{74},
		operation: func(i int) int {
			return i + 3
		},
		testDivisible: 17,
		trueAction:    0,
		falseAction:   1,
	})

	for j := 0; j < 20; j++ {
		for i := 0; i < len(monkeys); i++ {
			println("Monkey", i, "has", len(monkeys[i].holding), "items")
			monkeys[i].turn(monkeys)
		}
	}
	for i, monkey := range monkeys {
		println("Monkey", i, "has", len(monkey.holding), "items:")
		// for _, item := range monkey.holding {
		// 	println(item)
		// }
		println("Inspected", monkey.inspections, "items")
	}

}

func main_part1() {

	monkeys := []monkey{}

	// 0
	monkeys = append(monkeys, monkey{
		holding: []int{54, 89, 94},
		operation: func(i int) int {
			return i * 7
		},
		testDivisible: 17,
		trueAction:    5,
		falseAction:   3,
	})
	// 1
	monkeys = append(monkeys, monkey{
		holding: []int{66, 71},
		operation: func(i int) int {
			return i + 4
		},
		testDivisible: 3,
		trueAction:    0,
		falseAction:   3,
	})
	// 2
	monkeys = append(monkeys, monkey{
		holding: []int{76, 55, 80, 55, 55, 96, 78},
		operation: func(i int) int {
			return i + 2
		},
		testDivisible: 5,
		trueAction:    7,
		falseAction:   4,
	})
	// 3
	monkeys = append(monkeys, monkey{
		holding: []int{93, 69, 76, 66, 89, 54, 59, 94},
		operation: func(i int) int {
			return i + 7
		},
		testDivisible: 7,
		trueAction:    5,
		falseAction:   2,
	})
	// 4
	monkeys = append(monkeys, monkey{
		holding: []int{80, 54, 58, 75, 99},
		operation: func(i int) int {
			return i * 17
		},
		testDivisible: 11,
		trueAction:    1,
		falseAction:   6,
	})
	// 5
	monkeys = append(monkeys, monkey{
		holding: []int{69, 70, 85, 83},
		operation: func(i int) int {
			return i + 8
		},
		testDivisible: 19,
		trueAction:    2,
		falseAction:   7,
	})
	// 6
	monkeys = append(monkeys, monkey{
		holding: []int{89},
		operation: func(i int) int {
			return i + 6
		},
		testDivisible: 2,
		trueAction:    0,
		falseAction:   1,
	})
	// 7
	monkeys = append(monkeys, monkey{
		holding: []int{62, 80, 58, 57, 93, 56},
		operation: func(i int) int {
			return i * i
		},
		testDivisible: 13,
		trueAction:    6,
		falseAction:   4,
	})

	// 332*334

	for j := 0; j < 20; j++ {
		for i := 0; i < len(monkeys); i++ {
			// println("Monkey", i, "has", len(monkeys[i].holding), "items")
			monkeys[i].turn(monkeys)
		}
	}
	for i, monkey := range monkeys {
		println("Monkey", i, "has", len(monkey.holding), "items:")
		// for _, item := range monkey.holding {
		// 	println(item)
		// }
		println("Inspected", monkey.inspections, "items")
	}

}

func main() {

	monkeys := []monkey{}

	// 0
	monkeys = append(monkeys, monkey{
		holding: []int{54, 89, 94},
		operation: func(i int) int {
			return i * 7
		},
		testDivisible: 17,
		trueAction:    5,
		falseAction:   3,
	})
	// 1
	monkeys = append(monkeys, monkey{
		holding: []int{66, 71},
		operation: func(i int) int {
			return i + 4
		},
		testDivisible: 3,
		trueAction:    0,
		falseAction:   3,
	})
	// 2
	monkeys = append(monkeys, monkey{
		holding: []int{76, 55, 80, 55, 55, 96, 78},
		operation: func(i int) int {
			return i + 2
		},
		testDivisible: 5,
		trueAction:    7,
		falseAction:   4,
	})
	// 3
	monkeys = append(monkeys, monkey{
		holding: []int{93, 69, 76, 66, 89, 54, 59, 94},
		operation: func(i int) int {
			return i + 7
		},
		testDivisible: 7,
		trueAction:    5,
		falseAction:   2,
	})
	// 4
	monkeys = append(monkeys, monkey{
		holding: []int{80, 54, 58, 75, 99},
		operation: func(i int) int {
			return i * 17
		},
		testDivisible: 11,
		trueAction:    1,
		falseAction:   6,
	})
	// 5
	monkeys = append(monkeys, monkey{
		holding: []int{69, 70, 85, 83},
		operation: func(i int) int {
			return i + 8
		},
		testDivisible: 19,
		trueAction:    2,
		falseAction:   7,
	})
	// 6
	monkeys = append(monkeys, monkey{
		holding: []int{89},
		operation: func(i int) int {
			return i + 6
		},
		testDivisible: 2,
		trueAction:    0,
		falseAction:   1,
	})
	// 7
	monkeys = append(monkeys, monkey{
		holding: []int{62, 80, 58, 57, 93, 56},
		operation: func(i int) int {
			return i * i
		},
		testDivisible: 13,
		trueAction:    6,
		falseAction:   4,
	})

	// 332*334

	divisor := 1

	for _, monkey := range monkeys {
		divisor *= monkey.testDivisible
	}

	for j := 0; j < 10000; j++ {
		for i := 0; i < len(monkeys); i++ {
			// println("Monkey", i, "has", len(monkeys[i].holding), "items")
			monkeys[i].turnpart2(monkeys, divisor)
		}
	}
	for i, monkey := range monkeys {
		println("Monkey", i, "has inspectd", monkey.inspections, "items")
	}

}
