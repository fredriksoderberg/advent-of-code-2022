package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/fredriksoderberg/advent-of-code-2022/cast"
)

const filename = "./day9/input.txt"

func main() {
	moves := readInput(filename)
	visited := make(map[string]int)
	var headX, headY, tailX, tailY, steps int
	var dir, key string
	visited["0,0"] = 1
	for _, mov := range moves {
		dir = strings.Split(mov, " ")[0]
		steps = cast.StringToInt(strings.Split(mov, " ")[1])
		for steps > 0 {
			switch dir {
			case "U":
				headY++
			case "D":
				headY--
			case "R":
				headX++
			case "L":
				headX--
			}
			if !adjacent(headX, tailX, headY, tailY) {
				tailX, tailY = newPosition(tailX, tailY, headX, headY)
				key = cast.ToString(tailX) + "," + cast.ToString(tailY)
				if _, ok := visited[key]; !ok {
					visited[key] = 1
				}
			}
			steps--
		}
	}
	fmt.Println(cast.ToString(len(visited)))

	visited = make(map[string]int)
	headsX := make([]int, 10)
	headsY := make([]int, 10)
	visited["0,0"] = 1
	for _, mov := range moves {
		dir = strings.Split(mov, " ")[0]
		steps = cast.StringToInt(strings.Split(mov, " ")[1])
		for steps > 0 {
			switch dir {
			case "U":
				headsY[0]++
			case "D":
				headsY[0]--
			case "R":
				headsX[0]++
			case "L":
				headsX[0]--
			}
			for i := 1; i < 10; i++ {
				if !adjacent(headsX[i], headsX[i-1], headsY[i], headsY[i-1]) {
					headsX[i], headsY[i] = newPosition(headsX[i], headsY[i], headsX[i-1], headsY[i-1])
					if i == 9 {
						key = cast.ToString(headsX[i]) + "," + cast.ToString(headsY[i])
						if _, ok := visited[key]; !ok {
							visited[key] = 1
						}
					}
				}
			}
			steps--
		}
	}
	fmt.Println(cast.ToString(len(visited)))
}

func readInput(filename string) []string {
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
	return input
}

func adjacent(x1 int, x2 int, y1 int, y2 int) bool {
	if math.Abs(float64(x1-x2)) <= 1 && math.Abs(float64(y1-y2)) <= 1 {
		return true
	}
	return false
}

func newPosition(curX int, curY int, prevX int, prevY int) (int, int) {
	if math.Abs(float64(curX-prevX)) == 0 {
		if prevY > curY {
			return curX, curY + 1
		}
		return curX, curY - 1
	}
	if math.Abs(float64(curY-prevY)) == 0 {
		if prevX > curX {
			return curX + 1, curY
		}
		return curX - 1, curY
	}
	if prevX > curX && prevY > curY {
		return curX + 1, curY + 1
	}
	if prevX > curX && prevY < curY {
		return curX + 1, curY - 1
	}
	if prevX < curX && prevY > curY {
		return curX - 1, curY + 1
	}
	if prevX < curX && prevY < curY {
		return curX - 1, curY - 1
	}
	return 0, 0
}
