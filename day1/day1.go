package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"

	"github.com/fredriksoderberg/advent-of-code-2022/cast"
)

func main() {
	readFile, err := os.Open("./day1/input_day1.txt")
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	current := 0
	max := make([]int, 4)
	for fileScanner.Scan() {
		if fileScanner.Text() == "" {
			max = addAndSort(current, max)
			current = 0
		} else {
			current += cast.ToInt(fileScanner.Text())
		}
	}
	fmt.Println(cast.ToString(max[0]))
	top3 := max[0] + max[1] + max[2]
	fmt.Println(cast.ToString(top3))
	readFile.Close()
}

func addAndSort(val int, vals []int) []int {
	vals = append(vals, val)
	sort.Slice(vals, func(i, j int) bool {
		return vals[i] > vals[j]
	})
	return vals[:3]
}
