package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/fredriksoderberg/advent-of-code-2022/cast"
)

const filename = "./day8/input.txt"

func main() {
	rows := 99
	columns := 99
	trees := readInput(filename, rows, columns)
	cols := make([][]int, len(trees[0]))
	for i := 0; i < len(trees[0]); i++ {
		cols[i] = make([]int, 0)
		for j := 0; j < len(trees); j++ {
			cols[i] = append(cols[i], trees[j][i])
		}
	}

	visible := 0
	max := 0
	for r, row := range trees {
		for c, tree := range row {
			if r == 0 || c == 0 || r == rows-1 || c == columns-1 {
				visible++
			} else if isVisible(row[c+1:], tree) {
				// right
				visible++
			} else if isVisible(row[:c], tree) {
				// left
				visible++
			} else if isVisible(cols[c][r+1:], tree) {
				// below
				visible++
			} else if isVisible(cols[c][:r], tree) {
				// above
				visible++
			}

			right := numVisible(row[c+1:], tree, false)     // right
			left := numVisible(row[:c], tree, true)         // left
			below := numVisible(cols[c][r+1:], tree, false) // below
			above := numVisible(cols[c][:r], tree, true)    // above
			score := right * left * below * above

			if score > max {
				max = score
			}
		}
	}

	fmt.Println(cast.ToString(visible))
	fmt.Println(cast.ToString(max))
}

func readInput(filename string, rows int, cols int) [][]int {
	input := make([][]int, rows)
	for i := range input {
		input[i] = make([]int, cols)
	}
	readFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	i := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		for j, it := range line {
			input[i][j] = cast.RuneToInt(it)
		}
		i++
	}
	readFile.Close()
	return input
}

func isVisible(row []int, n int) bool {
	for _, v := range row {
		if v >= n {
			return false
		}
	}
	return true
}

func numVisible(row []int, n int, reverse bool) int {
	count := 0
	if reverse {
		for i := len(row) - 1; i >= 0; i-- {
			count++
			if row[i] >= n {
				return count
			}
		}
	} else {
		for _, v := range row {
			count++
			if v >= n {
				return count
			}
		}
	}

	return count
}
