package main

import (
	"fmt"
	"sort"

	"github.com/fredriksoderberg/advent-of-code-2022/cast"
)

type Monkey struct {
	items       []int
	inspections int
	worryChange func(int) int
	nextMonkey  func(int) int
}

func main() {
	monkeys := createMonkeys()
	round := 1
	for round <= 20 {
		for i := range monkeys {
			monkeys[i].inspections += len(monkeys[i].items)
			for _, it := range monkeys[i].items {
				newWorry := monkeys[i].worryChange(it) / 3
				nextMonkey := monkeys[i].nextMonkey(newWorry)
				monkeys[nextMonkey].items = append(monkeys[nextMonkey].items, newWorry)
				monkeys[i].items = monkeys[i].items[1:]
			}
		}
		round++
	}
	fmt.Println(cast.ToString(monkeyBusiness(monkeys)))

	monkeys = createMonkeys()
	round = 1
	lcd := 9699690
	for round <= 10000 {
		for i := range monkeys {
			monkeys[i].inspections += len(monkeys[i].items)
			for _, it := range monkeys[i].items {
				newWorry := monkeys[i].worryChange(it) % lcd
				nextMonkey := monkeys[i].nextMonkey(newWorry)
				monkeys[nextMonkey].items = append(monkeys[nextMonkey].items, newWorry)
				monkeys[i].items = monkeys[i].items[1:]
			}
		}
		round++
	}
	fmt.Println(cast.ToString(monkeyBusiness(monkeys)))
}

func monkeyBusiness(monkeys []Monkey) int {
	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].inspections > monkeys[j].inspections
	})
	return monkeys[0].inspections * monkeys[1].inspections
}

func createMonkeys() []Monkey {
	monkeys := make([]Monkey, 0)
	m := Monkey{
		items:       []int{59, 65, 86, 56, 74, 57, 56},
		inspections: 0,
		worryChange: func(i int) int {
			return i * 17
		},
		nextMonkey: func(i int) int {
			if i%3 == 0 {
				return 3
			}
			return 6
		},
	}
	monkeys = append(monkeys, m)
	m = Monkey{
		items:       []int{63, 83, 50, 63, 56},
		inspections: 0,
		worryChange: func(i int) int {
			return i + 2
		},
		nextMonkey: func(i int) int {
			if i%13 == 0 {
				return 3
			}
			return 0
		},
	}
	monkeys = append(monkeys, m)
	m = Monkey{
		items:       []int{93, 79, 74, 55},
		inspections: 0,
		worryChange: func(i int) int {
			return i + 1
		},
		nextMonkey: func(i int) int {
			if i%2 == 0 {
				return 0
			}
			return 1
		},
	}
	monkeys = append(monkeys, m)
	m = Monkey{
		items:       []int{86, 61, 67, 88, 94, 69, 56, 91},
		inspections: 0,
		worryChange: func(i int) int {
			return i + 7
		},
		nextMonkey: func(i int) int {
			if i%11 == 0 {
				return 6
			}
			return 7
		},
	}
	monkeys = append(monkeys, m)
	m = Monkey{
		items:       []int{76, 50, 51},
		inspections: 0,
		worryChange: func(i int) int {
			return i * i
		},
		nextMonkey: func(i int) int {
			if i%19 == 0 {
				return 2
			}
			return 5
		},
	}
	monkeys = append(monkeys, m)
	m = Monkey{
		items:       []int{77, 76},
		inspections: 0,
		worryChange: func(i int) int {
			return i + 8
		},
		nextMonkey: func(i int) int {
			if i%17 == 0 {
				return 2
			}
			return 1
		},
	}
	monkeys = append(monkeys, m)
	m = Monkey{
		items:       []int{74},
		inspections: 0,
		worryChange: func(i int) int {
			return i * 2
		},
		nextMonkey: func(i int) int {
			if i%5 == 0 {
				return 4
			}
			return 7
		},
	}
	monkeys = append(monkeys, m)
	m = Monkey{
		items:       []int{86, 85, 52, 86, 91, 95},
		inspections: 0,
		worryChange: func(i int) int {
			return i + 6
		},
		nextMonkey: func(i int) int {
			if i%7 == 0 {
				return 4
			}
			return 5
		},
	}
	monkeys = append(monkeys, m)
	return monkeys
}
