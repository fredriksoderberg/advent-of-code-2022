package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fredriksoderberg/advent-of-code-2022/cast"
)

const filename = "./day10/input.txt"
const addx = "addx"

type stack []string

func (s stack) Push(v string) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, string) {
	l := len(s)
	return s[:l-1], s[l-1]
}

func (s stack) Empty() bool {
	return len(s) == 0
}

func main() {
	ops := readInput(filename)
	var cycle, valX, sumStrength int
	cycle, valX = 1, 1
	for !ops.Empty() {
		if (cycle+20)%40 == 0 {
			sumStrength += cycle * valX
		}
		op, o := ops.Pop()
		ops = op
		opName := strings.Split(o, " ")[0]
		value := 0
		if opName == addx {
			value = cast.StringToInt(strings.Split(o, " ")[1])
			valX += value
		}
		cycle++
	}
	fmt.Println(cast.ToString(sumStrength))

	// make matrix of 6x40 with rune
	crt := make([][]rune, 6)
	for i := range crt {
		crt[i] = make([]rune, 40)
	}
	// fill matrix with runes
	for i := range crt {
		for j := range crt[i] {
			crt[i][j] = '.'
		}
	}
	//current position
	x, y := 0, 0
	cycle, valX = 1, 1
	ops = readInput(filename)
	for !ops.Empty() {
		if x == valX {
			crt[y][x] = '#'
		} else if x == valX-1 {
			crt[y][x] = '#'
		} else if x == valX+1 {
			crt[y][x] = '#'
		}
		op, o := ops.Pop()
		ops = op
		opName := strings.Split(o, " ")[0]
		value := 0
		if opName == addx {
			value = cast.StringToInt(strings.Split(o, " ")[1])
			valX += value
		}
		cycle++
		x++
		if x == 40 {
			x = 0
			y++
		}
	}

	// print matrix
	for i := range crt {
		for j := range crt[i] {
			fmt.Print(string(crt[i][j]))
		}
		fmt.Println()
	}
}

func readInput(filename string) stack {
	input := make([]string, 0)
	readFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		input = append(input, fileScanner.Text())
	}
	readFile.Close()
	ops := make(stack, 0)
	for i := len(input) - 1; i >= 0; i-- {
		if strings.Split(input[i], " ")[0] == "addx" {
			ops = ops.Push(input[i])
		}
		ops = ops.Push("noop")
	}
	return ops
}
