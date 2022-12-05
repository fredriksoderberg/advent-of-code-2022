package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fredriksoderberg/advent-of-code-2022/cast"
)

type stack []rune

func (s stack) Push(v rune) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, rune) {
	l := len(s)
	return s[:l-1], s[l-1]
}

func main() {
	crates, moves := readInput("./day5/input.txt")

	for _, move := range moves {
		c, o, d := parseMoves(move)
		i := 1
		for i <= c {
			cr, temp := crates[o-1].Pop()
			crates[o-1] = cr
			crates[d-1] = crates[d-1].Push(temp)
			i++
		}
	}

	top := strings.Builder{}
	for _, cr := range crates {
		_, r := cr.Pop()
		top.WriteRune(r)
	}

	fmt.Println(top.String())

	crates, moves = readInput("./day5/input.txt")

	for _, move := range moves {
		c, o, d := parseMoves(move)
		i := 1
		crateStash := make([]rune, 0)
		for i <= c {
			cr, temp := crates[o-1].Pop()
			crateStash = append(crateStash, temp)
			crates[o-1] = cr
			i++
		}
		for j := len(crateStash) - 1; j >= 0; j-- {
			crates[d-1] = crates[d-1].Push(crateStash[j])
		}
	}

	top = strings.Builder{}
	for _, cr := range crates {
		_, r := cr.Pop()
		top.WriteRune(r)
	}
	fmt.Println(top.String())
}

func emptyCrate(crate string) bool {
	for _, c := range crate {
		if c != ' ' {
			return false
		}
	}
	return true
}

func parseMoves(move string) (count int, origin int, dest int) {
	count = cast.ToInt(strings.Split(strings.Split(move, " from ")[0], "move ")[1])
	origin = cast.ToInt(strings.Split(strings.Split(move, " from ")[1], " to ")[0])
	dest = cast.ToInt(strings.Split(strings.Split(move, " from ")[1], " to ")[1])
	return
}

func readInput(input string) ([]stack, []string) {
	stackCount := 9
	crates := make([]stack, stackCount)
	for i := 0; i < stackCount; i++ {
		crates[i] = make(stack, 0)
	}

	moves := make([]string, 0)
	crateLine := make([]string, 0)

	readFile, err := os.Open(input)
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	// read crates
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if line[0:2] != " 1" {
			crateLine = append(crateLine, line)
		} else {
			fileScanner.Scan() // skip this line
			break
		}
	}
	for i := len(crateLine) - 1; i >= 0; i-- {
		line := crateLine[i]
		for i, j := 0, 0; i < len(line); i, j = i+4, j+1 {
			crate := line[i : i+2]
			if !emptyCrate(crate) {
				crates[j] = crates[j].Push(rune(crate[1]))
			}
		}
	}

	// read moves
	for fileScanner.Scan() {
		moves = append(moves, fileScanner.Text())
	}

	readFile.Close()
	return crates, moves
}
