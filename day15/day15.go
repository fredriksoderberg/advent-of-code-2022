package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"

	"github.com/fredriksoderberg/advent-of-code-2022/cast"
)

const filename = "./day15/input.txt"

type point struct {
	x int
	y int
}

func main() {
	//Read input file
	ppairs := readInput(filename)
	notpresent := make(map[point]rune)
	border := 10
	for _, ppair := range ppairs {
		if ppair[1].y == border {
			notpresent[ppair[1]] = 'B'
		}
		dist := manhattanDistance(ppair[0], ppair[1])
		nearestLine := math.Abs(float64(ppair[0].y - border))
		if dist >= int(nearestLine) {
			search := point{ppair[0].x, border}
			if _, ok := notpresent[search]; !ok {
				notpresent[search] = '#'
			}
			end := false
			// go right on the line
			for !end {
				search.x++
				if dist >= manhattanDistance(ppair[0], search) {
					if _, ok := notpresent[search]; !ok {
						notpresent[search] = '#'
					}
				} else {
					end = true
				}
			}
			search = point{ppair[0].x, border}
			end = false
			// go left on the line
			for !end {
				search.x--
				if dist >= manhattanDistance(ppair[0], search) {
					if _, ok := notpresent[search]; !ok {
						notpresent[search] = '#'
					}
				} else {
					end = true
				}
			}
		}
	}
	count := 0
	for _, v := range notpresent {
		if v == '#' {
			count++
		}
	}
	fmt.Println(count)
}

func readInput(filename string) [][]point {
	readFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	sc := bufio.NewScanner(readFile)
	defer readFile.Close()
	ppairs := make([][]point, 0)
	for sc.Scan() {
		p1 := point{}
		p2 := point{}
		re := regexp.MustCompile("-?[0-9]+")
		submatchall := re.FindAllString(sc.Text(), -1)
		for i, element := range submatchall {
			if i == 0 {
				p1.x = cast.StringToInt(element)
			}
			if i == 1 {
				p1.y = cast.StringToInt(element)
			}
			if i == 2 {
				p2.x = cast.StringToInt(element)
			}
			if i == 3 {
				p2.y = cast.StringToInt(element)
			}
		}
		ppairs = append(ppairs, []point{p1, p2})
	}
	return ppairs
}

func manhattanDistance(p1, p2 point) int {
	return int(math.Abs(float64(p1.x-p2.x)) + math.Abs(float64(p1.y-p2.y)))
}
